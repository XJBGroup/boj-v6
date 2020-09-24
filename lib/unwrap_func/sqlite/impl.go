package unwrap_func_sqlite

import (
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/mattn/go-sqlite3"
	"path"
)

func UnwrapSqliteError(err error) types.ServiceCode {
	if sqlError, ok := err.(sqlite3.Error); ok {
		switch sqlError.ExtendedCode {
		case 1062:
			return types.CodeDuplicatePrimaryKey
		case 1366:
			return types.CodeDatabaseIncorrectStringValue
		case 2067:
			return types.CodeUniqueConstraintFailed
		default:
			return types.CodeNotRecognizedDatabaseError + types.ServiceCode(sqlError.ExtendedCode)
		}
	}

	if err != nil {
		return types.CodeDatabaseError
	}

	return types.CodeOK
}

func Register(m module.Module) {
	m.Provide(path.Join("database_unwrap_functions", "sqlite3"), UnwrapSqliteError)
}
