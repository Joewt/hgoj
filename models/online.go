package models

type Online struct {
	Hash      string `orm:"pk"`
	Ip        string `orm:"size(46);"`
	Ua        string `orm:""`
	Refer     string `orm:"null"`
	Lastmove  int32
	Firsttime int32  `orm:"null"`
	Uri       string `orm:"null"`
}
