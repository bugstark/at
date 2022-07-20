package service

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

var Tencent *TencentConig

//腾讯云服务
type TencentConig struct {
	Appid  string
	Secret string
}

func InitTencent(appid, secret string) {
	Tencent = &TencentConig{
		Appid:  appid,
		Secret: secret,
	}
}

func (t *TencentConig) SendSmS(phone, sign, tempid string, args []string) error {
	credential := common.NewCredential(t.Appid, t.Secret)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{phone})
	request.SmsSdkAppId = common.StringPtr("1400616420") //默认
	request.SignName = common.StringPtr(sign)
	request.TemplateId = common.StringPtr(tempid)
	request.TemplateParamSet = common.StringPtrs(args)
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	fmt.Printf("%s", response.ToJsonString())
	return err
}
