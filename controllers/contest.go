package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/syserror"
	"github.com/yinrenxin/hgoj/tools"
	"sort"
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
	Rank int
	Nick string
	AC int64
	Total int64
	TotalTime float64
	CP []CPProblem
}

type ContestProblem struct {
	ProId int32

}

type CR []*ContestRank


type CPProblem struct {
	ProId int32
	Flag bool
	ACtime float64
	ErrNum int64
}

// @router /contest/add [get]
func (this *ContestController) ContestAddGet() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	month := map[string]int{
		"January": 1,
		"February": 2,
		"March": 3,
		"April": 4,
		"May": 5,
		"June": 6,
		"July": 7,
		"August": 8,
		"September": 9,
		"October": 10,
		"November": 11,
		"December": 12,
	}
	tnow := time.Now()
	y := tnow.Year()
	m := tnow.Month().String()
	d := tnow.Day()
	this.Data["year"] = y
	this.Data["month"] = month[m]
	this.Data["day"] = d
	logs.Info(y,m,d)
	this.TplName = "admin/addContest.html"
}


// @router /contest/add [post]
func (this *ContestController) ContestAddPost() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	title := this.GetMushString("title", "标题不能为空")
	desc := this.GetMushString("desc", "竞赛描述不能为空")
	proIds := this.GetMushString("proIds", "题目编号不能为空")
	role := this.GetMushString("role", "权限不能为空")
	limituser := this.GetString("limituser")
	startTimeDate := this.GetMushString("starttime[0]","开始时间不能为空")
	startTimeH := this.GetMushString("starttime[1]","开始时间不能为空")
	startTimeM := this.GetMushString("starttime[2]","开始时间不能为空")
	startTimeSlice := strings.Split(startTimeDate,"-")
	endTimeDate := this.GetMushString("endtime[0]","结束时间不能为空")
	endTimeH := this.GetMushString("endtime[1]","结束时间不能为空")
	endTimeM := this.GetMushString("endtime[2]","结束时间不能为空")
	endTimeSlice := strings.Split(endTimeDate,"-")
	now := time.Time{}
	startTime := time.Date(tools.StringToInt(startTimeSlice[0]),tools.StringToMonth(startTimeSlice[1]),tools.StringToInt(startTimeSlice[2]),tools.StringToInt(startTimeH),tools.StringToInt(startTimeM),0,0,now.Location())
	endTime := time.Date(tools.StringToInt(endTimeSlice[0]),tools.StringToMonth(endTimeSlice[1]),tools.StringToInt(endTimeSlice[2]),tools.StringToInt(endTimeH),tools.StringToInt(endTimeM),0,0,now.Location())
	if endTime.Sub(startTime).Seconds() < 0 {
		this.JsonErr("时间错误",6010,"")
	}
	cid, err := models.ContestAdd(title, desc, proIds, role, limituser, startTime,endTime)

	if err != nil {
		this.JsonErr(err.Error(),6001, "/contest/add")
	}
	temp := strconv.Itoa(int(cid))
	this.JsonOK("添加比赛成功","/contest/cid/"+temp)
}



// @router /contest/update/:cid [get]
func (this *IndexController) ContestUpdate() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}

	cid := this.Ctx.Input.Param(":cid")
	id, _ := tools.StringToInt32(cid)

	con, err := models.QueryContestByConId(id)
	if err != nil {
		this.Abort("500")
	}


	pro,err := models.QueryProblemByCid(con.ContestId)
	if err != nil {
		this.Abort("500")
	}
	var proids string
	var tempstr []string
	for _, v := range pro{
		tempstr = append(tempstr,strconv.Itoa(int(v.ProblemId)))
	}

	proids = strings.Join(tempstr,",")
	month := map[string]int{
		"January": 1,
		"February": 2,
		"March": 3,
		"April": 4,
		"May": 5,
		"June": 6,
		"July": 7,
		"August": 8,
		"September": 9,
		"October": 10,
		"November": 11,
		"December": 12,
	}
	tstart := con.StartTime
	tend := con.EndTime
	sy := tstart.Year()
	sm := month[tstart.Month().String()]
	sd := tstart.Day()
	sh := tstart.Hour()
	smi := tstart.Minute()

	startTimes := []int{sy,sm,sd,sh,smi}


	ey := tend.Year()
	em := month[tend.Month().String()]
	ed := tend.Day()
	eh := tend.Hour()
	emi := tend.Minute()

	endTimes := []int{ey,em,ed,eh,emi}

	this.Data["starttime"] = startTimes
	this.Data["endtime"] = endTimes
	this.Data["con"] = con
	this.Data["pids"] = proids
	this.TplName = "admin/editContest.html"
}



// @router /contest/update [post]
func (this *ContestController) ContestUpdatePost() {
	id := this.GetMushString("cid", "error")
	cid,_ := tools.StringToInt32(id)
	title := this.GetMushString("title", "标题不能为空")
	desc := this.GetMushString("desc", "竞赛描述不能为空")
	proIds := this.GetMushString("proIds", "题目编号不能为空")
	role := this.GetMushString("role", "权限不能为空")
	limituser := this.GetString("limituser")
	startTimeDate := this.GetMushString("starttime[0]","开始时间不能为空")
	startTimeH := this.GetMushString("starttime[1]","开始时间不能为空")
	startTimeM := this.GetMushString("starttime[2]","开始时间不能为空")
	startTimeSlice := strings.Split(startTimeDate,"-")
	endTimeDate := this.GetMushString("endtime[0]","结束时间不能为空")
	endTimeH := this.GetMushString("endtime[1]","结束时间不能为空")
	endTimeM := this.GetMushString("endtime[2]","结束时间不能为空")
	endTimeSlice := strings.Split(endTimeDate,"-")
	now := time.Time{}
	startTime := time.Date(tools.StringToInt(startTimeSlice[0]),tools.StringToMonth(startTimeSlice[1]),tools.StringToInt(startTimeSlice[2]),tools.StringToInt(startTimeH),tools.StringToInt(startTimeM),0,0,now.Location())
	endTime := time.Date(tools.StringToInt(endTimeSlice[0]),tools.StringToMonth(endTimeSlice[1]),tools.StringToInt(endTimeSlice[2]),tools.StringToInt(endTimeH),tools.StringToInt(endTimeM),0,0,now.Location())
	if endTime.Sub(startTime).Seconds() < 0 {
		this.JsonErr("时间错误",6010,"")
	}
	_, err := models.ContestUpdate(cid,title, desc, proIds, role, limituser, startTime,endTime)

	if err != nil {
		this.JsonErr(err.Error(),6001, "/contest/add")
	}
	temp := strconv.Itoa(int(cid))
	this.JsonOK("更新比赛成功","/contest/cid/"+temp)
}


// @router /contest/:page [get]
func (this *ContestController) ContestPage() {
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo)*pageContestSize
	con,_,totalNum,_ := models.QueryPageContest(start,pageContestSize)

	isPage, pagePrev,pageNext := PageCal(totalNum,pageNo,pageContestSize)

	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["con"] = con
	this.TplName = "contest.html"
}


// @router /contest/cid/:id [get]
func (this *ContestController) ContestCid() {

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
		Pac,Psub := models.QueryACSUBFromSolutionBYPidCi(v.ProblemId,cid)
		pro = append(pro, Pro{ProblemId:v.ProblemId,Title:v.Title,Accepted:Pac,Submit:Psub,Solved:v.Solved,Cid:cid})
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
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	pageNo := 0
	start := int(pageNo)*pageSize
	con,num, totalNum,err := models.QueryPageContest(start,pageSize)
	if err != nil {
		logs.Error(err)
	}
	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	pagePrev := 1
	pageNext := 2
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["conNum"] = num
	this.Data["con"] = con
	this.TplName = "admin/listContest.html"
}


// @router /contest/list/:page [get]
func (this *ContestController) ContestListPage() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo)*pageSize
	con,num, totalNum,err := models.QueryPageContest(start,pageSize)
	if err != nil {
		logs.Error(err)
	}

	isPage, pagePrev, pageNext := PageCal(totalNum,pageNo,pageSize)

	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["conNum"] = num
	this.Data["con"] = con
	this.TplName = "admin/listContest.html"
}


// @router /contest/status/cid/:cid [get]
func (this *ContestController) ContestStatus() {
	cid := this.Ctx.Input.Param(":cid")
	id, _ := tools.StringToInt32(cid)
	pageNo := 0
	start := int(pageNo)*pageStatusSize
	data,RESULT,_,totalNum,err := models.QueryPageSolutionByCid(id,start,pageStatusSize)
	if err != nil {
		logs.Error(err)
	}

	isPage := true
	if int(totalNum) < pageStatusSize {
		isPage = false
	}
	pagePrev := 1
	pageNext := 2
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["conid"] = cid
	this.Data["data"] = data
	this.Data["RES"] = RESULT
	this.TplName = "contest/statusContest.html"
}


// @router /contest/status/cid/:cid/:page [get]
func (this *ContestController) ContestStatusPage() {
	cid := this.Ctx.Input.Param(":cid")
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo)*pageStatusSize
	id, _ := tools.StringToInt32(cid)
	data,RESULT,_,totalNum,err := models.QueryPageSolutionByCid(id,start,pageStatusSize)
	if err != nil {
		logs.Error(err)
	}

	isPage, pagePrev,pageNext := PageCal(totalNum,pageNo,pageStatusSize)

	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
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
	var data []*ContestRank

	for k, v := range uids {
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
		data = append(data, &ContestRank{k,nick,ac,total, TotalTime,CPData})
	}
	//对排名进行排序
	sort.Sort(CR(data))
	for k, v := range data {
		v.Rank = k+1
	}
	this.Data["proids"] = proIds
	this.Data["data"] = data
	this.Data["conid"] = cid
	this.Data["contest"] = contestInfo
	this.TplName = "contest/contestrank.html"
}


func (I CR) Len() int {
	return len(I)
}
func (I CR) Less(i, j int) bool {
	if I[i].AC == I[j].AC {
		return I[i].TotalTime < I[j].TotalTime
	}
	return I[i].AC > I[j].AC
}
func (I CR) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
