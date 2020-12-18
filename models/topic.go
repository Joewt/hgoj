package models

import (
	_ "github.com/beego/beego/v2/adapter/orm"
)



type Topic struct {
	Tid			int32			`orm:"auto"`
	Title		string			`orm:"size(60);default()"`
	Status		int32			`orm:"size(2);default(0)"`
	TopLevel	int32			`orm:"size(2);default(0)"`
	Cid			int32			`orm:"null"`
	Pid			int32			`orm:""`
	AuthorId	int32			`orm:""`
}
