package models

import (
	_ "github.com/astaxie/beego/orm"
	"time"
)



type Problem struct {
	ProblemId		int32		`orm:"auto"`
	Title			string		`orm:"size(200);"`
	Description 	string		`orm:"type(text);null"`
	Input			string		`orm:"type(text);null"`
	Output			string		`orm:"type(text);null"`
	SampleInput		string		`orm:"type(text);null"`
	SampleOutput	string		`orm:"type(text);null"`
	Spj				string		`orm:"type(char);size(1);default(0)"`
	Hint			string		`orm:"type(text);null"`
	Source			string		`orm:"size(100);null"`
	InDate			time.Time	`orm:"null"`
	TimeLimit		int32		`orm:"default(0)"`
	MemoryLimit		int32		`orm:"default(0)"`
	Defunct			string		`orm:"type(char);size(1);default(N)"`
	Accepted		int32		`orm:"null;default(0)"`
	Submit			int32		`orm:"null;default(0)"`
	Solved			int32		`orm:"null;default(0)"`
}


func QueryAllProblem() ( []*Problem, int64, error){
	var pro []*Problem
	problem := new(Problem)
	qs := DB.QueryTable(problem)
	num, err := qs.All(&pro)
	if err != nil {
		return nil,num,err
	}
	return pro,num, nil
}

func QueryProblemById(id int32) (Problem, error) {
	pro := Problem{ProblemId:id}
	err := DB.Read(&pro,"ProblemId")
	if err != nil {
		return Problem{}, err
	}
	return pro, nil
}