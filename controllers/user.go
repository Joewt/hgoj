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

// @router /profile [get]
func (this *UserController) Profile() {
	//id := this.Ctx.Input.Param(":uid")
	//uid, _ := tools.StringToInt32(id)
	//if !this.IsAdmin {
	//	if uid != this.User.UserId {
	//		this.Abort("401")
	//	}
	//}

	if !this.IsLogin {
		this.Abort("401")
	}

	uid := this.User.UserId

	data,RESULT,err := models.QueryUserSolution(uid)
	if err != nil {
		logs.Error(err)
	}

	pros,_,err := models.QueryUserProblem(uid)
	if err != nil {
		logs.Error(err)
	}

	ac,sub,err := models.QueryACSubSolution(uid)
	if err != nil {
		ac = 0
		sub = 0
	}

	ResData, err := models.QueryResultUserSolution(uid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["resdata"] = ResData
	this.Data["problems"] = pros
	this.Data["data"] = data
	this.Data["RES"] = RESULT
	avatarUrl := tools.AvatarLink(this.User.Email,500)
	this.Data["avatar"] = avatarUrl
	this.Data["acnum"] = ac
	this.Data["subnum"] = sub
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


// @router /profile/update [get]
func (this *UserController) UserUpdateGet() {
	if !this.IsLogin {
		this.Abort("401")
	}
	this.TplName = "profileUpdate.html"
}


// @router /profile/update [post]
func (this *UserController) UserUpdatePost() {
	if !this.IsLogin {
		this.Abort("500")
	}
	//username := this.GetMushString("username", "用户名不能为空")
	nick := this.GetMushString("nick", "昵称不能为空")
	//email := this.GetMushString("email","邮箱不能为空")
	oldpwd   := this.GetMushString("oldpwd", "旧密码不能为空")
	newpwd  := this.GetMushString("newpwd", "新密码不能为空")
	//school := this.GetMushString("school", "学校不能为空")


	//if username != this.User.UserName {
	//	this.JsonErr("用户名不同", 9002, "")
	//}

	if oldpwd == newpwd {
		this.JsonErr("两次密码不能一样", 9000, "")
	}

	if tools.MD5(oldpwd) != this.User.Password {
		this.JsonErr("密码错误", 9001, "")
	}

	ok , err := models.UpdateUserInfo(this.User.UserId,nick,newpwd)
	if !ok {
		this.JsonErr(err.Error(), 9004, "")
	}
	this.DelSession(SESSION_USER_KEY)
	this.JsonOK("更新成功", "/login")
}



// @router /forgotpwd [get]
func (this *UserController) UserForgotPwd() {
	this.TplName = "forgotPassword.html"
}


// @router /forgotpwd/sendemail [post]
func (this *UserController) SendEmailForgot() {
	this.JsonErr("未知错误", 9010, "")
}