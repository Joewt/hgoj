package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
	"strconv"
	"strings"
	"time"
)

type ContestController struct {
	BaseController
	Visible bool
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


// @router /contest/cid/:id [get]
func (this *ContestController) ContestCid() {
	//cid := this.Ctx.Input.Param("id")
	this.Visible = true
	req :=  this.Ctx.Request.RequestURI
	temp := strings.Split(req,"/")
	cid, _ := tools.StringToInt32(temp[len(temp)-1])

	con, err := models.QueryContestByConId(cid)
	if err != nil {
		this.Abort("500")
	}

	//根据cid查找problem
	pro, err := models.QueryProblemByCid(con.ContestId)

	if err != nil {
		this.Visible = false
	}

	//进度条处理
	startTime := con.StartTime
	endTime := con.EndTime
	totalTime := endTime.Sub(startTime).Minutes()
	t := time.Now().Sub(startTime).Minutes()
	percentage := (t/totalTime)*100
	if t > totalTime {
		percentage = 100
	}
	if t < 0 {
		this.Visible = false
		percentage = 0
	}
	logs.Info("题目是否可见",this.Visible)
	this.Data["con"] = con
	this.Data["percent"] = percentage
	this.Data["pro"] = pro
	this.Data["visible"] = this.Visible
	this.TplName = "contest/indexContest.html"
}


// @router /contest/list [get]
func (this *ContestController) ContestList() {

	this.TplName = "admin/listContest.html"
}
