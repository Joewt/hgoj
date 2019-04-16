package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/yinrenxin/hgoj/models"
	"time"

	//"strconv"
	_ "github.com/yinrenxin/hgoj/models"
	_ "github.com/yinrenxin/hgoj/routers"
)

func main() {
	initTemplate()
	initSession()
	initStatic()
	beego.Run()
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
		if a == b {
			return "active"
		} else {
			return ""
		}
	})
	_ = beego.AddFuncMap("cal_rate", func(a,b time.Time)(float64){
		startTime := a
		endTime := b
		totalTime := endTime.Sub(startTime).Minutes()
		t := time.Now().Sub(startTime).Minutes()
		percentage := (t/totalTime)*100
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
}