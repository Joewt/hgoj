package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	_ "github.com/beego/beego/v2/adapter/orm"
	"github.com/yinrenxin/hgoj/syserror"
	"strconv"
	"strings"
	"time"
)

type Contest struct {
	ContestId   int32     `orm:"auto"`
	Title       string    `orm:"null"`
	StartTime   time.Time `orm:"default(null);type(datetime);null"`
	EndTime     time.Time `orm:"default(null);type(datetime);null"`
	Defunct     string    `orm:"type(char);size(1);default(N)"`
	Description string    `orm:"type(text);null"`
	Private     uint8     `orm:"type(4);default(0)"`
	Langmask    int       `orm:"default(0);description:(bits for LANG to mask)"`
	Password    string    `orm:"type(char);size(16);"`
	UserId      int32     `orm:"null"`
}

func ContestAdd(title, desc, proIds, role, limituser string, startTime time.Time, endTime time.Time, uid int32) (int32, error) {
	err := DB.Begin()
	var con Contest
	con.Title = title
	con.StartTime = startTime
	con.EndTime = endTime
	con.Defunct = "N"
	con.Description = desc
	con.Password = role
	con.Password = ""
	con.UserId = uid
	logs.Info("问题为：", proIds)

	cid, err2 := DB.Insert(&con)

	temp := strings.Split(proIds, ",")

	conPro := []ContestProblem{}
	for _, v := range temp {
		proId, err3 := strconv.Atoi(v)
		if err3 != nil {
			_ = DB.Rollback()
			return 0, err3
		}
		pro := Problem{ProblemId: int32(proId)}
		if DB.Read(&pro) != nil {
			_ = DB.Rollback()
			return 0, syserror.NoProError{}
		}
		conPro = append(conPro, ContestProblem{ProblemId: int32(proId), ContestId: int32(cid)})
	}

	_, err1 := DB.InsertMulti(4, conPro)

	if err2 != nil || err1 != nil {
		err = DB.Rollback()
		return 0, err
	}

	err = DB.Commit()
	return int32(cid), nil
}

func ContestUpdate(cid int32, title, desc, proIds, role, limituser string, startTime time.Time, endTime time.Time) (int32, error) {
	err := DB.Begin()
	con := Contest{ContestId: cid}
	con.Title = title
	con.StartTime = startTime
	con.EndTime = endTime
	con.Defunct = "N"
	con.Description = desc
	con.Password = role
	con.Password = ""

	_, err2 := DB.Update(&con, "title", "start_time", "defunct", "description", "password", "end_time")

	temp := strings.Split(proIds, ",")

	conPro := []ContestProblem{}
	for _, v := range temp {
		proId, err3 := strconv.Atoi(v)
		if err3 != nil {
			_ = DB.Rollback()
			return 0, err3
		}
		pro := Problem{ProblemId: int32(proId)}
		if DB.Read(&pro) != nil {
			_ = DB.Rollback()
			return 0, syserror.NoProError{}
		}
		if ok := QueryConProByPidCid(int32(proId), cid); ok {
			conPro = append(conPro, ContestProblem{ProblemId: int32(proId), ContestId: cid})
		}
	}

	_, err1 := DB.InsertMulti(4, conPro)

	if (len(conPro) > 0 && err1 != nil) || err2 != nil {
		err = DB.Rollback()
		return 0, err
	}

	err = DB.Commit()
	return cid, nil
}

func UpdateContestStatus(cid int32) bool {
	con := Contest{ContestId: cid}
	if DB.Read(&con) == nil {
		if con.Defunct == "Y" {
			con.Defunct = "N"
		} else {
			con.Defunct = "Y"
		}
		if num, err := DB.Update(&con); err == nil {
			logs.Info(num)
			return true
		}
	}
	return false
}

func QueryConProByPidCid(pid, cid int32) bool {
	pb := ContestProblem{ProblemId: pid, ContestId: cid}
	err := DB.Read(&pb, "ProblemId", "ContestId")
	if err != nil {
		return true
	}
	return false
}

func QueryAllContest() ([]*Contest, int64, error) {
	var con []*Contest
	contest := new(Contest)
	qs := DB.QueryTable(contest)
	num, err := qs.OrderBy("-contest_id").All(&con)
	if err != nil {
		return nil, num, err
	}
	return con, num, nil
}

func QueryPageContest(start, pageSize int) ([]*Contest, int64, int64, error) {
	var con []*Contest
	contest := new(Contest)
	qs := DB.QueryTable(contest)
	totalNum, err := qs.Count()
	num, err := qs.OrderBy("-contest_id").Limit(pageSize, start).All(&con)
	if err != nil {
		return nil, num, totalNum, err
	}
	return con, num, totalNum, nil
}

func QueryContestByConId(cid int32) (Contest, error) {
	con := Contest{ContestId: cid}

	err := DB.Read(&con, "ContestId")

	if err != nil {
		return Contest{}, err
	}

	return con, nil
}

func QueryACNumContestByCid(cid int32) (int32, int32) {
	var acNum []*ContestProblem
	var ac, sub int32
	contestproblem := new(ContestProblem)
	num, err := DB.QueryTable(contestproblem).Filter("contest_id", cid).All(&acNum)
	if err != nil {
		logs.Info(err)
	}

	logs.Info(num)
	for _, v := range acNum {
		ac += v.CAccepted
		sub += v.CSubmit
	}
	return ac, sub
}
