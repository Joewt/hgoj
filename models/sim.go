package models

type Sim struct {
	SId    int32 `orm:"pk"`
	SimSId int32 `orm:"null"`
	Sim    int32 `orm:"null"`
}
