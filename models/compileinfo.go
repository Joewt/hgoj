package models

import (
	_ "github.com/astaxie/beego/orm"
)



type CompileInfo struct {
	SolutionId 	int32	`orm:"pk"`
	Error  		string 	`orm:"type(text);null"`
}