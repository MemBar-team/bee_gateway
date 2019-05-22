package user

import (
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*Users
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
	orm.RegisterModel(new(Users))

}
