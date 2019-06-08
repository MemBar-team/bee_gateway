package user

import (
	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"time"

	"github.com/bee_getway/models"
)

type User struct {
	Id       string `gorm:"pk;unique;" json:"id" gorm:"pk;unique;"`
	UserType *uint8 `json:"user_type"`
	Email    string `gorm:"unique;" json:"email" gorm:"unique"`
	Password string `json:"password"`
	Modified *time.Time
	Created  *time.Time
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
