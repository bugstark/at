package captcha

import "github.com/mojocn/base64Captcha"

// TODO redis 验证码储存
// type Store interface {
// 	// Set sets the digits for the captcha id.
// 	Set(id string, value string) error

// 	// Get returns stored digits for the captcha id. Clear indicates
// 	// whether the captcha must be deleted from the store.
// 	Get(id string, clear bool) string

// 	//Verify captcha's answer directly
// 	Verify(id, answer string, clear bool) bool
// }

var store = base64Captcha.DefaultMemStore
var driver base64Captcha.Driver = base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)

// 生成验证码
func Generate() (id, b64s string, err error) {
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err = c.Generate()
	return
}

// 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	return store.Verify(id, val, true)
}
