package user

import (
	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"time"

	"github.com/bee_getway/models"
)

type User struct {
	Id       string `gorm:"pk;unique;"`
	UserType uint8
	Email    string `gorm:"unique"`
	Password string
	Modified time.Time
	Created  time.Time
}

func init() {
	devmode, err := beego.AppConfig.Bool("devmode")
	if err != nil {
		panic(err.Error())
	}
	if devmode {
		db := models.GormConnect()
		spew.Dump("created table Users")
		db.AutoMigrate(&User{})
	}
}
