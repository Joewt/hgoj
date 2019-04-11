package tools

import (
	"crypto/md5"
	"io"
	"strconv"
	"regexp"
)

func StringToInt32(s string) (int32, error) {
	id , err := strconv.ParseInt(s,10,32)
	id64 := int32(id)
	if err != nil {
		return -1, err
	}
	return id64, nil
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

