package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/tools"
	"github.com/yinrenxin/hgoj/models"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /profile/:uid [get]
func (this *UserController) Profile() {
	this.TplName = "profile.html"
}


// @router /user/reg [post]
func (this *UserController) UserReg() {
	username := this.GetMushString("username", "用户名不能为空")
	nick := this.GetMushString("nick", "昵称不能为空")
	email := this.GetMushString("email","邮箱不能为空")
	pwd   := this.GetMushString("pwd", "密码不能为空")
	pwd1  := this.GetMushString("pwd2", "确认密码不能为空")
	school := this.GetMushString("school", "学校不能为空")
	logs.Info(nick, school)
	//获取客户端ip
	ip := this.Ctx.Request.RemoteAddr

	if tools.CheckEmail(email) == false {
		this.JsonErr("邮箱格式错误", 1102, "/reg")
	}

	if strings.Compare(pwd, pwd1) != 0 {
		this.JsonErr("两次密码不同", 1100, "/reg")
	}
	//判断是否有同一个用户
	if models.FindUserByEmail(email) == false || models.FindUserByUname(username) == false {
		this.JsonErr("已经有该用户了",1101, "/reg")
	}


	//保存用户信息
	uid, err := models.SaveUser(username, nick, email, pwd, school, ip)
	if err != nil {
		this.JsonErr("注册失败", 112, "/reg")
	}
	//
	//user, _ := models.QueryUserById(id)
	//
	this.SetSession(SESSION_USER_KEY,uid)

	this.JsonOK("注册成功", "/index")
}