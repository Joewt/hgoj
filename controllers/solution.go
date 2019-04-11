package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/syserror"
)

type SolutionController struct {
	BaseController
}

// @router /submit [post]
func (this *SolutionController) Submit() {
	source, code_length := this.FilterSource("source", "代码不能为空")
	proId := this.GetString("proid")
	uid := "1"

	logs.Info(proId, source)
	sid, err := models.AddSolution(proId, source, uid, code_length)
	logs.Info("solutionid ：", sid, "err:", err)
	if err != nil {
		this.JsonErr("保存代码错误", syserror.SAVE_CODE_ERR, "problem")
	}
	this.JsonOK("提交成功","/status")
}