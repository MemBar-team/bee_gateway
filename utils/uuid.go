package utils

import (
	"github.com/satori/go.uuid"
)

type NewUUID struct {
	UUID uuid.UUID
}

func CreateUUID() string {
	newOne := NewUUID{}
	newOne.UUID = uuid.NewV4()
	u4 := newOne.UUID.String()
	return u4
}
