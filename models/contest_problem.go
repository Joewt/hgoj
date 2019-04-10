package models

import (
	_ "github.com/astaxie/beego/orm"
)



type ContestProblem struct {
	ProblemId		int32		`orm:"pk"`
	ContestId		int32		`orm:"null"`
	Title			string		`orm:"size(200)"`
	Num				int32		`orm:"default(0)"`
	CAccepted		int32		`orm:"default(0)"`
	CSubmit			int32		`orm:"default(0)"`
}
