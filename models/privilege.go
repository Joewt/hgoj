package models

import (
	_ "github.com/astaxie/beego/orm"
)



type Privilege struct {
	UserId			int32		`orm:"pk"`
	RightStr		string		`orm:"type(char);size(30);null"`
	Defunct			string		`orm:"type(char);size(1);default(N)"`
}

