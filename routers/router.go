package routers

import (
	"github.com/yinrenxin/hgoj/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.ProblemController{})
	beego.Include(&controllers.CeinfoController{})
	beego.Include(&controllers.SolutionController{})

}
