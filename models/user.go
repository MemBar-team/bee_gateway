package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"test", 1, "z03177279@gmail.com", "test"}
	UserList["test"] = &u
}

type User struct {
	UserID       string
	UserType     uint8
	UserEmail    string
	UserPassword string
}

func AddUser(u User) string {
	u.UserID = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.UserID] = &u
	return u.UserID
}

func GetUser(UserId string) (u *User, err error) {
	if u, ok := UserList[UserId]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

//func GetAllUsers() map[string]*User {
//	return UserList
//}

//
//func UpdateUser(uid string, uu *User) (a *User, err error) {
//	if u, ok := UserList[uid]; ok {
//		if uu.Username != "" {
//			u.Username = uu.Username
//		}
//		if uu.Password != "" {
//			u.Password = uu.Password
//		}
//		return u, nil
//	}
//	return nil, errors.New("User Not Exist")
//}

func Login(userEmail, password string) bool {
	for _, u := range UserList {
		if u.UserEmail == userEmail && u.UserPassword == password {
			return true
		}
	}
	return false
}

//func DeleteUser(uid string) {
//	delete(UserList, uid)
//}
