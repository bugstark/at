package service

import (
	"fmt"
	"log"
	"time"

	"github.com/imroc/req"
	platform "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniconfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	oaconfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var OFFICILACCOUNT *officialaccount.OfficialAccount //公众号
var MINIPROGRAM *miniprogram.MiniProgram            //小程序

func InitOfficialAccount(appid, appsecret, akurl string, cache cache.Cache) {
	OFFICILACCOUNT = platform.NewWechat().GetOfficialAccount(&oaconfig.Config{AppID: appid, AppSecret: appsecret, Cache: cache})
	if akurl != "" {
		OFFICILACCOUNT.SetAccessTokenHandle(CustomTokenHandle{Appid: appid, Akurl: akurl, Cache: cache})
	}
}

func InitMiniProgram(appid, appsecret, akurl string, cache cache.Cache) {
	MINIPROGRAM = platform.NewWechat().GetMiniProgram(&miniconfig.Config{AppID: appid, AppSecret: appsecret, Cache: cache})
	if akurl != "" {
		MINIPROGRAM.SetAccessTokenHandle(CustomTokenHandle{Appid: appid, Akurl: akurl, Cache: cache})
	}
}

type CustomTokenHandle struct {
	Appid string
	Akurl string
	Cache cache.Cache
}

// 自定义获取access_token的方法
func (c CustomTokenHandle) GetAccessToken() (accessToken string, err error) {
	key := fmt.Sprintf("custom_%s_access_token", c.Appid)
	token := c.Cache.Get(key)
	if token != nil {
		log.Println("get access_token from cache")
		return token.(string), nil
	}
	r, err := req.Get(c.Akurl)
	if err != nil {
		return "", err
	}
	type ApiRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var res ApiRes
	err = r.ToJSON(&res)
	if err != nil {
		return "", err
	}
	err = c.Cache.Set(key, res.AccessToken, time.Second*time.Duration(res.ExpiresIn))
	if err != nil {
		return "", err
	}
	log.Println("get access_token from " + c.Akurl)
	return res.AccessToken, nil
}
