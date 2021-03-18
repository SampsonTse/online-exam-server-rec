package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "online-exam-server-rec/routers"
)

func init() {
	// connect mysql
	orm.RegisterModel("default", "mysql", beego.AppConfig.String("sqlconn"))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
	orm.Debug = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
