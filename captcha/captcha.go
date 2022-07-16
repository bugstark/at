package captcha

import "github.com/mojocn/base64Captcha"

var store = base64Captcha.DefaultMemStore
var driver base64Captcha.Driver = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)

// 生成验证码
func Generate() (id, b64s string, err error) {
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return
}

// 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	return store.Verify(id, val, true)
}
