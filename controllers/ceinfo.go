package controllers

type CeinfoController struct {
	BaseController
}


// @router /ceinfo/:k [get]
func (this *CeinfoController) Ceinfo() {
	//todo 验证session
	this.TplName = "ceinfo.html"
}
