package utils

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/bee_gateway/models/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	UserID = "id"
	UserEmail = "email"
	UserType = "usertype"
	UserIat = "iat"
	Userexp = "exp"
)

var (
	secretKey string
	ErrAbsent  = "token absent"  // token存在しない
	ErrInvalid = "token invalid" // token無効
	ErrExpired = "token expired" // token機関超過
	ErrOther   = "other error"   // token未知エラー
)
func init() {
	secretKey = beego.AppConfig.String("secretKey")
}

func CreateJWT(user entity.User) string {

	// headerのセット
	// claimsのセット
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		UserID: user.Id,
		UserEmail: user.Email,
		UserType: user.UserType,
		UserIat: time.Now(),
		Userexp: time.Now().Add(time.Hour * 24).Unix(),
	})

	//電子署名
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func ParseByte(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return "", errors.New(ErrOther)
	}
	return []byte(secretKey), nil
}


func ParseToken(stringToken string) (entity.User,error) {
	if stringToken == "" {
		return entity.User{}, errors.New(ErrAbsent)
		}
	token, err := jwt.Parse(stringToken, ParseByte)

	if err != nil {
		return entity.User{},errors.New(ErrOther)
	}
	if token == nil {
		return entity.User{}, errors.New(ErrOther)
	}

	return entity.User{},nil

}