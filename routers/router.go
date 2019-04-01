package routers

import (
	"github.com/yinrenxin/hgoj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
