package service

import (
	"log"
	"testing"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/officialaccount"
)

func TestWeChatModel_GetMPUserAccessToken(t *testing.T) {
	type fields struct {
		OFFICILACCOUNT *officialaccount.OfficialAccount
		MINIPROGRAM    *miniprogram.MiniProgram
		AppID          string
		AppSecret      string
		UseUNI         bool
		Cache          cache.Cache
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				AppID:     "wxd4989d54758be297",
				AppSecret: "",
				UseUNI:    true,
				Cache:     cache.NewMemory(),
			},
			args: args{
				code: "003UAc100XQsdO1Krv200Xhff83UAc1i",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WeChatModel{
				OFFICILACCOUNT: tt.fields.OFFICILACCOUNT,
				MINIPROGRAM:    tt.fields.MINIPROGRAM,
				AppID:          tt.fields.AppID,
				AppSecret:      tt.fields.AppSecret,
				UseUNI:         tt.fields.UseUNI,
				Cache:          tt.fields.Cache,
			}
			w.InitOfficialAccount()
			log.Println(w.GetMPUserAccessToken(tt.args.code))
		})
	}
}
