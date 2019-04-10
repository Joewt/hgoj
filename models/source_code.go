package models

import (
	_ "github.com/astaxie/beego/orm"
)



type SourceCode struct {
	SolutionId			int32			`orm:"pk"`
	Source				string			`orm:"type(text)"`
}
