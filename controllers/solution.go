package controllers

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/syserror"
	"github.com/yinrenxin/hgoj/tools"
	"time"
)

type SolutionController struct {
	BaseController
}

// @router /submit [post]
func (this *SolutionController) Submit() {
	if !this.IsLogin {
		this.Abort("500")
	}
	source, code_length := this.FilterSource("source", "代码不能为空")
	proId := this.GetString("proid")
	lang := this.GetString("language")
	conId, _ := tools.StringToInt32(this.GetString("conid"))
	req := this.Ctx.Request
	addr := req.RemoteAddr
	ip := tools.SplitIP(addr)
	uid := this.User.UserId


	if conId != 0 {
		con, err := models.QueryContestByConId(conId)
		if err != nil {
			this.JsonErr("系统错误", 9120,"")
		}
		t := time.Now().Sub(con.EndTime).Seconds()
		if t > 0 {
			this.JsonErr("比赛已结束",9121,"")
		}
	}

	_, err := models.AddSolution(proId, source, uid, code_length, lang, conId, ip)
	if err != nil {
		this.JsonErr("保存代码错误", syserror.SAVE_CODE_ERR, "problem")
	}
	this.JsonOK("提交成功","/status")
}


// @router /status/:page [get]
func (this *SolutionController) StatusPage() {
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo)*pageSize
	data,RESULT,_,totalNum,err := models.QueryPageSolution(start,pageSize)
	if err != nil {
		logs.Error(err)
	}
	isPage, pagePrev, pageNext := PageCal(totalNum,pageNo,pageSize)

	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["data"] = data
	this.Data["RES"] = RESULT
	this.TplName = "status.html"
}
