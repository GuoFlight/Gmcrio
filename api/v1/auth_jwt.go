package v1

import (
	"Gmicro/conf"
	"Gmicro/db"
	"Gmicro/logger"
	"Gmicro/myctx"
	"context"
	"github.com/GuoFlight/gerror/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const HeaderNameAuth = "Authorization"

var GJwt Jwt

type Jwt struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (j Jwt) Auth(c *gin.Context) bool {
	ctx := myctx.GenFromGin(c)
	tokenStr, ok := j.GetTokenFromReq(c)
	if !ok {
		c.JSON(401, Res{
			Code: 1,
			Msg:  "no token",
		})
		c.Abort()
		return false
	}
	jwtCliams, gerr := j.ParseToken(ctx, tokenStr)
	if gerr != nil {
		c.JSON(401, Res{
			Code:    1,
			Msg:     gerr.Error(),
			TraceId: gerr.TraceID,
		})
		c.Abort()
		return false
	}
	c.Set(conf.KeyUsername, jwtCliams.Username)
	c.Next()
	return true
}

// GetTokenFromReq 从请求头中获取token
func (Jwt) GetTokenFromReq(c *gin.Context) (string, bool) {
	tokens, ok := c.Request.Header[HeaderNameAuth]
	if !ok {
		return "", false
	}
	return tokens[0], true
}

// ParseToken 解析Token
func (Jwt) ParseToken(ctx context.Context, tokenStr string) (*Jwt, *gerror.Gerr) {
	token, err := jwt.ParseWithClaims(tokenStr, &Jwt{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(conf.GConf.Jwt.Secret), nil
	})
	if err != nil {
		return nil, logger.ErrWithCtx(ctx, err.Error())
	}
	if !token.Valid { // 过期也会返回error
		return nil, gerror.NewErr("invalid token")
	}

	if claims, ok := token.Claims.(*Jwt); ok { // 校验token
		return claims, nil
	}
	return nil, gerror.NewErr("invalid token")
}

// Login 用户登录，返回token
func (j Jwt) Login(c *gin.Context) {
	ctx := myctx.GenFromGin(c)
	// 获取请求参数
	var req ReqLogin
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, Res{
			Code: 400,
			Msg:  "invalid request",
		})
		return
	}
	// 验证账号密码
	ok, gerr := db.GUser.CheckPwd(ctx, req.Username, req.Password)
	if gerr != nil {
		c.JSON(http.StatusOK, Res{
			Code:    1,
			Msg:     gerr.Error(),
			TraceId: gerr.TraceID,
		})
		return
	}
	if !ok {
		c.JSON(http.StatusOK, Res{
			Code: 401,
			Msg:  "invalid username or password",
		})
		return
	}

	// 生成token
	token, err := j.GenToken(ctx, req.Username)
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code:    1,
			Msg:     err.Error(),
			TraceId: gerr.TraceID,
		})
		return
	}
	c.JSON(http.StatusOK, Res{
		Code: 0,
		Data: token,
	})
	return
}

func (Jwt) GenToken(ctx context.Context, username string) (string, *gerror.Gerr) {
	jwtClaims := Jwt{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(conf.GConf.Jwt.Expire) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenStr, err := token.SignedString([]byte(conf.GConf.Jwt.Secret))
	if err != nil {
		return "", logger.ErrWithCtx(ctx, err.Error())
	}
	return tokenStr, nil
}
