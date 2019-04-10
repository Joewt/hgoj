package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Runtimeinfo struct {
	SolutionId			int32			`orm:"pk"`
	Error				string			`orm:"type(text);null"`
}
