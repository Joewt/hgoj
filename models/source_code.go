package models

type SourceCode struct {
	SolutionId int32  `orm:"pk"`
	Source     string `orm:"type(text)"`
}
