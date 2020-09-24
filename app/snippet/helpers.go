package snippet

func ConvertErrorToString(err error) string {
	var errString string
	if err != nil {
		errString = err.Error()
	} else {
		errString = ""
	}
	return errString
}
