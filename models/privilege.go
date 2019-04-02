package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Privilege struct {
	Id				int32
	UserId			int32		`orm:"null"`
	RightStr		string		`orm:"type(char);size(30);null"`
	Defunct			string		`orm:"type(char);size(1);default(N)"`
}

