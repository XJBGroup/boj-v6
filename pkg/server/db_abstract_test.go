package server

import (
	"bytes"
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io/ioutil"
	"testing"
)

type xDB struct {}

func (x *xDB) AcceptArg(arg int) {
	fmt.Println(arg)
}

func NewDB(arg module.Module) (*xDB, error) {
	return &xDB{}, nil
}

func BadNewDB() {

}

func TestDB(t *testing.T) {
	s := &Server{}
	s.Module = make(module.Module)
	sugar.HandlerError0(InstantiateLogger()(s))
	sugar.HandlerError0(reflectCallInitDB(s, "x", ModuleInjectFunc(NewDB)))
}

func TestCheckNewDBSignature(t *testing.T) {
	t.Run("no", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				_ = err.(int)
			}
		}()
		ModuleInjectFunc(NewDB)
	})
	t.Run("yes", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				_ = err.(int)
			}
		}()
		ModuleInjectFunc(BadNewDB)
	})
}

func BenchmarkDB(b *testing.B) {
	s := &Server{}
	s.Module = make(module.Module)
	sugar.HandlerError0(InstantiateLogger()(s))
	rf := ModuleInjectFunc(NewDB)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = reflectCallInitDB(s, "x", rf)
		s.Module = make(module.Module)
	}
}

func TestReadX(t *testing.T) {
	b := bytes.Split(sugar.HandlerError(ioutil.ReadFile("db_abstract_test.go")).([]byte), []byte{'\n'})
	fmt.Println(string(b[len(b)-1]))
}