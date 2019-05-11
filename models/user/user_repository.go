package user

import (
	"github.com/astaxie/beego/orm"
)

type UserRepository struct {
}

func (this *UserRepository) AddUser(u *User) (s string, err error) {
	dbCon := orm.NewOrm()
	if created, _, err := dbCon.ReadOrCreate(&u, "Email"); err == nil {
		if created {
			return "created new user", nil
		} else {
			return "already actived", nil
		}
	} else {
		return "create failed", err
	}
}

func (this *UserRepository) GetUser(UserId string) (u User, err error) {
	userInfo := User{Id: UserId}
	dbCon := orm.NewOrm()
	err = dbCon.Read(&userInfo)

	if err == orm.ErrNoRows || err == orm.ErrMissPK {

		return User{}, err
	}
	return userInfo, nil
}

func (this *UserRepository) GetAllUsers() map[string]*User {
	return UserList
}

//
//func (this *UserRepository) UpdateUser(uid string, uu *User) (a *User, err error) {
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

func (this *UserRepository) Login(userEmail, password string) bool {
	for _, u := range UserList {
		if u.Email == userEmail && u.Password == password {
			return true
		}
	}
	return false
}

//func DeleteUser(uid string) {
//	delete(UserList, uid)
//}
