package models

import (
	_ "github.com/astaxie/beego/orm"
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
}