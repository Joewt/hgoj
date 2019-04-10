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


func QueryAllSolution() ([]*Solution, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.OrderBy("-solution_id").All(&data)
	if err != nil {
		return nil,err
	}
	return data, nil
}



func AddSolution(pid string, source string, uid string, codeLen int)(int64, error){
	var Solu Solution
	var SoluCode SourceCode
	err := DB.Begin()
	Solu.ProblemId = stringToint32(pid)
	Solu.UserId = stringToint32(uid)
	Solu.InDate = time.Now()
	Solu.Language = 1
	Solu.Ip = "127.0.0.1"
	Solu.CodeLength = int32(codeLen)
	Solu.Result = 1

	sid, err := DB.Insert(&Solu)

	SoluCode.SolutionId = int32(sid)
	SoluCode.Source = source

	scid, err := DB.Insert(&SoluCode)
	if err != nil {
		return scid, err
	}

	if sid == 0 || scid == 0 {
		err = DB.Rollback()
		return sid, err
	} else {
		err = DB.Commit()
		return sid, err
	}
}