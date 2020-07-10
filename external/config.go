package external

type ConfigLoader func(cfg interface{}, cfgPath string) bool
