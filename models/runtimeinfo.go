package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
)



type Runtimeinfo struct {
	SolutionId			int32			`orm:"pk"`
	Error				string			`orm:"type(text);null"`
}
