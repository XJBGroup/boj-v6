package db

import "github.com/Myriad-Dreamin/boj-v6/types"

type BasicDB interface {
	UnwrapError(err error) (code types.ServiceCode)
}
