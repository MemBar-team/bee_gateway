package user

import (
	"github.com/astaxie/beego/orm"
	"github.com/bee_getway/models"
	"time"
)

type UserRepository struct {
}

func (this *UserRepository) AddUser(u *User) (s string, err error) {
	//dbCon := orm.NewOrm()
	db := models.GormConnect()
	defer db.Close()
	now := time.Now()
	u.Create = &now
	db.Create(&u)
	return "created new user ", nil
}

func (this *UserRepository) SearchUser(userEmail string, password string) ( User, bool) {
	db := models.GormConnect()
	defer db.Close()
	users := []User{}
	db.Find(&users, "email=? and password=?", userEmail, password)
	totalUsers := len(users)
	if totalUsers != 1{
		return User{}, false
	}
	// usersリスト処理
	return users[0],true
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

//func (this *UserRepository) GetAllUsers() map[string]*User {
//	return UserList
//}

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

//func DeleteUser(uid string) {
//	delete(UserList, uid)
//}
