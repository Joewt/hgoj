package controllers

type IndexController struct {
	BaseController
}


// @router / [get]
func (this *IndexController) Index() {

	this.TplName = "index.html"
}

// @router /about [get]
func (this *IndexController) IndexAbout() {
	this.TplName = "about.html"
}


// @router /message [get]
func (this *IndexController) IndexMessage() {
	this.TplName = "message.html"
}

// @router /comment [get]
func (this *IndexController) IndexComment() {
	this.TplName = "comment.html"
}

// @router /login [get]
func (this *IndexController) IndexUser() {
	this.TplName = "login.html"
}

// @router /reg [get]
func (this *IndexController) IndexReg() {
	this.TplName = "reg.html"
}

// @router /admin [get]
func (this *IndexController) IndexAdmin() {
	this.TplName = "admin/index.html"
}

//// @router /create [get]
//func (this *IndexController) IndexCreate() {
//	if this.IsAdmin == false {
//		this.Abort401(syserror.New("权限不足",nil))
//	}
//	this.Data["key"] = this.UUID()
//	this.TplName = "editor.html"
//}