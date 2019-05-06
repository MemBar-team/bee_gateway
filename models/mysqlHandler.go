package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

	orm.RegisterModel(
		new(User))
	orm.RunSyncdb("default", false, true)

	UserList = make(map[string]*User)
	u := User{"test", 1, "z03177279@gmail.com", "test"}
	UserList["test"] = &u
}
