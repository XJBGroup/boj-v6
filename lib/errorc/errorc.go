package errorc

import (
	"github.com/Myriad-Dreamin/boj-v6/types"
	"reflect"
)

type Code = types.ServiceCode

func MaybeSelectError(anyObj interface{}, err error) (Code, string) {
	if err != nil {
		return types.CodeSelectError, err.Error()
	}
	if reflect.ValueOf(anyObj).IsNil() {
		return types.CodeNotFound, "not found"
	}
	return types.CodeOK, ""
}

func MaybeQueryExistenceError(exists bool, err error) (Code, string) {
	if err != nil {
		return types.CodeSelectError, err.Error()
	}
	if !exists {
		return types.CodeNotFound, "not found"
	}
	return types.CodeOK, ""
}

type UpdateFieldsable interface {
	UpdateFields(fields []string) (int64, error)
}

func UpdateFields(obj UpdateFieldsable, fields []string) (Code, string) {
	_, err := obj.UpdateFields(fields)
	if err != nil {
		return types.CodeUpdateError, err.Error()
	}
	return types.CodeOK, ""
}

type Creatable interface {
	Create() (int64, error)
}

var checkInsertError func(err error) (Code, string)

func RegisterCheckInsertError(f func(err error) (Code, string)) {
	checkInsertError = f
}

func CheckInsertError(err error) (Code, string) {
	return checkInsertError(err)
}

func CreateObj(createObj Creatable) (Code, string) {
	affected, err := createObj.Create()
	if err != nil {
		if code, errs := CheckInsertError(err); code != types.CodeOK {
			return code, errs
		}
	} else if affected == 0 {
		return types.CodeInsertError, "affect nothing"
	}
	return types.CodeOK, ""
}
