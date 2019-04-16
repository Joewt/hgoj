package routers

import (
	"github.com/yinrenxin/hgoj/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.ProblemController{})
	beego.Include(&controllers.CeinfoController{})
	beego.Include(&controllers.SolutionController{})
	beego.Include(&controllers.ContestController{})
	beego.Include(&controllers.BlogController{})

}
