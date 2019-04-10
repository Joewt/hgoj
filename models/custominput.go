package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Custominput struct {
	SolutionId			int32			`orm:"pk"`
	InputText			string			`orm:"type(text);null"`
}
