package oss

import (
	"github.com/Myriad-Dreamin/boj-v6/external"
	mcore "github.com/Myriad-Dreamin/boj-v6/lib/core"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Module struct {
	OSSEngine *external.OSSEngine
	mcore.LoggerModule

	Opened bool
}

func NewModule() Module {
	return Module{
		Opened: false,
	}
}

func (m *Module) install(dep module.Module) bool {
	//path string, opts *opt.Options

	path := "."

	db, err := NewLevelDB(path, nil)
	if err != nil {
		panic(err)
	}

	err = dep.ProvideImpl(new(*external.OSSEngine), db)
	if err != nil {
		panic(err)
	}

	m.OSSEngine = db
	m.Opened = true

	return m.LoggerModule.Install(dep)
}

// todo: add close handler
func (m *Module) Close(_ module.Module) bool {
	if m.Opened {
		if err := m.OSSEngine.Close(); err != nil {
			m.Logger.Error("close oss engine error", "error", err)
			return false
		}
	}
	return true
}
