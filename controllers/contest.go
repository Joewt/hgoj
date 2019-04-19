package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/syserror"
	"github.com/yinrenxin/hgoj/tools"
	"strconv"
	"strings"
	"time"
)

type ContestController struct {
	BaseController
	Visible bool
}


type Pro struct {
	ProblemId		int32
	Title			string
	Accepted		int32
	Submit			int32
	Solved			int32
	Cid 			int32
}

type ContestRank struct {
	Nick string
	AC int64
	Total int64
	TotalTime float64
	CP []CPProblem
}

type ContestProblem struct {
	ProId int32

}


type CPProblem struct {
	ProId int32
	Flag bool
	ACtime float64
	ErrNum int64
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
	if !this.IsLogin {
		this.Abort401(syserror.UnKnowError{})
	}
	this.Visible = true
	req :=  this.Ctx.Request.RequestURI
	temp := strings.Split(req,"/")
	cid, _ := tools.StringToInt32(temp[len(temp)-1])

	con, err := models.QueryContestByConId(cid)
	if err != nil {
		this.Abort("500")
	}

	//根据cid查找problem
	pros, err := models.QueryProblemByCid(con.ContestId)

	if err != nil {
		this.Visible = false
	}


	var pro []Pro

	for _, v := range pros{
		pro = append(pro, Pro{ProblemId:v.ProblemId,Title:v.Title,Accepted:v.Accepted,Submit:v.Submit,Solved:v.Solved,Cid:cid})
	}

	//根据cid查找 ac数
	ac, sub := models.QueryACNumContestByCid(cid)

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
	this.Data["conid"] = cid
	this.Data["con"] = con
	this.Data["percent"] = percentage
	this.Data["problem"] = pro
	this.Data["visible"] = this.Visible
	this.Data["accepted"] = ac
	this.Data["submit"] = sub
	this.TplName = "contest/indexContest.html"
}


// @router /contest/problem/:id/:cid [get]
func (this *ProblemController) ProblemContest() {
	id := this.Ctx.Input.Param(":id")
	cid := this.Ctx.Input.Param(":cid")
	ids , err := tools.StringToInt32(id)
	if err != nil {
		this.Abort401(err)
		logs.Error(err)
	}
	pro,err := models.QueryProblemById(ids)
	if err != nil {
		this.Abort401(err)
		logs.Error(err)
	}

	c, _ := tools.StringToInt32(cid)
	pros, _ := models.QueryProblemByCid(c)
	f := false
	for _, v := range pros {
		if v.ProblemId == ids {
			f = true
		}
	}
	if !f {
		this.Abort404(syserror.UnKnowError{})
	}

	this.Data["conid"] = cid
	this.Data["problem"] = pro
	this.TplName = "contest/proContest.html"
}


// @router /contest/list [get]
func (this *ContestController) ContestList() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	con,num, err := models.QueryAllContest()
	if err != nil {
		logs.Error(err)
	}
	this.Data["conNum"] = num
	this.Data["con"] = con
	this.TplName = "admin/listContest.html"
}

// @router /contest/status/cid/:cid [get]
func (this *ContestController) ContestStatus() {
	cid := this.Ctx.Input.Param(":cid")
	id, _ := tools.StringToInt32(cid)
	data,RESULT,err := models.QueryAllSolutionByCid(id)
	if err != nil {
		logs.Error(err)
	}
	this.Data["conid"] = cid
	this.Data["data"] = data
	this.Data["RES"] = RESULT
	this.TplName = "contest/statusContest.html"
}


// @router /contestrank/cid/:cid [get]
func (this *ContestController) ContestRank() {
	cid := this.Ctx.Input.Param(":cid")
	c, _ := tools.StringToInt32(cid)
	pros, _ := models.QueryProblemByCid(c)

	contestInfo,_ := models.QueryContestByConId(c)

	var proIds []ContestProblem
	for _, v := range pros {
		//logs.Info(k,v.ProblemId)
		proIds = append(proIds, ContestProblem{v.ProblemId})
	}
	uids, _ := models.QueryAllUserIdByCid(c)
	var data []ContestRank

	for _, v := range uids {
		nick, ac, total := models.QueryACNickTotalByUid(v, c)
		var CPData []CPProblem
		for _, p := range proIds {
			qpid,flag,actime,ErrNum := models.QueryJudgeTimeFromSolutionByUidCidPid(v,p.ProId,c,contestInfo.StartTime)
				CPData = append(CPData, CPProblem{qpid, flag, actime, ErrNum})
		}
		var TotalTime float64
		for _, TT := range CPData {
			TotalTime += TT.ACtime
		}
		data = append(data, ContestRank{nick,ac,total, TotalTime,CPData})
	}

	//对排名进行排序

	this.Data["proids"] = proIds
	this.Data["data"] = data
	this.Data["conid"] = cid
	this.TplName = "contest/contestrank.html"
}