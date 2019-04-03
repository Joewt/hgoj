package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/yinrenxin/hgoj/models"
	_ "github.com/yinrenxin/hgoj/models"
	_ "github.com/yinrenxin/hgoj/routers"
)

func main() {
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