package service

import (
	"errors"
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

type WeChatModel struct {
	OFFICILACCOUNT *officialaccount.OfficialAccount
	MINIPROGRAM    *miniprogram.MiniProgram
	AppID          string
	AppSecret      string
	UseUNI         bool
	Cache          cache.Cache
}

var Wechat *WeChatModel

func InitWeChat(appid, appsceret string, useuni bool, cache cache.Cache, isminiprogram bool) bool {
	Wechat = &WeChatModel{
		AppID:     appid,
		AppSecret: appsceret,
		UseUNI:    useuni,
		Cache:     cache,
	}
	if isminiprogram {
		Wechat.InitMiniProgram()
	} else {
		Wechat.InitOfficialAccount()
	}
	_, err := Wechat.GetAccessToken()
	if err != nil {
		panic(err)
	}
	return true
}

func (w *WeChatModel) InitOfficialAccount() {
	w.OFFICILACCOUNT = platform.NewWechat().GetOfficialAccount(&oaconfig.Config{AppID: w.AppID, AppSecret: w.AppSecret, Cache: w.Cache})
	if w.UseUNI {
		w.OFFICILACCOUNT.SetAccessTokenHandle(w)
	}
}

func (w *WeChatModel) InitMiniProgram() {
	w.MINIPROGRAM = platform.NewWechat().GetMiniProgram(&miniconfig.Config{AppID: w.AppID, AppSecret: w.AppSecret, Cache: w.Cache})
	if w.UseUNI {
		w.MINIPROGRAM.SetAccessTokenHandle(w)
	}
}

// 自定义获取access_token的方法
func (w *WeChatModel) GetAccessToken() (accessToken string, err error) {
	key := fmt.Sprintf("custom_%s_access_token", w.AppID)
	token := w.Cache.Get(key)
	if token != nil {
		log.Println("get access_token from cache")
		return token.(string), nil
	}
	r, err := req.Get("https://uni.an2.cn/open/token?appid=" + w.AppID)
	if err != nil {
		return "", err
	}
	type ApiRes struct {
		Msg  string `json:"msg"`
		Data struct {
			AccessToken string `json:"access_token"`
			ExpiresAt   int    `json:"expires_at"`
			ExpiresIn   int    `json:"expires_in"`
		} `json:"data"`
	}
	var res ApiRes
	err = r.ToJSON(&res)
	if err != nil {
		return "", err
	}
	if res.Msg != "" {
		return "", errors.New(res.Msg)
	}
	err = w.Cache.Set(key, res.Data.AccessToken, time.Second*time.Duration(res.Data.ExpiresIn))
	if err != nil {
		return "", err
	}
	log.Println("get access_token from Uni")
	return res.Data.AccessToken, nil
}

func (w *WeChatModel) GetMPUserAccessToken(code string) {
	ak, err := w.OFFICILACCOUNT.GetAccessToken()
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/component/access_token?appid=%s&code=%s&grant_type=authorization_code&component_appid=wxbd5c9e30b31c7cc4&component_access_token=%s", w.AppID, code, ak)
	log.Println(url)
	res, err := req.Get(url)
	if err != nil {
		panic(err)
	}
	log.Println(res.String())
}
