package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/bee_getway/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dblink := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbname"),
	)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("dbdriver"), dblink)

	orm.RunSyncdb("default", false, true)

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
