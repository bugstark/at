package service

import (
	"fmt"
	"strings"

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

// Phone 逗号分隔，最多200个
func (t *TencentConig) SendSmS(phone, sign, tempid string, args []string) error {
	credential := common.NewCredential(t.Appid, t.Secret)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)
	request := sms.NewSendSmsRequest()
	phonelist := strings.Split(phone, ",")
	if len(phonelist) > 200 {
		return errors.NewTencentCloudSDKError("0", "超出200个手机号", "0000-0000-0000")
	}
	request.PhoneNumberSet = common.StringPtrs(phonelist)
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
