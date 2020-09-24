package snippet

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/errorc"
	"github.com/Myriad-Dreamin/boj-v6/lib/jwt"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/tidwall/gjson"
	"net/http"
	"reflect"
	"strconv"
)

var ResponseOK = serial.Response{Code: types.CodeOK}

func CheckInsertError(c controller.MContext, checker func(err error) types.ServiceCode, err error) bool {
	if code := checker(err); code != types.CodeOK {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{Code: code})
		return true
	}
	return false
}

func MissID(c controller.MContext) {
	c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
		Code:   types.CodeInvalidParameters,
		ErrorS: "id missing in the path",
	})
}

func AuthFailed(c controller.MContext, errorString string) {
	c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
		Code:   types.CodeAuthenticatePolicyError,
		ErrorS: errorString,
	})
}

func ParseUint(c controller.MContext, key string) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: "bad negative id",
		})
		return 0, false
	}
	return uint(id), true
}

func BindRequest(c controller.MContext, req interface{}) bool {
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: err.Error(),
		})
		return false
	}
	return true
}

func RawJson(c controller.MContext) (gjson.Result, bool) {
	if b, err := c.GetRawData(); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: err.Error(),
		})
		return gjson.Result{}, false
	} else {
		return gjson.ParseBytes(b), true
	}
}

func ParseUintAndBind(c controller.MContext, key string, req interface{}) (uint, bool) {
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: err.Error(),
		})
		return 0, false
	}
	if id < 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: "bad negative id",
		})
		return 0, false
	}
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: err.Error(),
		})
		return 0, false
	}
	return uint(id), true
}

func RosolvePageVariable(c controller.MContext) (int, int, bool) {
	spage, ok := c.GetQuery("page")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: "missing page number",
		})
		return 0, 0, false
	}
	page, err := strconv.Atoi(spage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUnserializeDataError,
			ErrorS: "can not convert page number to integer",
		})
		return 0, 0, false
	}
	spageSize, ok := c.GetQuery("page_size")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: "missing page size",
		})
		return 0, 0, false
	}
	pageSize, err := strconv.Atoi(spageSize)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUnserializeDataError,
			ErrorS: "can not convert page size to integer",
		})
		return 0, 0, false
	}
	if page <= 0 || pageSize <= 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInvalidParameters,
			ErrorS: "bad negative params",
		})
		return 0, 0, false
	}
	return page, pageSize, true
}

func MaybeGetRawDataError(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeGetRawDataError,
			ErrorS: err.Error(),
		})
		return true
	}
	return false
}

func MaybeCountError(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeSelectError,
			ErrorS: err.Error(),
		})
		return true
	}

	return false
}

type applyContext struct{ controller.MContext }

func (ctx applyContext) applyError(code errorc.Code, errs string) bool {
	if code != types.CodeOK {
		ctx.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code:   code,
			ErrorS: errs,
		})
		return true
	}
	return false
}

func MaybeSelectError(c controller.MContext, anyObj interface{}, err error) bool {
	return applyContext{c}.applyError(errorc.MaybeSelectError(anyObj, err))
}

func MaybeQueryExistenceError(c controller.MContext, exists bool, err error) bool {
	return applyContext{c}.applyError(errorc.MaybeQueryExistenceError(exists, err))
}

func DoReport(c controller.MContext, err error) bool {
	if err != nil {
		if _, ok := err.(*serial.ErrorSerializer); !ok {
			if _, ok = err.(*serial.Response); !ok {
				panic(err)
			}
		}

		c.AbortWithStatusJSON(http.StatusOK, err)
	}
	return err != nil
}

func MaybeSelectErrorWithTip_(anyObj interface{}, err error, missError string) error {
	if err != nil {
		return &serial.ErrorSerializer{
			Code:   types.CodeSelectError,
			ErrorS: err.Error(),
		}
	}
	if reflect.ValueOf(anyObj).IsNil() {
		return &serial.ErrorSerializer{
			Code:   types.CodeNotFound,
			ErrorS: missError,
		}
	}

	return nil
}

func MaybeSelectErrorWithTip(c controller.MContext, anyObj interface{}, err error, missError string) bool {
	return DoReport(c, MaybeSelectErrorWithTip_(anyObj, err, missError))
}

func MaybeMissingError(c controller.MContext, has bool, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeSelectError,
			ErrorS: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &serial.Response{
			Code: types.CodeNotFound,
		})
		return true
	}

	return false
}

func MaybeMissingErrorWithTip(c controller.MContext, has bool, err error, missError string) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeSelectError,
			ErrorS: err.Error(),
		})
		return true
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeNotFound,
			ErrorS: missError,
		})
		return true
	}

	return false
}
func MaybeOnlySelectError(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeSelectError,
			ErrorS: err.Error(),
		})
		return true
	}

	return false
}

type Deletable interface {
	Delete() (int64, error)
}

func DeleteObj(c controller.MContext, affected int64, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeDeleteError,
			ErrorS: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code: types.CodeDeleteNoEffect,
		})
		return false
	}
	return true
}

func CreateObj(c controller.MContext, checker func(err error) types.ServiceCode, affected int64, err error) bool {
	if err != nil {
		if CheckInsertError(c, checker, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code: types.CodeInsertError,
		})
		return false
	}
	return true
}

func CreateObjWithTip(c controller.MContext, checker func(err error) types.ServiceCode, affected int64, err error, tip string) bool {
	if err != nil {
		if CheckInsertError(c, checker, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: "create " + tip + " failed: " + err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeInsertError,
			ErrorS: "create " + tip + " has no effect",
		})
		return false
	}
	return true

}

type Updatable interface {
	Update() (int64, error)
}

func UpdateObj(c controller.MContext, checker func(err error) types.ServiceCode, updateObj Updatable) bool {
	affected, err := updateObj.Update()
	if err != nil {
		if CheckInsertError(c, checker, err) {
			return false
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeUpdateError,
			ErrorS: err.Error(),
		})
		return false
	} else if affected == 0 {
		c.AbortWithStatusJSON(http.StatusOK, &serial.Response{
			Code: types.CodeUpdateError,
		})
		return false
	}
	return true
}

func UpdateFields(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
			Code:   types.CodeUpdateError,
			ErrorS: err.Error(),
		})
		return false
	}
	return true
}

func GetCustomFields(c controller.MContext) *types.CustomFields {
	claims, _ := c.Get("claims")
	return claims.(*jwt.CustomClaims).CustomField.(*types.CustomFields)
}

func DoRollback(rollbacks []func()) {
	var l = len(rollbacks)
	for i := l - 1; i >= 0; i-- {
		rollbacks[i]()
	}
	return
}
