package user

import (
	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"time"

	"github.com/bee_getway/models"
)

type User struct {
	Id       string `gorm:"pk;unique;" json:"id"`
	UserType uint8  `json:"user_type"`
	Email    string `gorm:"unique;" json:"email"`
	Password string `json:"password"`
	Modified *time.Time
	Create   *time.Time
}

func init() {
	devmode, err := beego.AppConfig.Bool("devmode")
	if err != nil {
		panic(err.Error())
	}
	if devmode {
		db := models.GormConnect()
		if err != nil {
			panic(err.Error())
		}
		spew.Dump("created table Users")
		db.AutoMigrate(&User{})
	}
}
