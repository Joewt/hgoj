package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Runtimeinfo struct {
	Id					int32
	SolutionId			int32			`orm:"default(0)"`
	Error				string			`orm:"type(text);null"`
}
