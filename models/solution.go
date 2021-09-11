package models

import (
	"time"
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
