package ginhelper

//func ResetPassword(c controller.MContext, obj *user.User, password string) bool {
//	_, err := obj.ResetPassword(password)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
//			Code:  types.CodeUpdateError,
//			Error: err.Error(),
//		})
//		return false
//	}
//	return true
//}
//
//func AuthenticatePassword(c controller.MContext, user *user.User, password string) bool {
//	if ok, err := user.AuthenticatePassword(password); err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, &serial.ErrorSerializer{
//			Code:  types.CodeAuthenticatePasswordError,
//			Error: err.Error(),
//		})
//		return false
//	} else if !ok {
//		c.JSON(http.StatusOK, &serial.Response{
//			Code: types.CodeUserWrongPassword,
//		})
//		return false
//	}
//	return true
//}
