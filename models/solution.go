package models

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)



type Solution struct {
	SolutionId			int32			`orm:"auto"`
	ProblemId			int32			`orm:"default(0)"`
	UserId				int32			`orm:"null"`
	Time				int32			`orm:"auto_now_add;type(date);default(0)"`
	Memory				int32			`orm:"default(0)"`
	InDate				time.Time		`orm:"auto_now_add;type(date);default('2019-05-01 19:00:00')"`
	Result				int16			`orm:"default(0)"`
	Language			uint			`orm:"default(0)"`
	Ip					string  		`orm:"size(46)"`
	ContestId			int32			`orm:"null;default(0)"`
	Valid				int8			`orm:"default(1)"`
	Num					int8			`orm:"default(-1)"`
	CodeLength			int32			`orm:"default(0)"`
	Judgetime			time.Time		`orm:"auto_now;type(datetime)"`
	PassRate			float64 		`orm:"digits(3);decimals(2);default(0)"`
	LintError			uint			`orm:"default(0)"`
	Judger				string			`orm:"type(char);size(16);default(LOCAL)"`
}
