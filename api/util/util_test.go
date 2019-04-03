package util

import (
	"fmt"
	"testing"
)

func TestNewuuid(t *testing.T) {
	uid := Newuuid()
	fmt.Println(uid)
}
