package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
)



type Custominput struct {
	SolutionId			int32			`orm:"pk"`
	InputText			string			`orm:"type(text);null"`
}
