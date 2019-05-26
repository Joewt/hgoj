package controllers


// @router /problem/editdata/:pid [get]
func (this *ProblemController) ProblemTestDataEdit() {
	this.TplName = "admin/editData.html"
}