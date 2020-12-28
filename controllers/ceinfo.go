package controllers

import (
	"github.com/beego/beego/v2/adapter/logs"
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


	compileinfo, _ := models.QueryCompileInfoBySid(id)


	this.Data["username"] = user.UserName
	this.Data["solu"] = solution
	this.Data["source"] = source
	this.Data["compileinfo"] = compileinfo
	this.TplName = "ceinfo.html"
}


// @router /contest/ceinfo/:sid [get]
func (this *CeinfoController) CeinfoContest() {
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

	compileinfo, _ := models.QueryCompileInfoBySid(id)


	this.Data["username"] = user.UserName
	this.Data["solu"] = solution
	this.Data["source"] = source
	this.Data["conid"] = solution.ContestId
	this.Data["compileinfo"] = compileinfo
	this.TplName = "contest/ceinfo.html"
}
