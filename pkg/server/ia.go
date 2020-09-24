package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"reflect"
)

var et = reflect.TypeOf(new(error)).Elem()
var mt = reflect.TypeOf(new(module.Module)).Elem()

func ModuleInjectFunc(f interface{}) reflect.Value {
	rf := reflect.ValueOf(f)
	t := rf.Type()
	if t.NumIn() != 1 {
		panic(fmt.Errorf("module injector must has the only module object as input, got num in %v", t.NumIn()))
	}
	if t.In(0) != mt {
		panic(fmt.Errorf("module injector must has input signature shape: (module.Module), got (%v)", t.In(0)))
	}
	if t.NumOut() != 2 {
		panic(fmt.Errorf("module injector must has return signature shape: (object, error), got num out %v", t.NumOut()))
	}
	if t.Out(1) != et {
		panic(fmt.Errorf("module injector must has return signature shape: (object, error), got (%v, %v)", t.Out(0), t.Out(1)))
	}
	return rf
}

type InitializeAction = func(srv *Server) error

func (srv *Server) applyInitializeAction(actions []InitializeAction) error {
	for _, ia := range actions {
		err := ia(srv)
		if err != nil {
			return err
		}
	}
	return nil
}
