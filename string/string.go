package string

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func ParseInt(s string) int {
	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func ParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func ParseBool(s string) bool {
	if s == "\x01" || s == "true" {
		return true
	} else if s == "false" {
		return false
	}

	i := ParseInt(s)
	return i != 0
}

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// CamelToSnakeCase This function transform camelcase in snakecase LoremIpsum in lorem_ipsum
func CamelToSnakeCase(camel string) string {
	var buf bytes.Buffer
	for _, c := range camel {
		if 'A' <= c && c <= 'Z' {
			// just convert [A-Z] to _[a-z]
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c - 'A' + 'a')
			continue
		}
		buf.WriteRune(c)
	}
	return strings.ReplaceAll(buf.String(), " ", "")
}

func GetRandomName() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func GetMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func IsStringsEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func GetMaxLenStr(strs ...string) string {
	m := 0
	i := 0
	for j, str := range strs {
		l := len(str)
		if l > m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func GetMinLenStr(strs ...string) string {
	m := int(^uint(0) >> 1)
	i := 0
	for j, str := range strs {
		l := len(str)
		if l < m {
			m = l
			i = j
		}
	}
	return strs[i]
}

func ReadStringFromPath(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func WriteStringToPath(s string, path string) {
	err := os.WriteFile(path, []byte(s), 0o644)
	if err != nil {
		panic(err)
	}
}

// SnakeString transform XxYy to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	result := strings.ToLower(string(data[:]))
	return strings.ReplaceAll(result, " ", "")
}

func IsChinese(str string) bool {
	var flag bool
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			flag = true
			break
		}
	}
	return flag
}

func GetMaskedPhone(phone string) string {
	re := regexp.MustCompile(`(\d{3})\d+(\d{4})`)
	return re.ReplaceAllString(phone, "$1****$2")
}

func GetMaskedEmail(email string) string {
	if email == "" {
		return ""
	}

	tokens := strings.Split(email, "@")
	username := maskString(tokens[0])
	domain := tokens[1]
	domainTokens := strings.Split(domain, ".")
	domainTokens[len(domainTokens)-2] = maskString(domainTokens[len(domainTokens)-2])
	return fmt.Sprintf("%s@%s", username, strings.Join(domainTokens, "."))
}

func maskString(str string) string {
	if len(str) <= 2 {
		return str
	} else {
		return fmt.Sprintf("%c%s%c", str[0], strings.Repeat("*", len(str)-2), str[len(str)-1])
	}
}
