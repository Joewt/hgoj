package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Online struct {
	Id					int32
	Hash				string			`orm:"size(32)"`
	Ip					string  		`orm:"size(46);"`
	Ua					string			`orm:""`
	Refer				string			`orm:"null"`
	Lastmove			int32
	Firsttime			int32			`orm:"null"`
	Uri					string			`orm:"null"`
}
