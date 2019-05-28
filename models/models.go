package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
var DB orm.Ormer
func init() {
	orm.Debug = true
	dbHost := beego.AppConfig.String("dbhost")
	dbPort := beego.AppConfig.String("dbport")
	dbUser := beego.AppConfig.String("dbuser")
	dbPassword := beego.AppConfig.String("dbpassword")
	dbName := beego.AppConfig.String("dbname")
	if dbPort == "" {
		dbPort = "3306"
	}
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)


	orm.RegisterModel(new(Users), new(Topic), new(SourceCode),
		new(Solution), new(Sim), new(ShareCode), new(Reply),
		new(Problem), new(Privilege), new(Printer), new(Online),
		new(Mail), new(News), new(LoginLog), new(Custominput),
		new(ContestProblem), new(Contest), new(CompileInfo), new(Balloon))
	_ = orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()
	DB.Using("default") // 默认使用 default，你可以指定为其他数据库

}