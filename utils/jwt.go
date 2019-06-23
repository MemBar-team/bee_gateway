package utils

import (
	"github.com/astaxie/beego"
	"github.com/bee_gateway/models/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateJWT(user entity.User) string {

	secretKey := beego.AppConfig.String("secretKey")

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["userType"] = user.UserType
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	//電子署名
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}