package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Sim struct {
	Id			int32
	SId			int32		`orm:"null"`
	SimSId		int32		`orm:"null"`
	Sim			int32		`orm:"null"`
}
