package middlewares

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// Claims Jwt参数
type Claims struct {
	UID      uint   //可以是用户id，可以是管理员id
	OpenId   string //关联的openid
	RoleKey  string //角色id
	UserName string //用户名
	DeptId   uint   //部门id
	Isadmin  bool   //是否是管理员
	jwt.StandardClaims
}

type Config struct {
	Jwtkey string
}

// Jwt 中间件
func Jwt(key string, whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(whitelist, c.Request.URL.Path) {
			c.Next()
			return
		}
		Authorization := c.GetHeader("Authorization")
		l := len("Bearer")
		if len(Authorization) > l+1 && Authorization[:l] == "Bearer" {
			Authorization = Authorization[l+1:]
		} else {
			c.AbortWithStatusJSON(401, gin.H{
				"msg":  "请登陆后访问",
				"data": nil,
			})
			return
		}
		conf := &Config{Jwtkey: key}
		data, err := conf.ParseToken(Authorization)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"msg":  "Token 解析失败:" + err.Error(),
				"data": nil,
			})
			return
		}
		c.Set("UID", data.UID)
		c.Set("RoleKey", data.RoleKey)
		c.Set("OpenId", data.OpenId)
		c.Set("UserName", data.UserName)
		c.Set("DeptId", data.DeptId)
		c.Set("Isadmin", data.Isadmin)
		c.Next()
	}
}

// ParseToken 解析jwt
func (c *Config) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.Jwtkey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("非法令牌")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("已过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("未生效，请检查设备时间")
			} else {
				return nil, errors.New("验证失败，请重新获取")
			}
		}
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// GenerateToken 生成JWT
func (c *Config) GenerateToken(uid uint, DeptId uint, openid, rolekey, username string, isadmin bool) (string, error) {
	claims := Claims{
		uid,
		openid,
		rolekey,
		username,
		DeptId,
		isadmin,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60*60,                 //允许误差一小时
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), //7天自动过期
			Issuer:    "BugStark",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(c.Jwtkey))
	return token, err
}
