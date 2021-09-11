package models

type Runtimeinfo struct {
	SolutionId int32  `orm:"pk"`
	Error      string `orm:"type(text);null"`
}
