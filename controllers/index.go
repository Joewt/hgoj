package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
)

type IndexController struct {
	BaseController
}


// @router /index [get]
func (this *IndexController) Index() {
	logs.Info("xxxx")
	this.TplName = "index.html"
}

// @router /faqs [get]
func (this *IndexController) IndexFaqs() {
	this.TplName = "faqs.html"
}


// @router /problemset [get]
func (this *IndexController) IndexProblemset() {

	pros,num,err := models.QueryAllProblem()
	if err != nil {
		logs.Error(err)
	}

	logs.Info("pros:---",pros,num)

	this.Data["problems"] = pros
	this.TplName = "problem.html"
}

// @router /status [get]
func (this *IndexController) IndexStatus() {
	data,RESULT,err := models.QueryAllSolution()
	if err != nil {
		logs.Error(err)
	}
	this.Data["data"] = data
	this.Data["RES"] = RESULT
	this.Data["m"] = map[string]interface{} {
		"a": 1,
		"1": map[string]float64{
			"c": 4,
		},
	}
	this.TplName = "status.html"
}

// @router /contest [get]
func (this *IndexController) IndexContest() {
	this.TplName = "contest.html"
}

// @router /login [get]
func (this *IndexController) IndexUser() {
	this.TplName = "login.html"
}

// @router /reg [get]
func (this *IndexController) IndexReg() {
	this.TplName = "reg.html"
}

// @router /admin [get]
func (this *IndexController) IndexAdmin() {
	this.TplName = "admin/index.html"
}

// @router /skin-config.html [get]
func (this *IndexController) IndexSkinConfig() {
	this.TplName = "admin/skin-conf.html"
}

//// @router /create [get]
//func (this *IndexController) IndexCreate() {
//	if this.IsAdmin == false {
//		this.Abort401(syserror.New("权限不足",nil))
//	}
//	this.Data["key"] = this.UUID()
//	this.TplName = "editor.html"
//}