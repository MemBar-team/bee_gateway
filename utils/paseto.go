package utils
import (
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/bee_gateway/models/entity"
	"github.com/o1egl/paseto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
	"strconv"
	"time"
)

var (
	privateKey []byte
	publicKey []byte
)

type Auth struct {
	id string
	email string
	userType uint8
	iat time.Time
	exp time.Time
}

const (
	PrivateKEY string = "privateKey"
	PublicKKEY string = "publicKey"
	ErrAbsent  string = "token absent"  // token存在しない
	ErrInvalid string = "token invalid" // token無効
	ErrExpired string = "token expired" // token機関超過
	ErrOther   string = "other error"   // token未知エラー
	ErrCreateFail string = "token create fail" //token作成失敗
	UserID = "id"
	UserEmail = "email"
	UserType = "userType"
	UserIat = "iat"
	Userexp = "exp"
)


func init() {
	pvKey, _ := hex.DecodeString(beego.AppConfig.String(PrivateKEY))
	privateKey = ed25519.PrivateKey(pvKey)

	pbKey, _ := hex.DecodeString(beego.AppConfig.String(PublicKKEY))
	publicKey = ed25519.PublicKey(pbKey)
}

func CreateToken(user entity.User) (string,error){

	//　json作成
	jsonToken := paseto.JSONToken{
		Expiration:time.Now().Add(24* time.Hour),
	}
	// クライム追加
	jsonToken.Set(UserID,user.Id)
	jsonToken.Set(UserEmail,user.Email)
	jsonToken.Set(UserType, string(*user.UserType))
	jsonToken.Set(UserIat,time.Now().String())
	jsonToken.Set(UserEmail,time.Now().Add(time.Hour * 24).String())
	footer := "dmason"

	// トークン作成
	v2 := paseto.NewV2()

	// トークンサイン
	token, err := v2.Sign(privateKey, jsonToken, footer)
	if err != nil {
		return "",errors.New(ErrCreateFail)
	}

	return token, nil
}