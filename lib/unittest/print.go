package unittest

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io"
	"reflect"
	"strings"
)

func writeLevel(w io.Writer, level int) {
	for i := 0; i < level; i++ {
		sugar.HandlerError(w.Write([]byte{' '}))
	}
}

type defaultNode []map[string]interface{}

type metaNode map[string]interface{}

type usingNode map[string]string

type usingForceNode map[string]string

type selectorNode []SelectorDef

type caseNode map[string]interface{}

type versionNode string

type nameNode string

type inheritNode []string

func debugPrint(w io.Writer, v interface{}, path string, level int) {
	var mm = false
	switch v := v.(type) {
	case defaultNode:
		for i := range v {
			debugPrint(w, v[i], fmt.Sprintf("Default(%s)", dotJoin(path, "default")), level)
		}
	case usingNode:
		debugPrint(w, (map[string]string)(v), fmt.Sprintf("Using(%s)", path), level)
	case usingForceNode:
		debugPrint(w, (map[string]string)(v), fmt.Sprintf("UsingFor(%s)", path), level)
	case metaNode:
		debugPrint(w, (map[string]interface{})(v), fmt.Sprintf("Meta(%s)", dotJoin(path, "meta")), level)
	case versionNode:
		debugPrint(w, (string)(v), dotJoin(path, "version"), level)
	case nameNode:
		debugPrint(w, (string)(v), dotJoin(path, "name"), level)
	default:
		mm = true
	}
	if !mm {
		return
	}
	k := reflect.TypeOf(v).Kind()
	if (k == reflect.Ptr || k == reflect.Map) && reflect.ValueOf(v).IsNil() {
		return
	}
	writeLevel(w, level)
	switch v := v.(type) {
	case SpecV1:
		path = dotJoin(path, "spec")
		sugar.HandlerError(fmt.Fprintf(w, "Spec(%s):\n", path))
		debugPrint(w, versionNode(v.Version), path, level+1)
		if v.Meta != nil {
			sugar.HandlerError(w.Write([]byte{'\n'}))
		}
		debugPrint(w, v.Meta, path, level+1)
		if v.TestDefs != nil {
			sugar.HandlerError(w.Write([]byte{'\n'}))
		}
		debugPrint(w, v.TestDefs, path, level+1)
		debugPrint(w, defaultNode(v.Default), path, level+1)
		if v.Selector != nil {
			sugar.HandlerError(w.Write([]byte{'\n'}))
		}
		debugPrint(w, selectorNode(v.Selector), path, level+1)
		sugar.HandlerError(w.Write([]byte{'\n'}))
		debugPrint(w, v.PackageDefs, path, level+1)
		sugar.HandlerError(w.Write([]byte{'\n'}))
		debugPrint(w, metaNode(v.Meta), path, level+1)
	case selectorNode:
		sugar.HandlerError(w.Write([]byte("Selectors:")))
		for i := range v {
			sugar.HandlerError(w.Write([]byte{'\n'}))
			debugPrint(w, v[i], fmt.Sprintf("[%v]", i), level+1)
		}
	case SelectorDef:
		sugar.HandlerError(w.Write([]byte("Selector:\n")))
		debugPrint(w, nameNode(v.Name), path, level+1)
		if v.Case != nil {
			sugar.HandlerError(w.Write([]byte{'\n'}))
		}
		debugPrint(w, caseNode(v.Case), path, level+1)
	case []PackageDef:
		sugar.HandlerError(fmt.Fprintf(w, "PackageDefs:"))
		for i := range v {
			sugar.HandlerError(w.Write([]byte{'\n'}))
			debugPrint(w, v[i], path, level+1)
		}
	case []TestDef:
		sugar.HandlerError(fmt.Fprintf(w, "TestDefs:"))
		for i := range v {
			sugar.HandlerError(w.Write([]byte{'\n'}))
			debugPrint(w, v[i], path, level+1)
		}
	case PackageDef:
		sugar.HandlerError(fmt.Fprintf(w, "%s: %s", v.Path, v.Namespace))
	case TestDef:
		path = dotJoin(path, v.Name)
		sugar.HandlerError(fmt.Fprintf(w, "%s:", path))
		if v.Using != nil {
			sugar.HandlerError(w.Write([]byte{'\n'}))
		}
		debugPrint(w, usingNode(v.Using), path, level+1)
		if v.Meta != nil {
			sugar.HandlerError(w.Write([]byte{'\n'}))
		}
		debugPrint(w, metaNode(v.Meta), path, level+1)
		sugar.HandlerError(w.Write([]byte{'\n'}))
		debugPrint(w, inheritNode(v.Inherit), path, level+1)
		sugar.HandlerError(w.Write([]byte{'\n'}))
		debugPrint(w, v.Cases, path, level+1)
	case inheritNode:
		sugar.HandlerError(fmt.Fprintf(w, "Inherit: <%s>", strings.Join(v, ",")))
	case string:
		if len(path) != 0 {
			sugar.HandlerError(w.Write([]byte(path)))
			sugar.HandlerError(w.Write([]byte{':'}))
			sugar.HandlerError(w.Write([]byte{' '}))
		}
		sugar.HandlerError(w.Write([]byte(v)))
	case map[string]interface{}:
		sugar.HandlerError(w.Write([]byte(path)))
		sugar.HandlerError(w.Write([]byte{':'}))
		for k := range v {
			sugar.HandlerError(w.Write([]byte{'\n'}))
			debugPrint(w, v[k], k, level+1)
		}
	case map[string]string:
		sugar.HandlerError(w.Write([]byte(path)))
		sugar.HandlerError(w.Write([]byte{':'}))
		for k := range v {
			sugar.HandlerError(w.Write([]byte{'\n'}))
			debugPrint(w, v[k], k, level+1)
		}
	case map[interface{}]interface{}:
		sugar.HandlerError(w.Write([]byte(path)))
		sugar.HandlerError(w.Write([]byte{':'}))
		for k := range v {
			sugar.HandlerError(w.Write([]byte{'\n'}))
			debugPrint(w, v[k], k.(string), level+1)
		}
	default:
		panic(fmt.Sprintf("unknown type: %t", v))
	}
}
