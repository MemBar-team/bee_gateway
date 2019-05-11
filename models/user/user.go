package user

import (
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

type User struct {
	Id       string `orm:"pk;" json:"id"`
	UserType uint8  `json:"user_type"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Modified time.Time
	Create   time.Time
}

func init() {
	orm.RegisterModel(
		new(User))
}
