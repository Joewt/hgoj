package models

type ContestProblem struct {
	Id        int32
	ProblemId int32  `orm:"default(0)"`
	ContestId int32  `orm:"null"`
	Title     string `orm:"size(200)"`
	Num       int32  `orm:"default(0)"`
	CAccepted int32  `orm:"default(0)"`
	CSubmit   int32  `orm:"default(0)"`
}
