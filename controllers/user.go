package controllers

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"time"

	"github.com/astaxie/beego"
	"github.com/bee_getway/models/user"
)

// Operations about Users
type UserController struct {
	beego.Controller
	user.UserRepository
}

type User struct {
	Id       string `gorm:"pk;unique;" json:"id"`
	UserType uint8  `json:"user_type"`
	Email    string `gorm:"unique;" json:"email"`
	Password string `json:"password"`
	Modified *time.Time
	Created  *time.Time
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) CreateUser() {
	var userData User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &userData)
	if err != nil {
		panic("jsont to obj is faild")
	}
	strMsg, err := u.AddUser(&userData)
	if err != nil {
		panic("db insert user failed")
	}
	u.Data["json"] = strMsg
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var logindata LoginData
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &logindata)
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(logindata)
	userData, ok := u.UserLogin(logindata.Email, logindata.Password)
	if ok {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	spew.Dump(userData)
	//jwt.
	//	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
//func (u *UserController) Get() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		user, err := u.GetUser(uid)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = user
//		}
//	}
//	u.ServeJSON()
//}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
//func (u *UserController) Put() {
//	uid := u.GetString(":uid")
//	if uid != "" {
//		var user models.User
//		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//		uu, err := models.UpdateUser(uid, &user)
//		if err != nil {
//			u.Data["json"] = err.Error()
//		} else {
//			u.Data["json"] = uu
//		}
//	}
//	u.ServeJSON()
//}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
//func (u *UserController) Delete() {
//	uid := u.GetString(":uid")
//	models.DeleteUser(uid)
//	u.Data["json"] = "delete success!"
//	u.ServeJSON()
//}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
