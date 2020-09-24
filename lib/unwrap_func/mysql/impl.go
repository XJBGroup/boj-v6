package unwrap_func_sqlite

import (
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/go-sql-driver/mysql"
	"path"
)

func UnwrapMySQLError(err error) types.ServiceCode {
	if sqlError, ok := err.(*mysql.MySQLError); ok {
		switch sqlError.Number {
		case 1062:
			return types.CodeDuplicatePrimaryKey
		case 1366:
			return types.CodeDatabaseIncorrectStringValue
		default:
			return types.CodeNotRecognizedDatabaseError + types.ServiceCode(sqlError.Number)
		}
	}

	if err != nil {
		return types.CodeDatabaseError
	}

	return types.CodeOK
}

func Register(m module.Module) {
	m.Provide(path.Join("database_unwrap_functions", "mysql"), UnwrapMySQLError)
}
