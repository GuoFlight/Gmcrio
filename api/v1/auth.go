package v1

import "github.com/gin-gonic/gin"

var GAuth Auth

type Auth struct{}

// Auth 用于认证的中间件
func (Auth) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		GJwt.Auth(c)
	}
}
