package user

import (
	"time"
)

type Users struct {
	Id       string `orm:"pk;unique;" json:"id"`
	UserType uint8  `json:"user_type"`
	Email    string `orm:"unique;" json:"email"`
	Password string `json:"password"`
	Modified time.Time
	Create   time.Time
}

func init() {
	//orm.RegisterModel(new(Users))
	db := gormConnect()

}
