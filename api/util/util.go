package util

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func Newuuid() string {
	uid := uuid.Must(uuid.NewV4())
	return fmt.Sprintf("%s", uid)
}