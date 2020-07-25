package plugin

import (
	"context"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Logger = external.Logger
type ConfigLoader = external.ConfigLoader
type ServerConfig = config.ServerConfig
type Module = module.Module

type Plugin interface {
	Configuration(logger Logger, loader ConfigLoader, cfg *ServerConfig) (plg Plugin)
	Inject(module Module) (plg Plugin)
	Work(ctx context.Context)
}
