package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/syserror"

	"github.com/yinrenxin/hgoj/models"
	//"github.com/yinrenxin/hgoj/syserror"
	"github.com/yinrenxin/hgoj/tools"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /profile/:uid [get]
func (this *UserController) Profile() {
	id := this.Ctx.Input.Param(":uid")
	uid, _ := tools.StringToInt32(id)
	if !this.IsAdmin {
		if uid != this.User.UserId {
			this.Abort("401")
		}
	}
	avatarUrl := tools.AvatarLink(this.User.Email)
	this.Data["avatar"] = avatarUrl
	this.TplName = "profile.html"
}


// @router /user/reg [post]
func (this *UserController) UserReg() {
	logs.Info("是否登录:",this.IsLogin)
	if this.IsLogin {
		this.Abort("500")
	}
	username := this.GetMushString("username", "用户名不能为空")
	nick := this.GetMushString("nick", "昵称不能为空")
	email := this.GetMushString("email","邮箱不能为空")
	pwd   := this.GetMushString("pwd", "密码不能为空")
	pwd1  := this.GetMushString("pwd2", "确认密码不能为空")
	school := this.GetMushString("school", "学校不能为空")
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
	//panic("xx")
	logs.Info("到这里不会执行了")
	//保存用户信息
	uid, err := models.SaveUser(username, nick, email, pwd, school, ip)
	if err != nil {
		this.JsonErr("注册失败", 112, "/reg")
	}

	//
	user, _ := models.QueryUserById(uid)
	//
	this.SetSession(SESSION_USER_KEY,user)

	this.JsonOK("注册成功", "/index")
}


// @router /login [post]
func (this *UserController) Login() {
	ue := this.GetMushString("ue", "用户名或邮箱不能为空")
	pwd   := this.GetMushString("pwd", "密码不能为空")
	id, err := models.QueryUserByUEAndPwd(ue, pwd)

	if err != nil {
		//this.Abort500(syserror.New("登录失败",err))
		this.JsonErr(err.Error(), 2000, "/login")
	}

	user, _ := models.QueryUserById(id)
	if err != nil {
		logs.Warn(err)
	}

	this.SetSession(SESSION_USER_KEY,user)

	this.JsonOK("登录成功","/")
}


// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/index", 302)
}


// @router /user/list [get]
func (this *UserController) UserList() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	user,_, err := models.QueryAllUser()
	if err != nil {
		this.JsonErr("未知错误", 4000, "/index")
	}
	this.Data["user"] = user
	this.TplName = "admin/userList.html"
}