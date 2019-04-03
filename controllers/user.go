package controllers


type UserController struct {
	BaseController
}

// @router /profile/:uid [get]
func (this *UserController) Profile() {
	this.TplName = "profile.html"
}