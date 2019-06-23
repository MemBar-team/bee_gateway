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
	UserType = "userType"
	UserIat = "iat"
	Userexp = "exp"
)
type Auth struct {
	id string
	email string
	userType uint8
	iat time.Time
	exp time.Time
}

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
		Userexp: time.Now().Add(time.Hour * 24),
	})

	//電子署名
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func ParseByte(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return "", errors.New(ErrInvalid)
	}
	return []byte(secretKey), nil
}


func ParseToken(stringToken string) (Auth,error) {
	if stringToken == "" {
		return Auth{}, errors.New(ErrAbsent)
		}
	token, err := jwt.Parse(stringToken, ParseByte)

	if err != nil {
		return Auth{},errors.New(ErrOther)
	}
	if token == nil {
		return Auth{}, errors.New(ErrOther)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Auth{}, errors.New(ErrOther)
	}
	userId,ok := claims[UserID].(string)
	if !ok{
		return Auth{}, errors.New(ErrOther)
	}
	userEmail,ok := claims[UserEmail].(string)
	if !ok{
		return Auth{}, errors.New(ErrOther)
	}
	userType ,ok := claims[UserType].(uint8)
	if !ok{
		return Auth{}, errors.New(ErrOther)
	}
	userIat ,ok := claims[UserIat].(time.Time)
	if !ok{
		return Auth{}, errors.New(ErrOther)
	}
	userExp ,ok := claims[Userexp].(time.Time)
	if !ok{
		return Auth{}, errors.New(ErrOther)
	}

	return Auth {
		userId,
		userEmail,
		userType,
		userIat,
		userExp,
	},nil

}


func authenticate(stringToken string) (Auth,error) {

	// 認証
	auth, err := ParseToken(stringToken)
	if err != nil {
		return Auth{},err
	}
	now := time.Now()
	if now.After(auth.exp) {
		return Auth{}, errors.New(ErrExpired)
	}

	return auth,nil
}