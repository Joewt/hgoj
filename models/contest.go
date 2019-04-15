package models

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/orm"
	"github.com/yinrenxin/hgoj/syserror"
	"strconv"
	"strings"
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
}


func ContestAdd(title, desc,proIds,role,limituser string,startTime, endTime time.Time) (int32, error) {
	err := DB.Begin()
	var con Contest
	con.Title = title
	con.StartTime = startTime
	con.EndTime = endTime
	con.Defunct = "Y"
	con.Description = desc
	con.Password = role
	con.Password = ""
	logs.Info("问题为：", proIds)

	cid, err2 := DB.Insert(&con)

	temp := strings.Split(proIds,",")

	conPro := []ContestProblem{}
	for _,v := range temp {
		proId,err3 := strconv.Atoi(v)
		if err3 != nil {
			_ = DB.Rollback()
			return 0, err3
		}
		pro := Problem{ProblemId:int32(proId)}
		if DB.Read(&pro) != nil {
			_ = DB.Rollback()
			return 0, syserror.NoProError{}
		}
		conPro = append(conPro, ContestProblem{ProblemId:int32(proId),ContestId:int32(cid)})
	}

	_, err1 := DB.InsertMulti(4, conPro)

	if err2 != nil ||  err1 != nil {
		err = DB.Rollback()
		return 0, err
	}

	err = DB.Commit()
	return int32(cid), nil
}