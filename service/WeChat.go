package service

import (
	"github.com/imroc/req"
	"github.com/silenceper/wechat/cache"
	platform "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniconfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	oaconfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var OFFICILACCOUNT *officialaccount.OfficialAccount //公众号
var MINIPROGRAM *miniprogram.MiniProgram            //小程序

func InitOfficialAccount(appid, appsecret, akurl string) {
	OFFICILACCOUNT = platform.NewWechat().GetOfficialAccount(&oaconfig.Config{AppID: appid, AppSecret: appsecret, Cache: cache.NewMemory()})
	if akurl != "" {
		OFFICILACCOUNT.SetAccessTokenHandle(CustomTokenHandle{Appid: appid, Akurl: akurl})
	}

}

func InitMiniProgram(appid, appsecret, akurl string) {
	MINIPROGRAM = platform.NewWechat().GetMiniProgram(&miniconfig.Config{AppID: appid, AppSecret: appsecret, Cache: cache.NewMemory()})
	if akurl != "" {
		MINIPROGRAM.SetAccessTokenHandle(CustomTokenHandle{Appid: appid, Akurl: akurl})
	}
}

type CustomTokenHandle struct {
	Appid string
	Akurl string
}

func (c CustomTokenHandle) GetAccessToken() (accessToken string, err error) {
	res, err := req.Get(c.Akurl)
	if err != nil {
		return "", err
	}
	type ApiRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	token := new(ApiRes)
	err = res.ToJSON(token)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
