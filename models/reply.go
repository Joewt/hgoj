package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
	"time"
)



type Reply struct {
	Rid				int32			`orm:"auto"`
	AuthorId		int32			`orm:"null"`
	Time			time.Time		`orm:"auto_now_add;type(datetime);default(2019-05-01 19:00:00)"`
	Content			string			`orm:"type(text);null"`
	TopicId			int32			`orm:"null"`
	Status			int8			`orm:"default(0)"`
	Ip				string  		`orm:"size(46);null"`
}
