package controllers


type UserController struct {
	BaseController
}

// @router /profile/:uid [get]
func (this *UserController) Profile() {
	this.TplName = "profile.html"
}


// @router /user/reg [post]
func (this *UserController) UserReg() {

	this.JsonErr("报错了", 1000,"/reg")
}