package types

type ServiceCode = int

const (
	// Generic Code

	CodeOK ServiceCode = iota
	// CodeBindError indicates a parameter missing error
	CodeBindError
	// CodeUnserializeDataError indicates a parsing data error
	CodeUnserializeDataError
	// CodeInvalidParameters tells some wrong data was in the request
	CodeInvalidParameters
	// GetRawDataError tells some wrong data was in the request
	CodeGetRawDataError
)

const (
	// Generic Code -- Database
	// CodeInsertError occurs when insert object into database
	CodeInsertError ServiceCode = iota + 100
	// CodeSelectError occurs when select object from database
	CodeSelectError
	// CodeUpdateError occurs when update object to database
	CodeUpdateError
	// CodeDeleteError occurs when delete object from database
	CodeDeleteError
	// CodeNotFound occurs when object with specific condition is not in the
	CodeNotFound
	// CodeDeleteNoEffect occurs when deleting object has no effect
	CodeDeleteNoEffect

	// database
	// CodeDuplicatePrimaryKey occurs when the object's primary key conflicts
	// with something that was already in the database
	CodeDuplicatePrimaryKey

	// CodeDatabaseIncorrectStringValue occurs when ?
	CodeDatabaseIncorrectStringValue

	CodeUniqueConstraintFailed
)

const (
	// Generic Code -- Authentication
	// CodeAuthGenerateTokenError occurs when insert object into database
	CodeAuthGenerateTokenError ServiceCode = iota + 1000
	CodeAuthenticatePasswordError
	CodeAuthenticatePolicyError

	CodeChangeOwnerError
	CodeGroupCreateError
)

const (
	CodeUserIDMissing ServiceCode = iota + 10000
	CodeUserWrongPassword
)

const (
	CodeSubmissionUploaded ServiceCode = iota + 11000
	CodeFSExecError
	CodeUploadFileError
	CodeConfigModifyError
	CodeStatError
)
