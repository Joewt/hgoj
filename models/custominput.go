package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Custominput struct {
	Id					int32
	SolutionId			int32			`orm:"null;default(0)"`
	InputText			string			`orm:"type(text);null"`
}
