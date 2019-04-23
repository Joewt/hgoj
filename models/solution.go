package models

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/orm"
	"time"
)


var JUDGERES = map[int]string{
	0: "判题中",
	1: "等待重判",
	2: "编译中",
	3: "运行并评判",
	4: "正确",
	5: "格式错误",
	6: "答案错误",
	7: "时间超限",
	8: "内存超限",
	9: "输出超限",
	10: "运行错误",
	11: "编译错误",
	12: "编译成功",
	13: "运行完成",
}


var JUDGERESCLSAA = map[int]string{
	0: "warning",
	1: "danger",
	2: "warning",
	3: "info",
	4: "success",
	5: "warning",
	6: "danger",
	7: "warning",
	8: "warning",
	9: "warning",
	10: "danger",
	11: "danger",
	12: "info",
	13: "success",
}


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


func QuerySolutionBySid(sid int32) (Solution, error) {
	Solu := Solution{SolutionId:sid}
	err := DB.Read(&Solu)
	if err != nil {
		return Solution{}, err
	}
	return Solu, nil
}


func QueryAllSolution() ([]*Solution, map[int]string, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.OrderBy("-solution_id").All(&data)
	if err != nil {
		return nil,JUDGERES,err
	}
	return data, JUDGERES,nil
}


func QueryPageSolution(start , pageSize int) ([]*Solution, map[int]string,int64,int64, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	totalNum, _ := qs.Count()
	num, err := qs.OrderBy("-solution_id").Limit(pageSize, start).All(&data)
	if err != nil {
		return nil,JUDGERES,num,totalNum,err
	}
	return data, JUDGERES,num,totalNum,nil
}


func QueryUserSolution(uid int32) ([]*Solution, map[int]string, error){
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.Filter("user_id",uid).Limit(10).OrderBy("-solution_id").All(&data)
	if err != nil {
		return nil,JUDGERES,err
	}
	return data, JUDGERES,nil
}


func QueryACSubSolution(uid int32) (int64,int64, error) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	sub, err := qs.Filter("user_id", uid).Count()
	if err != nil {
		return 0,0,err
	}

	ac, err := qs.Filter("user_id", uid).Filter("result",4).Count()

	if err != nil {
		return 0,0,err
	}
	return ac, sub, nil
}


func QueryACSUBFromSolutionBYPidCi(pid,cid int32) (int32,int32) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	sub, err := qs.Filter("problem_id", pid).Filter("contest_id",cid).Count()
	if err != nil {
		return 0,0
	}
	ac, err := qs.Filter("problem_id", pid).Filter("contest_id",cid).Filter("result", 4).Count()
	if err != nil {
		return 0,0
	}

	return int32(ac),int32(sub)
}


func QueryResultUserSolution(uid int32) (map[int]int64, error) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)

	data := make(map[int]int64)
	for k, _ := range JUDGERES {
		num, _ := qs.Filter("user_id", uid).Filter("result", k).Count()
		data[k] = num

	}
	return data, nil
}


func QueryAllSolutionByCid(cid int32)  ([]*Solution, map[int]string, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.Filter("contest_id",cid).OrderBy("-solution_id").All(&data)
	if err != nil {
		return nil,JUDGERES,err
	}
	return data, JUDGERES,nil
}


func QueryPageSolutionByCid(cid int32,start , pageSize int)  ([]*Solution, map[int]string, int64,int64,error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	totalNum, _ := qs.Filter("contest_id",cid).Count()
	num, err := qs.Filter("contest_id",cid).OrderBy("-solution_id").Limit(pageSize, start).All(&data)
	if err != nil {
		return nil,JUDGERES,num,totalNum,err
	}
	return data, JUDGERES,num,totalNum,nil
}


func QueryAllUserIdByCid(cid int32) ([]int32, int64) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	num, _ := qs.Distinct().Filter("contest_id",cid).All(&data,"user_id")

	var uid []int32
	for _, v := range data {
		uid = append(uid, v.UserId)
	}

	return uid, num
}


func QueryJudgeTimeFromSolutionByUidCidPid(uid,pid,cid int32, startTime time.Time)(int32,bool,float64,int64) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	num, _ := qs.Filter("problem_id", pid).Filter("user_id", uid).Filter("contest_id", cid).All(&data)
	var t  float64
	var i int64
	var flag bool
	var ErrNum int64
	flag = false
	for _, v := range data {
		 i = 0
		if v.Result == 4 {
			i++
			t = v.Judgetime.Sub(startTime).Seconds()
			flag = true
		}
	}

	total := t + float64(num-i)*20
	ErrNum = num - i
	return pid,flag,total, ErrNum

}


func QueryACNickTotalByUid(uid int32, cid int32) (string, int64,int64) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	ac, _ := qs.Filter("user_id", uid).Filter("contest_id", cid).Filter("result", 4).Count()
	user, _ := QueryUserById(uid)
	nick := user.Nick
	total, _ := qs.Filter("user_id", uid).Filter("contest_id", cid).Count()
	return nick, ac, total
}


func AddSolution(pid string, source string, uid int32, codeLen int, lang string, conid int32, ip string)(int64, error){
	var Solu Solution
	var SoluCode SourceCode
	err := DB.Begin()
	Solu.ProblemId = stringToint32(pid)
	Solu.UserId = uid
	Solu.InDate = time.Now()
	Solu.Language = uint(stringToint32(lang))
	Solu.CodeLength = int32(codeLen)
	Solu.Result = 0
	logs.Info("conid",conid)
	if conid != -1 {
		Solu.ContestId = conid
	}
	Solu.Ip = ip

	sid, err := DB.Insert(&Solu)

	SoluCode.SolutionId = int32(sid)
	SoluCode.Source = source

	scid, err := DB.Insert(&SoluCode)

	if sid == 0 || scid != 0 {
		err = DB.Rollback()
		return sid, err
	} else {
		err = DB.Commit()
		return sid, err
	}
}