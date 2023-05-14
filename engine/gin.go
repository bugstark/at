package engine

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

// Query 查询参数
type Query struct {
	Fields string `gorm:"-" form:"fields" json:"-" binding:"safe"`
	Sort   string `gorm:"-" form:"sort" json:"-" binding:"safe"`
	Order  string `gorm:"-" form:"order" json:"-" binding:"safe"`
	Size   int    `gorm:"-" form:"size" json:"-"`
	Page   int    `gorm:"-" form:"page" json:"-"`
}

func NewEngine(debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var app = gin.New()
	var safe validator.Func = func(fl validator.FieldLevel) bool {
		var st = fl.Field().String()
		if st == "" {
			return true
		}
		pass, err := regexp.MatchString("^[A-Za-z0-9,.]+$", st)
		if err != nil {
			log.Println("Safe Check:", err.Error())
			return false
		}
		return pass
	}
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validate.RegisterValidation("safe", safe)
	}
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
