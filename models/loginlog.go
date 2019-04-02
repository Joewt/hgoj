package models

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)



type LoginLog struct {
	Id			int32
	UserId		int32		`orm:"null"`
	Password 	string 		`orm:"size(32);null"`
	IP			string  	`orm:"size(46);null"`
	Time		time.Time	`orm:"auto_now_add;type(date);null"`
}
