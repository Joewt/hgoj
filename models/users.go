package models

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	_ "github.com/beego/beego/v2/adapter/orm"
	"github.com/yinrenxin/hgoj/syserror"
	"github.com/yinrenxin/hgoj/tools"
	"time"
)



type Users struct {
	UserId		int32		`orm:"auto"`
	UserName    string		`orm:"size(46)"`
	Email		string		`orm:"size(100);null"`
	Submit		int32 		`orm:"default(0);null"`
	Solved		int32		`orm:"default(0);null"`
	Defunct		string		`orm:"type(char);size(1);default(N)"`
	Ip			string  	`orm:"size(46)"`
	Accesstime  time.Time 	`orm:"auto_now_add;type(datetime);null"`
	Volume		int32		`orm:"default(1)"`
	Language 	int32  		`orm:"default(1)"`
	Password 	string 		`orm:"size(32);null;"`
	RegTime  	time.Time 	`orm:"auto_now_add;type(datetime);null"`
	Nick    	string 		`orm:"size(20)"`
	School 		string 		`orm:"size(20)"`
	Role        int32		`orm:"default(0);null"`
}



//func QueryUserByEmailAndPwd(email, pwd string) (uint,error) {
//	user := User{Email: email}
//	err := DB.Read(&user,"Email")
//	if err != nil{
//		err = syserror.New("没有该账号",nil)
//		return 247, err
//	}
//	if user.Pwd != common.MD5(pwd) {
//		err = syserror.New("密码错误", nil)
//		return 247, err
//	}
//	return user.Id, nil
//}

//func QueryUserByUname() (user Users, err error){
//	user = Users{UserId: id}
//	err = DB.Read(&user, "Id")
//	if err != nil {
//		logs.Warn(err)
//	}
//	return user, nil
//}


func FindUserByEmail(email string) bool {
	user := Users{Email: email}
	err := DB.Read(&user, "Email")
	if err != nil {
		return true
	}
	return false
}

func FindUserByUname(uname string) bool {
	user := Users{UserName: uname}
	err := DB.Read(&user, "UserName")
	if err != nil {
		return true
	}
	return false
}


func SaveUser(username,nick, email,pwd,school,ip string) (int32, error) {
	cnt, err := DB.QueryTable("users").Count()
	if cnt == 0 {
		user := new(Users)
		user.UserName = username
		user.Password = tools.MD5(pwd)
		user.Email = email
		user.Nick = nick
		user.Role = 1
		user.School = school
		user.Ip = ip
		id, err := DB.Insert(user)
		if err != nil {
			return int32(id), err
		}
		logs.Info("新注册一个管理员账户: ", id)
		return int32(id), nil
	}
	user := new(Users)
	user.UserName = username
	user.Password = tools.MD5(pwd)
	user.Email = email
	user.Nick = nick
	user.Role = 0
	user.School = school
	user.Ip = ip
	id, err := DB.Insert(user)
	if err != nil {
		fmt.Println("加油，你最棒！")
		return int32(id), err
	}
	logs.Info("新注册一个普通用户: ", id)
	return int32(id), nil
}


func QueryUserById(id int32) (Users,error) {
	user := Users{UserId: id}
	err := DB.Read(&user, "UserId")
	if err != nil {
		logs.Warn(err)
		return Users{},err
	}
	return user, nil
}


func QueryUserByUEAndPwd(ue, pwd string) (int32, error) {
	user := Users{Email: ue}
	err := DB.Read(&user,"Email")
	if err != nil{
		user = Users{UserName:ue}
		err = DB.Read(&user, "UserName")
		if err != nil {
			err = syserror.New("没有该账号",nil)
			return 0, err
		}
	}
	if user.Password != tools.MD5(pwd) {
		err = syserror.New("密码错误", nil)
		return 0, err
	}
	return user.UserId, nil
}


func QueryLimitUser()([]*Users,int64, error) {
	var u []*Users
	user := new(Users)
	qs := DB.QueryTable(user)
	num, err := qs.OrderBy("-user_id").Limit(50).All(&u)
	if err != nil {
		return nil,num,err
	}
	return u,num, nil
}

func QueryAllUser() ([]*Users,int64, error) {
	var u []*Users
	user := new(Users)
	qs := DB.QueryTable(user)
	num, err := qs.OrderBy("-user_id").All(&u)
	if err != nil {
		return nil,num,err
	}
	return u,num, nil
}


func QueryUserByRole() ([]*Users,int64, error){
	var u []*Users
	user := new(Users)
	qs := DB.QueryTable(user)
	num, err := qs.OrderBy("-user_id").Filter("role__gt",0).All(&u)
	if err != nil {
		return nil,num,err
	}
	return u,num, nil
}


func QueryPageUser(start , pageSize int) ([]*Users,int64, int64,error) {
	var u []*Users
	user := new(Users)
	qs := DB.QueryTable(user)
	totalNum, _ := qs.Count()
	num, err := qs.OrderBy("-user_id").Limit(pageSize,start).All(&u)
	if err != nil {
		return nil,num,totalNum,err
	}
	return u,num, totalNum, nil
}



func UpdateUserInfo(uid int32,nick, pwd string) (bool, error) {
	user := Users{UserId:uid}
	user.Nick = nick
	user.Password = tools.MD5(pwd)
	_, err := DB.Update(&user, "nick", "password")
	if err != nil {
		return false, err
	}
	return true, nil
}

func QueryUidByUname(uname string) (int32) {
	user := Users{UserName: uname}
	err := DB.Read(&user, "UserName")
	if err != nil {
		return 0
	}
	return user.UserId
}

func UpdateUserRoleByUname(uname string, role int32) (bool) {

	user := Users{UserId:QueryUidByUname(uname)}
	user.Role = role
	_, err := DB.Update(&user,"role")
	if err != nil {
		return false
	}
	return true
}


func UpdateUserPwdByUname(uname string, pwd string) (bool) {
	user := Users{UserId:QueryUidByUname(uname)}
	user.Password = pwd
	_, err := DB.Update(&user,"password")
	if err != nil {
		return false
	}
	return true
}
