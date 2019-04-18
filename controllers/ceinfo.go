package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
)

type CeinfoController struct {
	BaseController
}


// @router /ceinfo/:sid [get]
func (this *CeinfoController) Ceinfo() {
	//todo 验证session
	if !this.IsLogin {
		this.Abort("401")
	}

	sid := this.Ctx.Input.Param(":sid")
	id, _ := tools.StringToInt32(sid)
	source := models.QuerySourceBySolutionId(id)


	solution, _ := models.QuerySolutionBySid(id)

	user, err := models.QueryUserById(solution.UserId)
	if err != nil {
		logs.Error(err)
	}


	if solution.UserId != this.User.UserId && !this.IsAdmin {
		this.Abort("401")
	}


	this.Data["username"] = user.UserName
	this.Data["solu"] = solution
	this.Data["source"] = source
	this.TplName = "ceinfo.html"
}
