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


// @router /problem/add [get]
func (this *ProblemController) ProblemAdd() {
	this.TplName = "admin/addProblem.html"
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
		this.JsonErr("更新失败", 1001, "/problem/add")
	}
	ok := mkdata(pid, "test.in", testinput,OJ_DATA)
	if !ok {
		this.JsonErr("syserror", syserror.FILE_WRITE_ERR,"/problem/add")
	}
	ok = mkdata(pid, "test.out", testoutput, OJ_DATA)
	if !ok {
		this.JsonErr("syserror", syserror.FILE_WRITE_ERR,"/problem/add")
	}
	//summary, err := getSummary(content)
	//if err != nil {
	//	summary = ""
	//}
	//u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	//err = models.InsertArticle(key, title, content,u.Id, summary)
	//if err != nil {
	//	this.Abort500(syserror.New("系统错误",err))
	//}
	this.JsonOK("添加题目成功", "/admin")
	//this.JsonOK("添加文章成功","/")
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
