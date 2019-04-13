package controllers


type BlogController struct {
	BaseController
}


// @router /article/index [get]
func (this *BlogController) BlogIndex() {
	this.TplName = "blog/index.html"
}


// @router /article/add [get]
func (this *BlogController) BlogAddGet() {
	this.TplName = "admin/addContest.html"
}


// @router /article/add [post]
func (this *BlogController) ContestAdd() {
	this.TplName = "admin/addContest.html"
}


// @router /article/list [post]
func (this *BlogController) ContestList() {
	this.TplName = "admin/listContest.html"
}