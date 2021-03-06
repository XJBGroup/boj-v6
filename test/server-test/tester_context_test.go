package server_test

import (
	"github.com/Myriad-Dreamin/boj-v6/test/tester"
	"reflect"
)

var srv *tester.Tester

const (
	normalUserIdKey       = "user/normal/key"
	normalUserPassword    = "yY11112222"
	normalUserNewPassword = "xX11122222"

	AnnouncementIdKey = "announcement/normal/key"
)

var intT = 1
var intType = reflect.TypeOf(intT)
var uintType = reflect.TypeOf(uint(1))

func RangeInt(l, r int) []int {
	var x = make([]int, r-l)
	for i := l; i < r; i++ {
		x[i-l] = i
	}
	return x
}
