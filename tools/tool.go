package tools

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/yinrenxin/hgoj/setting"
)

var CONTEST_PRO_KEY = map[int]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "I",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
	16: "Q",
	17: "R",
	18: "S",
	19: "T",
	20: "U",
	21: "V",
	22: "W",
	23: "X",
	24: "Y",
	25: "Z",
}

var LANGUAGE_MAP = map[int]string{
	1:  "C/C++",
	3:  "Java",
	6:  "Python",
	17: "Go",
}

func InitTools() {
	downDir := "./static/down"
	zipDir := "/home/judge/data/tempzip"
	err := os.Mkdir(downDir, os.ModePerm)
	if err != nil {
		logs.Error(err)
	}
	err2 := os.Mkdir(zipDir, os.ModePerm)
	if err2 != nil {
		logs.Error(err)
	}
}

func StringToInt32(s string) (int32, error) {
	id, err := strconv.ParseInt(s, 10, 32)
	id64 := int32(id)
	if err != nil {
		return -1, err
	}
	return id64, nil
}

func StringToInt(s string) int {
	id, err := strconv.ParseInt(s, 10, 32)
	id64 := int(id)
	if err != nil {
		return -1
	}
	return id64
}

func StringToMonth(s string) time.Month {
	id, err := strconv.ParseInt(s, 10, 32)
	id64 := time.Month(id)
	if err != nil {
		return -1
	}
	return id64
}

func IntToString(s int) string {
	return strconv.Itoa(s)
}

func SplitIP(ip string) string {
	temp := strings.Split(ip, ":")

	return temp[0]
}

func CheckEmail(email string) (b bool) {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !m {
		return false
	}
	return true
}

func MD5(s string) string {
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
	url = setting.GRAVATARSOURCE + HashEmail(email) + "?d=identicon&s=" + s

	return url
}

func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
