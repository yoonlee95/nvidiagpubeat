// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration `config:"period"`
	Query  string        `config:"query"`
	Env    string        `config:"env"`
}

var DefaultConfig = Config{
	Period: 1 * time.Second,
	Query:  "utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate",
	Env:    "production",
}
