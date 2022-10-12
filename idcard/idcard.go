package idcard

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type IDCardInfo struct {
	IDCardNo string
}

//实例化居民身份证结构体
func NewIDCard(IDCardNo string) *IDCardInfo {
	return &IDCardInfo{IDCardNo: IDCardNo}
}

func (s *IDCardInfo) Verification() bool {
	var (
		coefficient []int32 = []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		code        []byte  = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	)
	if len(s.IDCardNo) != 18 {
		return false
	}
	idByte := []byte(strings.ToUpper(s.IDCardNo))
	sum := int32(0)
	for i := 0; i < 17; i++ {
		sum += int32(byte(idByte[i])-byte('0')) * coefficient[i]
	}
	return code[sum%11] == idByte[17]
}

//根据身份证号获取生日（时间格式）
func (s *IDCardInfo) GetBirthDay() *time.Time {
	if s == nil {
		return nil
	}
	dayStr := s.IDCardNo[6:14] + "000001"
	birthDay, err := time.Parse("20060102150405", dayStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &birthDay
}

//根据身份证号获取生日（字符串格式 yyyy-MM-dd）
func (s *IDCardInfo) GetBirthDayStr() string {
	defaultDate := "1999-01-01 00:00:01"
	if s == nil {
		return defaultDate
	}

	birthDay := s.GetBirthDay()
	if birthDay == nil {
		return defaultDate
	}
	return birthDay.Format("2006-01-02")
}

//根据身份证号获取生日的年份
func (s *IDCardInfo) GetYear() string {
	if s == nil {
		return ""
	}
	return s.IDCardNo[6:10]
}

//根据身份证号获取生日的月份
func (s *IDCardInfo) GetMonth() string {
	if s == nil {
		return ""
	}
	return s.IDCardNo[10:12]
}

//根据身份证号获取生日的日份
func (s *IDCardInfo) GetDay() string {
	if s == nil {
		return ""
	}
	return s.IDCardNo[12:14]
}

//根据身份证号获取性别
func (s *IDCardInfo) GetSex() string {
	if s == nil {
		return "证件号错误"
	}
	sexStr := s.IDCardNo[16:17]
	if sexStr == "" {
		return "证件号错误"
	}
	i, err := strconv.Atoi(sexStr)
	if err != nil {
		return "证件号错误"
	}
	if i%2 != 0 {
		return "男"
	}
	return "女"
}

func (s *IDCardInfo) GetAge() int {
	if s == nil {
		return 19
	}
	birthDay := s.GetBirthDay()
	if birthDay == nil {
		return 19
	}
	now := time.Now()
	age := now.Year() - birthDay.Year()
	if now.Month() < birthDay.Month() {
		age = age - 1
	}
	if age <= 0 {
		return 19
	}
	if age <= 0 || age >= 150 {
		return 19
	}
	return int(age)
}
