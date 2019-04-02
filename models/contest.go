package models

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)



type Contest struct {
	ContestId	int32		`orm:"auto"`
	Title		string		`orm:"null"`
	StartTime	time.Time	`orm:"default(null);auto_now_add;type(datetime);null"`
	EndTime		time.Time	`orm:"default(null);auto_now_add;type(datetime);null"`
	Defunct		string		`orm:"type(char);size(1);default(N)"`
	Description string		`orm:"type(text);null"`
	Private		uint8		`orm:"type(4);default(0)"`
	Langmask	int			`orm:"default(0);description:(bits for LANG to mask)"`
	Password	string		`orm:"type(char);size(16);"`
	UserId		int32		`orm:"default(1)"`
}