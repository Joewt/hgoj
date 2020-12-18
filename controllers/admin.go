package controllers

import (
	"html/template"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
)

// @router /admin/rejudge [get]
func (this *ProblemController) Rejudge() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/rejudge.html"
}

// @router /admin/rejudge [post]
func (this *ProblemController) RejudgePost() {
	rjpid := this.GetString("rjpid")
	rjcid := this.GetString("rjcid")
	rjsid := this.GetString("rjsid")
	if rjpid != "" {
		pid, _ := tools.StringToInt32(rjpid)
		_, err := models.QueryProblemById(pid)
		if err == orm.ErrNoRows {
			this.JsonErr("找不到对应的数据", 13001, "")
		}
		if ok := models.UpdateSolutionResultByPid(pid); !ok {
			this.JsonErr("重判失败", 13002, "")
		}
		this.JsonOK("题目重判成功", "/status")

	}
	if rjcid != "" {
		cid, _ := tools.StringToInt32(rjcid)
		_, err := models.QueryContestByConId(cid)
		if err == orm.ErrNoRows {
			this.JsonErr("找不到对应的数据", 13001, "")
		}
		if ok := models.UpdateSolutionResultByCid(cid); !ok {
			this.JsonErr("重判失败", 13002, "")
		}
		this.JsonOK("比赛重判成功", "/contest/status/cid/"+rjcid)
	}
	if rjsid != "" {
		sid, _ := tools.StringToInt32(rjsid)
		_, err := models.QuerySolutionBySid(sid)
		if err == orm.ErrNoRows {
			this.JsonErr("找不到对应的数据", 13001, "")
		}
		if ok := models.UpdateSolutionResultBySid(sid); !ok {
			this.JsonErr("重判失败", 13002, "")
		}
		this.JsonOK("提交重判成功", "/status")
	}
	this.JsonErr("未知错误", 13000, "")
}
