package controllers

import "github.com/astaxie/beego/logs"

type ProblemController struct {
	BaseController
}


// @router /problem/:id [get]
func (this *ProblemController) Problem() {
	id := this.Ctx.Input.Param(":id")
	logs.Info("id="+id)

	this.TplName = "problems.html"
}