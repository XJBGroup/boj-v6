package oss

import (
	"errors"
	"github.com/Myriad-Dreamin/boj-v6/external"
)

var engine external.Engine

var ErrNotExist = errors.New("object not exists")

type byteObject []byte

func (b byteObject) Data() []byte { return b }
func (b byteObject) Free()        {}

func ToByteObject(obj []byte, err error) (external.ByteObject, error) {
	if err != nil {
		return nil, err
	}
	return byteObject(obj), nil
}

func RegisterEngine(e external.Engine) error {
	engine = e
	return nil
}
