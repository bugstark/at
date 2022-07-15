package service

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	v20180301 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
)

func (t *Tencent) GetBIZ(name, idcard, redirect, extra, RuleID string) (response *v20180301.DetectAuthResponse, err error) {
	credential := common.NewCredential(t.Appid, t.Secret)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	client, err := v20180301.NewClient(credential, "ap-chengdu", cpf)
	if err != nil {
		return nil, err
	}
	val := v20180301.NewDetectAuthRequest()
	val.Name = &name
	val.IdCard = &idcard
	val.Extra = &extra
	val.RedirectUrl = &redirect
	val.RuleId = common.StringPtr(RuleID)
	return client.DetectAuth(val)
}

func (t *Tencent) CheckBiz(biz, RuleID string) (response *v20180301.GetDetectInfoEnhancedResponse, err error) {
	credential := common.NewCredential(t.Appid, t.Secret)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "faceid.tencentcloudapi.com"
	client, err := v20180301.NewClient(credential, "ap-chengdu", cpf)
	if err != nil {
		return nil, err
	}
	val := v20180301.NewGetDetectInfoEnhancedRequest()
	val.BizToken = &biz
	val.RuleId = common.StringPtr(RuleID)
	return client.GetDetectInfoEnhanced(val)
}
