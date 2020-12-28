package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
	"time"
)


type ShareCode struct {
	ShareId			int32		`orm:"auto"`
	UserId			int32		`orm:"null"`
	Share_code		string		`orm:"type(text);null"`
	Language		string		`orm:"type(32);null"`
	ShareTime		time.Time	`orm:"auto_now;type(datetime);null"`
}
