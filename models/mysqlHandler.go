package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
)

func init(){
	dblink:=fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbport"),
		beego.AppConfig.String("dbname"),
		beego.AppConfig.String("psqlssl"),
	)
	spew.Dump(dblink)
	orm.RegisterDriver("postgres",orm.DRPostgres)
	orm.RegisterDataBase(beego.AppConfig.String("dbalianame"),beego.AppConfig.String("dbdrivename"), dblink)
	//orm.RegisterDataBase("default", "postgres", "user=postgres password=test dbname=test1 host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterModel(
		new(Comment))
	orm.RunSyncdb("default",false,true)
}