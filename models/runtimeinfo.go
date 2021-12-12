package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	_ "github.com/beego/beego/v2/adapter/orm"
)

type Runtimeinfo struct {
	SolutionId int32  `orm:"pk"`
	Error      string `orm:"type(text);null"`
}

func QueryRuntimeInfoBySid(sid int32) (string, error) {
	compile := Runtimeinfo{SolutionId: sid}
	err := DB.Read(&compile, "SolutionId")
	if err != nil {
		logs.Warn(err)
		return Runtimeinfo{}.Error, err
	}
	return compile.Error, nil
}
