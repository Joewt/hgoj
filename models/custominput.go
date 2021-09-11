package models

type Custominput struct {
	SolutionId int32  `orm:"pk"`
	InputText  string `orm:"type(text);null"`
}
