package service

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/imroc/req"
)

// 中国移动MAS短信平台
type MAS struct {
	EcName    string
	Apid      string
	SecretKey string
	Sign      string
	Apiurl    string
}

func NewMAS(ecname, apid, secretkey, sign, apiurl string) *MAS {
	return &MAS{
		EcName:    ecname,
		Apid:      apid,
		SecretKey: secretkey,
		Sign:      sign,
		Apiurl:    apiurl,
	}
}

// SendSMS 多个电话号码,英文逗号分隔
func (m *MAS) SendSMS(phone, content string) error {
	data := map[string]string{}
	data["ecName"] = m.EcName
	data["apId"] = m.Apid
	data["mobiles"] = phone
	data["content"] = content
	data["sign"] = m.Sign
	data["addSerial"] = ""
	data["mac"] = fmt.Sprintf("%x", md5.Sum([]byte(m.EcName+m.Apid+m.SecretKey+phone+content+m.Sign)))
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	base64stringjson := base64.StdEncoding.EncodeToString(jsondata)
	res, err := req.Post(m.Apiurl, base64stringjson)
	if err != nil {
		return err
	}
	type RESP struct {
		MsgGroup string `json:"msgGroup"`
		Rspcod   string `json:"rspcod"`
		Success  bool   `json:"success"`
	}
	var resp RESP
	err = res.ToJSON(&resp)
	if err != nil {
		return err
	}
	if resp.Success {
		return nil
	}
	log.Println(resp.Rspcod)
	return errors.New(resp.Rspcod)
}
