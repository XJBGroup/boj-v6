package mcore

import (
	"database/sql"
	core_cfg "github.com/Myriad-Dreamin/boj-v6/lib/core-cfg"
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type DormModule struct {
	DormDB *dorm.DB
}

func (m *DormModule) FromRaw(db *dorm.DB, dep module.Module) bool {
	m.DormDB = db
	dep.Provide(DefaultNamespace.DBInstance.DormDB, m.DormDB)
	return true
}

func (m *DormModule) FromRawSQL(db *sql.DB, dep module.Module) bool {
	instanceLogger := dep.RequireImpl(new(logger.Logger)).(logger.Logger)
	options := []interface{}{adapt(instanceLogger)}

	escaper := m.getEscaper(dep)
	if len(escaper) != 0 {
		options = append(options, dorm.Escaper(escaper))
	}

	xdb, err := dorm.FromRaw(db, options)

	m.DormDB = xdb
	dep.Provide(DefaultNamespace.DBInstance.DormDB, xdb)

	return Maybe(dep, "init dorm error", err)
}

func (m *DormModule) FromContext(dep module.Module) bool {
	m.DormDB = dep.RequireImpl(new(*dorm.DB)).(*dorm.DB)
	return true
}

func (m *DormModule) Install(dep module.Module) bool {
	return m.FromContext(dep)
}

func (m *DormModule) GetDormInstance() *dorm.DB {
	return m.DormDB
}

type DatabaseConfiguration interface {
	GetDatabaseConfiguration() core_cfg.DatabaseConfig
}

type RedisConfiguration interface {
	GetRedisConfiguration() core_cfg.RedisConfig
}

func (m *DormModule) getEscaper(dep module.Module) string {
	return dep.RequireImpl(new(DatabaseConfiguration)).(DatabaseConfiguration).GetDatabaseConfiguration().Escaper
}

type L struct {
	logger.Logger
}

func (l *L) With(kvs ...interface{}) dorm.Logger {
	return &L{l.Logger.With(kvs)}
}

func adapt(logger logger.Logger) dorm.Logger {
	return &L{logger}
}
