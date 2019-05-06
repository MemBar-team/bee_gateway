package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dblink := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbname"),
	)
	spew.Dump(dblink)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("dbdriver"), dblink)
	//orm.RegisterDataBase("default", "postgres", "user=postgres password=test dbname=test1 host=127.0.0.1 port=5432")
	orm.RegisterModel(
		new(User))
	orm.RunSyncdb("default", false, true)
}
