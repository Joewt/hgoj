package controllers

import (
	"html/template"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
)

type BlogController struct {
	BaseController
}

// @router /article/:artid [get]
func (this *BlogController) BlogIndex() {
	artid := this.Ctx.Input.Param(":artid")
	id, _ := tools.StringToInt32(artid)
	Art, err := models.QueryArtByArtId(id)
	if err != nil {
		this.Abort("500")
	}
	this.Data["art"] = Art
	this.TplName = "blog/index.html"
}

// @router /admin/art/add [get]
func (this *BlogController) BlogAddGet() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/addArt.html"
}

// @router /admin/art/add [post]
func (this *BlogController) BlogAddPost() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	title := this.GetMushString("title", "标题不能为空")
	content := this.GetMushString("content", "内容不能为空")

	_, err := models.AddArt(this.User.UserId, title, content)
	if err != nil {
		this.JsonErr("文章添加错误", 10000, "")
	}

	this.JsonOK("文章添加成功", "/admin/art/list")
}

// @router /admin/art/list [get]
func (this *BlogController) BlogList() {
	Art, err := models.QueryAllArt()
	if err != nil {
		logs.Error("文章查询错误", err)
	}
	this.Data["art"] = Art
	this.TplName = "admin/listArt.html"
}
