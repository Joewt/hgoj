package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"strconv"
	"time"
)

type ContestController struct {
	BaseController
}



// @router /contest/add [get]
func (this *ContestController) ContestAddGet() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	this.TplName = "admin/addContest.html"
}


// @router /contest/add [post]
func (this *ContestController) ContestAddPost() {
	title := this.GetMushString("title", "标题不能为空")
	desc := this.GetMushString("desc", "竞赛描述不能为空")
	proIds := this.GetMushString("proIds", "题目编号不能为空")
	role := this.GetMushString("role", "权限不能为空")
	limituser := this.GetString("limituser")

	now := time.Time{}
	startTime := time.Date(2019,1,1,9,00,0,0,now.Location())
	endTime := time.Date(2019,1,1,14,00,0,0,now.Location())


	cid, err := models.ContestAdd(title, desc, proIds, role, limituser, startTime,endTime)

	if err != nil {
		this.JsonErr(err.Error(),6001, "/contest/add")
	}
	temp := strconv.Itoa(int(cid))
	this.JsonOK("添加比赛成功","/contest/cid/"+temp)
}


// @router /contest/cid/:cid [get]
func (this *ContestController) ContestCid() {
	cid := this.Ctx.Input.Param("cid")
	logs.Info("cid: ",cid)
	this.TplName = "contest/indexContest.html"
}


// @router /contest/list [get]
func (this *ContestController) ContestList() {
	this.TplName = "admin/listContest.html"
}
