package common

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/satori/go.uuid"
)

type NewUUID struct {
	UUID uuid.UUID
}

func CreateUUID() string {
	newOne := NewUUID{}
	newOne.UUID = uuid.NewV4()
	u4 := newOne.UUID.String()
	spew.Dump(u4)
	return u4
}
