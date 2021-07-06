package main

import (
	"encoding/gob"
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
	"strconv"
	"strings"
	"time"

	//"strconv"
	_ "github.com/yinrenxin/hgoj/models"
	_ "github.com/yinrenxin/hgoj/routers"
	//_ "github.com/yinrenxin/hgoj/tools"
)

func main() {
	initTemplate()
	initSession()
	initStatic()
	initLogs()
	//启动定时任务
	go tools.StartCron()
	go tools.InitTools()
	beego.Run()
}


func initLogs() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/app.logs","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func initSession() {
	gob.Register(models.Users{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "hgoj"
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./data/session"
}


func initStatic() {
	beego.SetStaticPath("/","static")
	//beego.SetStaticPath("/problem", "static")
}


func initTemplate() {
	_ = beego.AddFuncMap("get_res", func(result int16) (string) {
		//res, _ := strconv.Atoi(result)
		return models.JUDGERES[int(result)]
	})


	_ = beego.AddFuncMap("get_res_class", func(result int16)(string){

		return models.JUDGERESCLSAA[int(result)]
	})

	_ = beego.AddFuncMap("menu_eq", func(a,b string)(string){
		if strings.Contains(a, b) {
			return "active"
		} else {
			return ""
		}
	})

	_ = beego.AddFuncMap("select_eq", func(a,b string)(string){
		if a == b {
			return "selected"
		} else {
			return ""
		}
	})

	_ = beego.AddFuncMap("cal_rate", func(a,b time.Time)(float64){
		startTime := a
		endTime := b
		totalTime := endTime.Sub(startTime).Minutes()
		t := time.Now().Sub(startTime).Minutes()
		value := t/totalTime
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
		percentage := value*100
		if t > totalTime {
			percentage = 100
		}

		if t < 0 {
			percentage = 0
		}
		return percentage
	})


	_ = beego.AddFuncMap("is_finish", func(a,b time.Time)(bool){
		startTime := a
		endTime := b
		totalTime := endTime.Sub(startTime).Minutes()
		t := time.Now().Sub(startTime).Minutes()
		if t > totalTime {
			return true
		}
		return false
	})

	_ = beego.AddFuncMap("RoleMap", func(role int32)(string){
		var temp string
		if role == 1 {
			temp = "管理员"
		}
		if role == 2 {
			temp = "教师"
		}
		if role == 0 {
			temp = "普通用户"
		}
		return temp
	})


	_ = beego.AddFuncMap("ParseTime", func(temp float64)(string){
		s := int(temp)
		h := s/(60*60)
		s = s - h*60*60
		m := s/60
		s = s - m*60
		th := strconv.Itoa(h)
		tm := strconv.Itoa(m)
		ts := strconv.Itoa(s)
		return th+":"+tm+":"+ts
	})

	_ = beego.AddFuncMap("Correctness_rate",func(i float64)(int){
		return int(i*100)
	})


	_ = beego.AddFuncMap("uidToUname",func(uid int32)(string){
		user,_ := models.QueryUserById(uid)
		return user.Nick
	})

	_ = beego.AddFuncMap("array_get", func(arr []int,i int)(int){
		return arr[i]
	})

	_ = beego.AddFuncMap("Language_map", func(i uint)(string){
		return tools.LANGUAGE_MAP[int(i)]
	})

	_ = beego.AddFuncMap("Format_time_to_d", func(t time.Time)(string){
		return t.Format("2006-01-02")
	})

	_ = beego.AddFuncMap("Format_time_to_s", func(t time.Time)(string){
		return t.Format("2006-01-02 15:04:05")
	})


}
