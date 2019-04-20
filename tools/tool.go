package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/yinrenxin/hgoj/setting"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StringToInt32(s string) (int32, error) {
	id , err := strconv.ParseInt(s,10,32)
	id64 := int32(id)
	if err != nil {
		return -1, err
	}
	return id64, nil
}


func StringToInt(s string) (int) {
	id , err := strconv.ParseInt(s,10,32)
	id64 := int(id)
	if err != nil {
		return -1
	}
	return id64
}



func StringToMonth(s string) (time.Month) {
	id , err := strconv.ParseInt(s,10,32)
	id64 := time.Month(id)
	if err != nil {
		return -1
	}
	return id64
}


func IntToString(s int) (string) {
	return strconv.Itoa(s)
}



func CheckEmail(email string) (b bool) {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !m {
		return false
	}
	return true
}


func MD5(s string)string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	var x string
	for _, v := range h.Sum(nil) {
		x += strconv.FormatInt(int64(v), 16)
	}
	return x
}



func HashEmail(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	h := md5.New()
	h.Write([]byte(email))
	return hex.EncodeToString(h.Sum(nil))
}
func AvatarLink(email string, size int) (url string) {
	s := IntToString(size)
	url = setting.GRAVATARSOURCE + HashEmail(email) + "?d=identicon&s="+s

	return url
}
