package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
)



type SourceCode struct {
	SolutionId			int32			`orm:"pk"`
	Source				string			`orm:"type(text)"`
}



func QuerySourceBySolutionId(sid int32) (string) {
	SC := SourceCode{SolutionId: sid}
	err := DB.Read(&SC)
	if err != nil {
		return ""
	}
	return SC.Source
}
