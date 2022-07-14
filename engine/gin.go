package engine

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewEngine(debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var app = gin.New()
	app.MaxMultipartMemory = 8 << 20
	app.Use(gin.Recovery())
	app.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{
			"msg":  fmt.Sprintf("not found '%s:%s'", c.Request.Method, c.Request.URL.Path),
			"data": nil,
		})
	})
	return app
}
