package userRepository

import (
	"github.com/astaxie/beego/orm"
	"github.com/bee_gateway/models"
	"github.com/bee_gateway/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/bee_gateway/models/entity"
	"time"
)

type UserRepository struct {
}

func (this *UserRepository) AddUser(u *entity.User) (s string, err error) {
	db := models.GormConnect()
	defer db.Close()
	now := time.Now()
	u.Created = &now
	u.Id = utils.CreateUUID()
	db.Create(&u)
	return "created new user ", nil
}

func (this *UserRepository) UserLogin(userEmail string, password string) (entity.User, bool) {
	db := models.GormConnect()
	defer db.Close()
	users := []entity.User{}
	db.Find(&users, "email=? and password=?", userEmail, password)
	totalUsers := len(users)
	if totalUsers != 1 {
		return entity.User{}, false
	}
	// usersリスト処理

	spew.Dump(users[0])
	return users[0], true
}

func (this *UserRepository) GetUser(UserId string) (u entity.User, err error) {
	userInfo := entity.User{Id: UserId}
	dbCon := orm.NewOrm()
	err = dbCon.Read(&userInfo)

	if err == orm.ErrNoRows || err == orm.ErrMissPK {

		return entity.User{}, err
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
