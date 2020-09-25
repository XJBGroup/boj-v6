package snippet

import (
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func ResetPassword(c controller.MContext, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUpdateError,
			ErrorS: err.Error(),
		})
		return false
	}
	return true
}

func AuthenticatePassword(c controller.MContext, ok bool, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeAuthenticatePasswordError,
			ErrorS: err.Error(),
		})
		return false
	} else if !ok {
		c.JSON(http.StatusOK, &serial.Response{
			Code: types.CodeUserWrongPassword,
		})
		return false
	}
	return true
}
