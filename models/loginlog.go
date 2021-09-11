package models

import "time"

type LoginLog struct {
	UserId   int32     `orm:"pk"`
	Password string    `orm:"size(32);null"`
	Ip       string    `orm:"size(46);null"`
	Time     time.Time `orm:"auto_now_add;type(date);null"`
}
