package v1

import (
	"Gmicro/conf"
	"Gmicro/logger"
	"github.com/GuoFlight/gerror"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	HeaderNameAuth = "Authorization"
	secretKey      = "MySecret@1216"
)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Auth struct{}

var GAuth Auth

func (a Auth) Login(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, Res{
			Code: 400,
			Msg:  "invalid request",
		})
		return
	}
	if user, ok := conf.GConf.Auth.Users[req.Username]; ok && req.Password == user.Password {
		token, err := a.GenToken(req.Username)
		if err != nil {
			c.JSON(500, Res{
				Code: 500,
				Msg:  err.Error(),
			})
			return
		}
		c.JSON(200, Res{
			Code: 0,
			Data: token,
		})
		return
	} else {
		c.JSON(401, Res{
			Code: 401,
			Msg:  "invalid username or password",
		})
		return
	}
}

func (a Auth) GenToken(username string) (string, error) {
	c := JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(conf.GConf.Auth.SecondTokenExpire) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secretKey))
}

func (a Auth) Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := a.GetTokenFromReq(c)
		if token == "" {
			c.JSON(401, Res{
				Code: 1,
				Msg:  "no Authorization",
			})
			c.Abort()
			return
		}
		jwtClaims, gerr := a.ParseToken(token)
		if gerr != nil {
			c.JSON(401, Res{
				Code: 1,
				Msg:  gerr.Error(),
			})
			c.Abort()
			return
		}
		c.Set("username", jwtClaims.Username)
		c.Next()
	}
}

func (a Auth) GetTokenFromReq(c *gin.Context) string {
	if token, ok := c.Request.Header[HeaderNameAuth]; ok {
		return token[0]
	} else {
		return ""
	}
}

// ParseToken 解析Token
func (a Auth) ParseToken(tokenString string) (*JwtClaims, *gerror.Gerr) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, logger.PrintErr(gerror.NewErr(err.Error()), nil)
	}
	if !token.Valid { // 过期也会返回error
		return nil, gerror.NewErr("invalid token")
	}

	if claims, ok := token.Claims.(*JwtClaims); ok { // 校验token
		return claims, nil
	}
	return nil, gerror.NewErr("invalid token")
}

// checkPermWrite 判断用户是否有写权限
func (a Auth) checkPermWrite(username string) bool {
	if user, ok := conf.GConf.Auth.Users[username]; ok && (user.Role == conf.RoleAdmin || user.Role == conf.RoleWriter) {
		return true
	}
	return false
}

// CheckPermWrite 判断用户是否有写权限。返回值：是否继续执行
func (a Auth) CheckPermWrite(c *gin.Context) bool {
	resNoPerm := Res{
		Code: 403,
		Msg:  "No permission to write",
	}
	username, ok := c.Get("username")
	if !ok {
		c.JSON(403, resNoPerm)
		return false
	}
	if !a.checkPermWrite(username.(string)) {
		c.JSON(403, resNoPerm)
		return false
	}
	return true
}
