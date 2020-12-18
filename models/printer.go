package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
	"time"
)



type Printer struct {
	PrinterId			int32			`orm:"auto"`
	UserId				int32			`orm:"null"`
	InDate				time.Time		`orm:"auto_now_add;type(date);default(2019-05-01 19:00:00)"`
	Status				int8			`orm:"default(0)"`
	Worktime			time.Time		`orm:"auto_now;type(datetime)"`
	Printer				string			`orm:"type(char);size(16);default(LOCAL)"`
	Content				string			`orm:"type(text)"`
}
