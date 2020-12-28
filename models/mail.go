package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
	"time"
)




type Mail struct {
	MailId			int32		`orm:"auto"`
	ToUser			string		`orm:"size(48);default()"`
	FromUser		string		`orm:"size(48);default()"`
	Title			string		`orm:"size(200);default()"`
	Content			string		`orm:"type(text);null"`
	NewMail			int8		`orm:"default(1)"`
	Reply			int8		`orm:"default(0);null"`
	InDate			time.Time	`orm:"null"`
	Defunct			string		`orm:"type(char);size(1);default(N)"`
}
