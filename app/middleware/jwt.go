package middleware

import (
	"github.com/cilidm/toolbox/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
			msg  string
			data interface{}
		)
		code = 200
		token := c.PostForm("token")
		if token == "" {
			code = 400
			msg = "请求参数错误"
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = 20001
				msg = "Token鉴权失败"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 20002
				msg = "Token已超时"
			}
		}
		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
