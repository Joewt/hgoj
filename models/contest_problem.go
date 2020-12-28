package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
)



type ContestProblem struct {
	Id				int32
	ProblemId		int32		`orm:"default(0)"`
	ContestId		int32		`orm:"null"`
	Title			string		`orm:"size(200)"`
	Num				int32		`orm:"default(0)"`
	CAccepted		int32		`orm:"default(0)"`
	CSubmit			int32		`orm:"default(0)"`
}



func QueryProblemByCid(cid int32)([]Problem, error) {
	qs := DB.QueryTable("contest_problem")
	var c []*ContestProblem
	qs.Filter("contest_id", cid).All(&c)
	//var proIds  []int32
	var p []Problem
	for _, v := range c {
		//proIds = append(proIds,v.ProblemId)
		pro := Problem{ProblemId:v.ProblemId}
		err := DB.Read(&pro,"ProblemId")
		if err != nil {
			return []Problem{}, err
		}
		p = append(p, pro)
	}
	return p, nil
}
