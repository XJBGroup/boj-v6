package main

import (
	"bytes"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io/ioutil"
	"unicode"
)

func fromSnakeToCamel(src []byte, big bool) []byte {
	if len(src) == 0 {
		return []byte{}
	}
	var b = bytes.NewBuffer(make([]byte, 0, len(src)))
	for i := range src {
		if src[i] == '_' {
			big = true
		} else {
			if big {
				big = false
				b.WriteByte(byte(unicode.ToUpper(rune(src[i]))))
			} else {
				b.WriteByte(src[i])
			}
		}
	}
	return b.Bytes()
}

func shadowTemplate(src string, target string, replaceFunc func([]byte) []byte) {
	var b, err = ioutil.ReadFile(src)
	sugar.HandlerError0(err)

	b = replaceFunc(b)

	sugar.HandlerError0(ioutil.WriteFile(target, b, 0644))
}

func objectTemplate(snakeRep string, src string, target string) {
	var obj = []byte(snakeRep)
	var middleObj = bytes.ReplaceAll(obj, []byte("_"), []byte("-"))
	var entity = fromSnakeToCamel(obj, true)

	shadowTemplate(src, target, func(b []byte) []byte {
		b = bytes.ReplaceAll(b, []byte("user."), bytes.Join([][]byte{obj, []byte(".")}, []byte{}))
		b = bytes.ReplaceAll(b, []byte("package user"), bytes.Join([][]byte{[]byte("package "), obj}, []byte{}))
		b = bytes.ReplaceAll(b, []byte("/user"), bytes.Join([][]byte{[]byte("/"), middleObj}, []byte{}))
		b = bytes.ReplaceAll(b, []byte("User"), entity)

		return b
	})
}

func main() {
	//var midRep = bytes.ReplaceAll(obj, []byte("_"), []byte("-"))
	objectTemplate("announcement", "app/user/db_generated.go", "app/announcement/db_generated.go")
	objectTemplate("submission", "app/user/db_generated.go", "app/submission/db_generated.go")
	objectTemplate("problem", "app/user/db_generated.go", "app/problem/db_generated.go")
	objectTemplate("comment", "app/user/db_generated.go", "app/comment/db_generated.go")
	objectTemplate("contest", "app/user/db_generated.go", "app/contest/db_generated.go")
	objectTemplate("group", "app/user/db_generated.go", "app/group/db_generated.go")
	//objectTemplate("problem_desc", "app/user/db_generated.go", "app/problem-desc/db_generated.go")

	shadowTemplate("app/user_problem/db.go", "app/user_problem/solved_db.go", func(b []byte) []byte {
		b = bytes.ReplaceAll(b, []byte("Tried"), []byte("Solved"))
		b = bytes.ReplaceAll(b, []byte("tried"), []byte("solved"))
		return b
	})
}
