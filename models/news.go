package models

import "time"



type News struct {
	NewsId			int32			`orm:"auto"`
	UserId			int32			`orm:"null"`
	Title			string			`orm:"size(200);null"`
	Content			string			`orm:"type(text)"`
	Time			time.Time		`orm:"auto_now_add;type(date);default(2019-05-01 19:00:00)"`
	Importance		int8			`orm:"default(0)"`
	Defunct			string			`orm:"type(char);size(1);default(N)"`
}