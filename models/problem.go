package models

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/orm"
	"github.com/yinrenxin/hgoj/syserror"
	"strconv"
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

func stringToint32(str string) int32 {
	d,_ := strconv.Atoi(str)
	return int32(d)
}


func QueryAllProblem() ( []*Problem, int64, error){
	var pro []*Problem
	problem := new(Problem)
	qs := DB.QueryTable(problem)
	num, err := qs.OrderBy("-problem_id").All(&pro)
	if err != nil {
		return nil,num,err
	}
	return pro,num, nil
}


func QueryUserProblem(uid int32) ([]*Problem, int64, error) {
	var data []*Solution
	Solutions := new(Solution)
	qss := DB.QueryTable(Solutions)
	_, err := qss.Filter("user_id",uid).Filter("result",4).Limit(10).OrderBy("-solution_id").All(&data,"problem_id","time","in_date","memory","language")
	if err != nil {
		return nil,0,err
	}

	var proIds []int32
	for _, v := range data {
		proIds = append(proIds, v.ProblemId)
	}

	var pro []*Problem
	if proIds == nil{
		return pro,0,nil
	}
	problem := new(Problem)
	qs := DB.QueryTable(problem)
	num, err := qs.Filter("problem_id__in", proIds).OrderBy("-problem_id").All(&pro)
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

func AddProblem(data ...string) (int64,error) {
	var pro Problem
	pro.Title = data[0]
	pro.TimeLimit = stringToint32(data[1])
	pro.MemoryLimit = stringToint32(data[2])
	pro.Description = data[3]
	pro.Input = data[4]
	pro.Output = data[5]
	pro.SampleInput = data[6]
	pro.SampleOutput = data[7]
	pro.InDate = time.Now()
	pro.Defunct = "Y"

	pid, err := DB.Insert(&pro)
	if err != nil {
		return pid, err
	}
	return pid, nil
}

func UpdateProblemById(id int32, data []string) (bool,error) {
	pro := Problem{ProblemId:id}
	if DB.Read(&pro) == nil {
		pro.Title = data[0]
		pro.TimeLimit = stringToint32(data[1])
		pro.MemoryLimit = stringToint32(data[2])
		pro.Description = data[3]
		pro.Input = data[4]
		pro.Output = data[5]
		pro.SampleInput = data[6]
		pro.SampleOutput = data[7]
		pro.InDate = time.Now()
		if num, err := DB.Update(&pro); err == nil {
			logs.Info(num)
			return true, err
		}
	}
	return false, syserror.UpdateProErr()
}

func DelProblemById(id int32) (bool) {
	if num, err := DB.Delete(&Problem{ProblemId: id}); err == nil {
		logs.Info(num)
		return true
	}
	return false
}