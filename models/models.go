package models

import (
	"sync"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DB orm.Ormer
var once sync.Once

func init() {
	once.Do(func() {
		orm.Debug = false
		dbHost := beego.AppConfig.String("dbhost")
		dbPort := beego.AppConfig.String("dbport")
		dbUser := beego.AppConfig.String("dbuser")
		dbPassword := beego.AppConfig.String("dbpassword")
		dbName := beego.AppConfig.String("dbname")
		if dbPort == "" {
			dbPort = "3306"
		}
		dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
		err := orm.RegisterDriver("mysql", orm.DRMySQL)
		if err != nil {
			logs.Error("orm.RegisterDriver err:", err)
		}
		err2 := orm.RegisterDataBase("default", "mysql", dsn)
		orm.SetMaxIdleConns("default", 1000)
		orm.SetMaxOpenConns("default", 2000)
		if err2 != nil {
			logs.Error("orm.RegisterDataBase err:", err2)
		}

		orm.RegisterModel(new(Users), new(Topic), new(SourceCode),
			new(Solution), new(Sim), new(ShareCode), new(Reply),
			new(Problem), new(Privilege), new(Printer), new(Online),
			new(Mail), new(News), new(LoginLog), new(Custominput),
			new(ContestProblem), new(Contest), new(Compileinfo), new(Balloon))
		_ = orm.RunSyncdb("default", false, true)

		DB = orm.NewOrm()
		err = DB.Using("default") // 默认使用 hgoj，你可以指定为其他数据库
		if err != nil {
			logs.Error("orm.NewOrm err:", err)
			//os.Exit(-1)
		}
	})

}
