package models

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/orm"
)



type Compileinfo struct {
	SolutionId 	int32	`orm:"pk"`
	Error  		string 	`orm:"type(text);null"`
}



func QueryCompileInfoBySid(sid int32) (string,error) {
	compile := Compileinfo{SolutionId: sid}
	err := DB.Read(&compile, "SolutionId")
	if err != nil {
		logs.Warn(err)
		return Compileinfo{}.Error,err
	}
	return compile.Error, nil
}