package routers

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/yinrenxin/hgoj/controllers"
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
	beego.Router("/friend", &controllers.FriendController{})

}
