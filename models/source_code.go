package models

import (
	_ "github.com/astaxie/beego/orm"
)



type SourceCode struct {
	Id					int32
	SolutionId			int32			`orm:"null"`
	Source				string			`orm:"type(text)"`
}
