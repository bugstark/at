package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Casbin(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := c.GetString("AuthID")
		pass, err := e.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"data": nil, "msg": err.Error()})
			return
		}
		if pass {
			c.Next()
		} else {
			c.AbortWithStatusJSON(403, gin.H{"data": nil, "msg": "无访问权限！"})
			return
		}
	}
}
