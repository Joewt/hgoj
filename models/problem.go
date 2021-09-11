package models

import "time"

type Problem struct {
	ProblemId    int32     `orm:"auto"`
	Title        string    `orm:"size(200);"`
	Description  string    `orm:"type(text);null"`
	Input        string    `orm:"type(text);null"`
	Output       string    `orm:"type(text);null"`
	SampleInput  string    `orm:"type(text);null"`
	SampleOutput string    `orm:"type(text);null"`
	Spj          string    `orm:"type(char);size(1);default(0)"`
	Hint         string    `orm:"type(text);null"`
	Source       string    `orm:"size(100);null"`
	InDate       time.Time `orm:"default(null);type(datetime);null"`
	TimeLimit    int32     `orm:"default(0)"`
	MemoryLimit  int32     `orm:"default(0)"`
	Defunct      string    `orm:"type(char);size(1);default(N)"`
	Accepted     int32     `orm:"null;default(0)"`
	Submit       int32     `orm:"null;default(0)"`
	Solved       int32     `orm:"null;default(0)"`
}
