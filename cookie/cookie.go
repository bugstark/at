package cookie

import (
	"net/http"
)

func CookieStr2HttpCookie(Cookiestr string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", Cookiestr)
	request := http.Request{Header: header}
	return request.Cookies()
}
