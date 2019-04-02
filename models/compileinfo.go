package models

import (
	_ "github.com/astaxie/beego/orm"
)



type CompileInfo struct {
	Id			int32
	SolutionId 	int32 	`orm:"null"`
	Error  		string 	`orm:"type(text);null"`
}