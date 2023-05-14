package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func Casbin(e *casbin.Enforcer, whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		if slices.Contains(whitelist, obj) {
			c.Next()
			return
		}
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
