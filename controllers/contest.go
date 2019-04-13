package controllers

type ContestController struct {
	BaseController
}



// @router /contest/add [get]
func (this *ContestController) ContestAddGet() {
	this.TplName = "admin/addContest.html"
}


// @router /contest/add [post]
func (this *ContestController) ContestAdd() {
	this.TplName = "admin/addContest.html"
}


// @router /contest/list [get]
func (this *ContestController) ContestList() {
	this.TplName = "admin/listContest.html"
}