package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	_ "github.com/beego/beego/v2/adapter/orm"
	client "github.com/beego/beego/v2/client/orm"
)

var JUDGERES = map[int]string{
	0:  "判题中",
	1:  "等待重判",
	2:  "编译中",
	3:  "运行并评判",
	4:  "正确",
	5:  "格式错误",
	6:  "答案错误",
	7:  "时间超限",
	8:  "内存超限",
	9:  "输出超限",
	10: "运行错误",
	11: "编译错误",
	12: "编译成功",
	13: "运行完成",
}

var JUDGERESCLSAA = map[int]string{
	0:  "warning",
	1:  "danger",
	2:  "warning",
	3:  "info",
	4:  "success",
	5:  "warning",
	6:  "danger",
	7:  "warning",
	8:  "warning",
	9:  "warning",
	10: "danger",
	11: "danger",
	12: "info",
	13: "success",
}

type Solution struct {
	SolutionId int32     `orm:"auto"`
	ProblemId  int32     `orm:"default(0)"`
	UserId     int32     `orm:"null"`
	Time       int32     `orm:"auto_now_add;type(date);default(0)"`
	Memory     int32     `orm:"default(0)"`
	InDate     time.Time `orm:"auto_now_add;type(date);default('2019-05-01 19:00:00')"`
	Result     int16     `orm:"default(0)"`
	Language   uint      `orm:"default(0)"`
	Ip         string    `orm:"size(46)"`
	ContestId  int32     `orm:"null;default(0)"`
	Valid      int8      `orm:"default(1)"`
	Num        int8      `orm:"default(-1)"`
	CodeLength int32     `orm:"default(0)"`
	Judgetime  time.Time `orm:"null;type(datetime);default(null)"`
	PassRate   float64   `orm:"digits(3);decimals(2);default(0)"`
	LintError  uint      `orm:"default(0)"`
	Judger     string    `orm:"type(char);size(16);default(LOCAL)"`
}

func UpdateSolutionResultBySid(sid int32) bool {
	Solu := Solution{SolutionId: sid}
	Solu.Result = 1
	_, err := DB.Update(&Solu, "result")
	if err != nil {
		logs.Error(err)
		return false
	}
	return true
}

func UpdateSolutionResultByCid(cid int32) bool {
	_, err := DB.QueryTable("solution").Filter("contest_id", cid).Update(client.Params{
		"result": 1,
	})
	if err != nil {
		logs.Error(err)
		return false
	}
	return true
}

func UpdateSolutionResultByPid(pid int32) bool {
	_, err := DB.QueryTable("solution").Filter("problem_id", pid).Update(client.Params{
		"result": 1,
	})
	if err != nil {
		logs.Error(err)
		return false
	}
	return true
}

func QueryTimeAndMemoryByuidpid(uid, pid int32) (int32, int32) {
	var data Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.Filter("user_id", uid).Filter("problem_id", pid).Filter("result", 4).All(&data)
	if err != nil {
		return 0, 0
	}
	return data.Time, data.Memory
}

func QuerySolutionBySid(sid int32) (Solution, error) {
	Solu := Solution{SolutionId: sid}
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
		return nil, JUDGERES, err
	}
	return data, JUDGERES, nil
}

func QueryTotalNumAcNumSolution(t string) (int64, int64) {
	//var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	totalNum, err := qs.Filter("in_date", t).Count()
	acNum, err := qs.Filter("in_date", t).Filter("result", 4).Count()
	if err != nil {
		return 0, 0
	}
	return totalNum, acNum
}

func QueryPageSolution(start, pageSize int) ([]*Solution, map[int]string, int64, int64, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	totalNum, _ := qs.Count()
	num, err := qs.OrderBy("-solution_id").Limit(pageSize, start).All(&data)
	if err != nil {
		return nil, JUDGERES, num, totalNum, err
	}
	return data, JUDGERES, num, totalNum, nil
}

func QueryUserSolution(uid int32) ([]*Solution, map[int]string, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.Filter("user_id", uid).Limit(10).OrderBy("-solution_id").All(&data)
	if err != nil {
		return nil, JUDGERES, err
	}
	return data, JUDGERES, nil
}

func QueryACSubSolution(uid int32) (int64, int64, error) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	sub, err := qs.Filter("user_id", uid).Count()
	if err != nil {
		return 0, 0, err
	}

	ac, err := qs.Filter("user_id", uid).Filter("result", 4).Count()

	if err != nil {
		return 0, 0, err
	}
	return ac, sub, nil
}

func QueryACSUBFromSolutionBYPidCi(pid, cid int32) (int32, int32) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	sub, err := qs.Filter("problem_id", pid).Filter("contest_id", cid).Count()
	if err != nil {
		return 0, 0
	}
	ac, err := qs.Filter("problem_id", pid).Filter("contest_id", cid).Filter("result", 4).Count()
	if err != nil {
		return 0, 0
	}

	return int32(ac), int32(sub)
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

func QueryAllSolutionByCid(cid int32) ([]*Solution, map[int]string, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	_, err := qs.Filter("contest_id", cid).OrderBy("-solution_id").All(&data)
	if err != nil {
		return nil, JUDGERES, err
	}
	return data, JUDGERES, nil
}

func QueryPageSolutionByCid(cid int32, start, pageSize int) ([]*Solution, map[int]string, int64, int64, error) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	totalNum, _ := qs.Filter("contest_id", cid).Count()
	num, err := qs.Filter("contest_id", cid).OrderBy("-solution_id").Limit(pageSize, start).All(&data)
	if err != nil {
		return nil, JUDGERES, num, totalNum, err
	}
	return data, JUDGERES, num, totalNum, nil
}

func QueryAllUserIdByCid(cid int32) ([]int32, int64) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	num, _ := qs.Distinct().Filter("contest_id", cid).All(&data, "user_id")

	var uid []int32
	for _, v := range data {
		uid = append(uid, v.UserId)
	}

	return uid, num
}

func QueryAllUserByCid(cid int32) ([]*UserInfos, int64) {
	var UserData []TempUser

	num1, _ := DB.Raw("SELECT users.user_id,users.nick from solution,users WHERE `contest_id` = ? and solution.user_id = users.user_id group by users.user_id", cid).QueryRows(&UserData)

	var UserAc []*TempUserAc
	DB.Raw("SELECT `user_id`,count(`solution_id`) as ac from solution WHERE `contest_id` = ? and `result` = 4 group by `user_id`", cid).QueryRows(&UserAc)

	var UserSubmitNum []TempUserSubmit
	DB.Raw("SELECT `user_id`,count(`solution_id`) as total from solution WHERE `contest_id` = ? group by `user_id`", cid).QueryRows(&UserSubmitNum)

	var UserInfo []*UserInfos
	for _, v := range UserData {
		var ac int32
		var total int32
		for _, v2 := range UserAc {
			if v.UserId == v2.UserId {
				ac = v2.Ac
			}
		}
		for _, v3 := range UserSubmitNum {
			if v.UserId == v3.UserId {
				total = v3.Total
			}
		}
		UserInfo = append(UserInfo, &UserInfos{v, ac, total})
	}
	return UserInfo, num1
}

type TempUser struct {
	UserId int32
	Nick   string
}

type TempUserAc struct {
	UserId int32
	Ac     int32
}

type TempUserSubmit struct {
	UserId int32
	Total  int32
}

type UserInfos struct {
	TempUser
	Ac    int32
	Total int32
}

func QueryAllAcNumByCid(cid int32) ([]*TempUserAc, int64) {
	var UserAc []*TempUserAc
	num, _ := DB.Raw("SELECT `user_id`,count(solution_id) from solution WHERE `contest_id` = ? group by user_id", cid).QueryRows(&UserAc)
	return UserAc, num
}

func QueryJudgeTimeFromSolutionByUidCidPid(uid, pid, cid int32, startTime time.Time) (int32, bool, float64, int64) {
	var data []*Solution
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	num, _ := qs.Filter("problem_id", pid).Filter("user_id", uid).Filter("contest_id", cid).All(&data)
	var t float64
	var i int64
	var flag bool
	var ErrNum int64
	flag = false
	for _, v := range data {
		if v.Result == 4 {
			i++
			t = v.Judgetime.Sub(startTime).Seconds()
			flag = true
		}
	}
	var total float64
	if flag {
		total = t + float64(num-i)*20*60
	}
	ErrNum = num - i
	return pid, flag, total, ErrNum

}

func QueryACNickTotalByUid(uid int32, cid int32) (string, int64, int64) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	ac, _ := qs.Filter("user_id", uid).Filter("contest_id", cid).Filter("result", 4).Count()
	user, _ := QueryUserById(uid)
	nick := user.Nick
	total, _ := qs.Filter("user_id", uid).Filter("contest_id", cid).Count()
	return nick, ac, total
}

func QueryACNickTotalByUidPid(uid int32, cid int32, pid int32) (string, int64, int64) {
	Solutions := new(Solution)
	qs := DB.QueryTable(Solutions)
	ac, _ := qs.Filter("user_id", uid).Filter("contest_id", cid).Filter("result", 4).Filter("problem_id", pid).Count()
	user, _ := QueryUserById(uid)
	nick := user.Nick
	total, _ := qs.Filter("user_id", uid).Filter("contest_id", cid).Filter("problem_id", pid).Count()
	return nick, ac, total
}

func AddSolution(pid string, source string, uid int32, codeLen int, lang string, conid int32, ip string) (int64, error) {
	var Solu Solution
	var SoluCode SourceCode
	err := DB.Begin()
	Solu.ProblemId = stringToint32(pid)
	Solu.UserId = uid
	Solu.InDate = time.Now()
	Solu.Language = uint(stringToint32(lang))
	Solu.CodeLength = int32(codeLen)
	Solu.Result = 0
	logs.Info("conid", conid)
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
