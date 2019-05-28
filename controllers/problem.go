package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/syserror"
	"io/ioutil"
	"os"
	"strconv"
	//"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
)

type ProblemController struct {
	BaseController
}



type Problems struct {
	num int64
	pageRange []int
	pageNum int
}


// @router /problemset/:page [get]
func (this *ProblemController) ProblemSetPage(){
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo)*pageSize
	pros,_,totalNum, err := models.QueryPageProblem(start, pageSize)
	if err != nil {
		logs.Error(err)
	}
	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	temp := int(totalNum) / pageSize
	var t  []int
	for i := 0; i <= temp;i ++ {
		t = append(t, i+1)
	}
	proData := new(Problems)
	proData.pageRange = t
	proData.num = totalNum
	pageRange := t
	pagePrev := pageNo
	pageNext := pageNo + 2
	if int(pageNo) == temp {
		pageNext = pageNo + 1
	}
	if pageNo == 0 {
		pagePrev = pageNo + 1
	}


	this.Data["pageData"] = proData
	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["problems"] = pros
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.TplName = "problem.html"
}


// @router /problem/:id [get]
func (this *ProblemController) Problem() {
	id := this.Ctx.Input.Param(":id")
	ids , err := tools.StringToInt32(id)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}
	pro,err := models.QueryProblemById(ids)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}

	this.Data["problem"] = pro
	this.TplName = "problems.html"
}


// @router /problem/edit/:id [get]
func (this *ProblemController) ProblemEdit() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	id := this.Ctx.Input.Param(":id")
	ids , err := tools.StringToInt32(id)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}
	pro,err := models.QueryProblemById(ids)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}

	//临时处理(读测试数据并展示)
	baseDir := OJ_DATA+"/"+strconv.Itoa(int(ids))

	testin := baseDir+"/test.in"
	testout := baseDir+"/test.out"

	in, err := ioutil.ReadFile(testin)
	if err != nil {
		this.Abort("401")
	}

	out, err := ioutil.ReadFile(testout)
	if err != nil {
		this.Abort("401")
	}

	pro.Input = string(in)
	pro.Output = string(out)


	this.Data["PRO"] = pro
	this.TplName = "admin/editProblem.html"
}


// @router /problem/update [post]
func (this *ProblemController) ProblemUpdate() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	proId := this.GetString("proid")
	temp ,err := strconv.Atoi(proId)
	id := int32(temp)
	//refer := this.Ctx.Request.Referer()
	//re := regexp.MustCompile("")
	//logs.Info(re.FindString(refer))
	pro,err := models.QueryProblemById(id)
	if err != nil {
		this.JsonErr("问题未找到", syserror.PROBLEM_NOT_FOUND, "/problem/edit/"+proId)
		logs.Error(pro,err)
	}

	title := this.GetMushString("title", "标题不能为空")
	memory := this.GetMushString("memory", "限制内存不能为空")
	time := this.GetMushString("time", "限制时间不能为空")
	desc := this.GetMushString("desc", "描述不能为空")
	input := this.GetMushString("input", "input不能为空")
	output := this.GetMushString("output", "output不能为空")
	sampleinput := this.GetMushString("sampleinput", "sampleinput不能为空")
	sampleoutput := this.GetMushString("sampleoutput", "sampleoutput不能为空")
	testinput := this.GetMushString("testinput", "testinput不能为空")
	testoutput := this.GetMushString("testoutput", "testoutput不能为空")
	data := []string{title,memory,time,desc,input,output,sampleinput,sampleoutput}
	ok, err := models.UpdateProblemById(id,data)
	if !ok {
		this.JsonErr("更新失败", syserror.UPDATE_PROBLEM_ERR, "problem/edit/"+proId)
	}

	pid := int64(id)

	ok = mkdata(pid, "test.in", testinput,OJ_DATA)
	if !ok {
		this.JsonErr("syserror", syserror.FILE_WRITE_ERR,"/problem/add")
	}
	ok = mkdata(pid, "test.out", testoutput, OJ_DATA)
	if !ok {
		this.JsonErr("syserror", syserror.FILE_WRITE_ERR,"/problem/add")
	}
	this.JsonOK("更新成功", "/problem/list")
}


// @router /problem/add [get]
func (this *ProblemController) ProblemAdd() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	this.TplName = "admin/addProblem.html"
}


// @router /problem/del [post]
func (this *ProblemController) ProblemDel() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	proId := this.GetString("proid")
	logs.Info("要删除的proId：",proId)
	temp, _ := strconv.Atoi(proId)
	id := int32(temp)
	ok := models.DelProblemById(id)
	if !ok {
		this.JsonErr("删除失败", syserror.DEL_PROBLEM_ERR, "/problem/list")
	}
	this.JsonOK("删除成功", "/problem/list")
}

// @router /problem/list [get]
func (this *ProblemController) ProblemList() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	pageNo := 0
	start := pageNo*pageSize
	pros,_,totalNum,err := models.QueryPageProblem(start,pageSize)
	if err != nil {
		logs.Error(err)
	}

	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	temp := int(totalNum) / pageSize
	var t  []int
	for i := 0; i <= temp;i ++ {
		t = append(t, i+1)
	}
	pageRange := t
	pagePrev := pageNo + 1
	pageNext := pageNo + 2


	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["PROS"] = pros
	this.TplName = "admin/listProblem.html"
}


// @router /problem/list/:page [get]
func (this *ProblemController) ProblemListPage() {
	if !this.IsAdmin && !this.IsTeacher{
		this.Abort("401")
	}
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo)*pageSize
	pros,_,totalNum,err := models.QueryPageProblem(start,pageSize)
	if err != nil {
		logs.Error(err)
	}

	isPage,pageRange,pagePrev,pageNext := PageRangeCal(totalNum,pageNo,pageSize)

	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["PROS"] = pros
	this.TplName = "admin/listProblem.html"
}


// @router /problem/add [post]
func (this *ProblemController) ProblemAddPost() {
	title := this.GetMushString("title", "标题不能为空")
	memory := this.GetMushString("memory", "限制内存不能为空")
	time := this.GetMushString("time", "限制时间不能为空")
	desc := this.GetMushString("desc", "描述不能为空")
	input := this.GetMushString("input", "input不能为空")
	output := this.GetMushString("output", "output不能为空")
	sampleinput := this.GetMushString("sampleinput", "sampleinput不能为空")
	sampleoutput := this.GetMushString("sampleoutput", "sampleoutput不能为空")
	testinput := this.GetMushString("testinput", "testinput不能为空")
	testoutput := this.GetMushString("testoutput", "testoutput不能为空")

	pid, err := models.AddProblem(title,time,memory,desc,input,output,sampleinput,sampleoutput)
	if err != nil {
		this.JsonErr("更新失败", syserror.ADD_PROBLEM_ERR, "/problem/add")
	}
	ok := mkdata(pid, "test.in", testinput,OJ_DATA)
	if !ok {
		this.JsonErr("写入文件权限不足", syserror.FILE_WRITE_ERR,"/problem/add")
	}
	ok = mkdata(pid, "test.out", testoutput, OJ_DATA)
	if !ok {
		this.JsonErr("写入文件权限不足", syserror.FILE_WRITE_ERR,"/problem/add")
	}

	this.JsonOK("添加题目成功", "/admin")
}


// @router /problem/fileupload [post]
func (this *ProblemController) Fileupload() {
	logs.Info(this.Ctx.Request)
	f, h, err := this.GetFile("testcase")
	logs.Info(h.Filename)
	if err != nil {
		logs.Error("error:--- ",err)
	}
	defer f.Close()
	this.SaveToFile("testcase", "/Users/joe/go/src/github.com/yinrenxin/hgoj/static/upload/" + h.Filename)
}

func mkdata(pid int64, filename string, input string, oj_data string) bool {
	baseDir := oj_data+"/"+strconv.Itoa(int(pid))
	err := os.MkdirAll(baseDir, 0777)
	if err != nil {
		logs.Error("目录创建失败",err)
		return false
	}
	name := baseDir+"/"+filename
	logs.Info(name)
	data := []byte(input)
	if ioutil.WriteFile(name,data,0644) != nil {
		logs.Error("文件写入失败,文件名",name)
		return false
	}
	return true
}
