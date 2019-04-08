package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/syserror"

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
	logs.Info(title)
	memory := this.GetMushString("memory", "限制内存不能为空")
	logs.Info(memory)
	//time := this.GetMushString("time", "限制时间不能为空")
	//desc := this.GetMushString("desc", "描述不能为空")
	//input := this.GetMushString("input", "input不能为空")
	//output := this.GetMushString("output", "output不能为空")
	//sampleinput := this.GetMushString("sampleinput", "sampleinput不能为空")
	//sampleoutput := this.GetMushString("sampleoutput", "sampleoutput不能为空")
	//testinput := this.GetMushString("testinput", "testinput不能为空")
	//testoutput := this.GetMushString("testoutput", "testoutput不能为空")
	//
	//
	//
	//
	//
	//logs.Info(title,memory,time,desc,input,output,sampleinput,sampleoutput,testinput,testoutput)
	////
	//_, ok := models.QueryByKeyArt(key)
	//if ok == true {
	//	this.Abort500(syserror.New("已经有该文章了，请勿重复添加",nil))
	//}
	//summary, err := getSummary(content)
	//if err != nil {
	//	summary = ""
	//}
	//u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	//err = models.InsertArticle(key, title, content,u.Id, summary)
	//if err != nil {
	//	this.Abort500(syserror.New("系统错误",err))
	//}
	this.Abort500(syserror.New("系统错误",nil))
	//this.JsonOK("添加文章成功","/")
}
