package database

import (
	mcore "github.com/Myriad-Dreamin/boj-v6/lib/core"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Module struct {
	mcore.GormModule
	mcore.RawSQLModule
	mcore.DormModule
	mcore.LoggerModule

	Opened bool
}

func NewModule() Module {
	return Module{
		Opened: false,
	}
}

func (m *Module) install(
	initFunc func(dep module.Module) bool, dep module.Module) bool {
	return m.LoggerModule.Install(dep) &&
		initFunc(dep) &&
		m.RawSQLModule.FromRaw(m.GormDB.DB(), dep) &&
		m.DormModule.FromRawSQL(m.RawDB, dep)
}

func (m *Module) InstallFromContext(dep module.Module) bool {
	m.Opened = m.install(m.GormModule.FromContext, dep)
	return m.Opened
}

func (m *Module) Install(dep module.Module) bool {
	m.Opened = m.install(m.GormModule.InstallFromConfiguration, dep)
	return m.Opened
}

func (m *Module) InstallMock(dep module.Module) bool {
	m.Opened = m.install(m.GormModule.InstallMockFromConfiguration, dep)
	return m.Opened
}

//func (m *Module) RegisterRedis(dep module.Module) bool {
//	return splayer.RegisterRedis(dep)
//}
//
//func (m *Module) InstallRedis(dep module.Module) bool {
//	return splayer.Install(dep)
//}

func (m *Module) Close(dep module.Module) bool {
	if m.Opened {
		if err := m.GormDB.Close(); err != nil {
			m.Logger.Error("close DB error", "error", err)
			return false
		}
	}
	return true
}
