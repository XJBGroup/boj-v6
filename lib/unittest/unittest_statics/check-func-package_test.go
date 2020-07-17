package unittest_statics

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/unittest/unittest_types"
	"reflect"
	"testing"
)

func Test_RepositionCtx(t *testing.T) {
	var p = &LinkedContext{name: "root"}
	a := &LinkedContext{name: "a"}
	p.Insert("a", a)
	json := &LinkedContext{name: "json"}
	p.Insert("json", json)
	aa := &LinkedContext{name: "a"}
	a.Insert("a", aa)
	aaa := &LinkedContext{name: "a"}
	aa.Insert("a", aaa)

	type args struct {
		p unittest_types.LinkedContext
		k string
	}
	tests := []struct {
		name string
		args args
		want unittest_types.LinkedContext
	}{
		{"easy", args{a, "a"}, a},
		{"easy", args{aa, "a"}, aa},
		{"easy", args{aaa, "a"}, aaa},
		{"easy", args{a, ".a"}, a},
		{"easy", args{aa, ".a"}, a},
		{"easy", args{aaa, ".a"}, a},
		{"easy", args{a, ".a.a"}, aa},
		{"easy", args{aa, ".a.a"}, aa},
		{"easy", args{aaa, ".a.a"}, aa},
		{"easy", args{a, ".a.a.a"}, aaa},
		{"easy", args{aa, ".a.a.a"}, aaa},
		{"easy", args{aaa, ".a.a.a"}, aaa},
		{"easy", args{a, "a.a"}, aa},
		{"easy", args{aa, "a.a"}, aaa},
		{"easy", args{aaa, "a.a"}, nil},
		{"easy", args{a, "a.a.a"}, aaa},
		{"easy", args{aa, "a.a.a"}, nil},
		{"easy", args{aaa, "a.a.a"}, nil},
		{"global_find", args{p, ".json"}, json},
		{"global_find", args{a, ".json"}, json},
		{"global_find", args{aa, ".json"}, json},
		{"global_find", args{aaa, ".json"}, json},
		{"global_find", args{p, "root.json"}, json},
		{"global_find", args{a, "root.json"}, json},
		{"global_find", args{aa, "root.json"}, json},
		{"global_find", args{aaa, "root.json"}, json},
		{"global_find", args{p, "json"}, nil},
		{"global_find", args{a, "json"}, nil},
		{"global_find", args{aa, "json"}, nil},
		{"global_find", args{aaa, "json"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepositionCtx(tt.args.p, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RepositionCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
