package common

import (
	"github.com/bee_gateway/vendor/github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
)

func CreateUUID() string {
	u,err := uuid.NewRandom()
	if err != nil {
		beego.BeeLogger.Error("error")
	}
	spew.Dump(u)

	return uu
}