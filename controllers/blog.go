package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/models"
)

type BlogController struct {
	BaseController
}


// @router /article/index [get]
func (this *BlogController) BlogIndex() {
	this.TplName = "blog/index.html"
}


// @router /admin/art/add [get]
func (this *BlogController) BlogAddGet() {
	this.TplName = "admin/addArt.html"
}


// @router /admin/art/add [post]
func (this *BlogController) BlogAddPost() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	title := this.GetMushString("title", "标题不能为空")
	content := this.GetMushString("content", "内容不能为空")

	_, err := models.AddArt(this.User.UserId,title,content)
	if err != nil {
		this.JsonErr("文章添加错误",10000,"")
	}

	this.JsonOK("文章添加成功","/admin/art/list")
}



// @router /admin/art/list [get]
func (this *BlogController) BlogList() {
	Art, err := models.QueryAllArt()
	if err != nil {
		logs.Error("文章查询错误",err)
	}
	this.Data["art"] = Art
	this.TplName = "admin/listArt.html"
}