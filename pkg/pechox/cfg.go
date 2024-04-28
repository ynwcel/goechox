package pechox

import (
	"time"
)

type cfg struct {
	Debug          bool          `json:"debug"`
	Listen         int           `json:"listen"`
	ReadTimeout    time.Duration `json:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"`
	IdleTimeout    time.Duration `json:"idle_timeout"`
	MaxHeaderBytes int           `json:"max_header_bytes"`
}

func default_cfg() *cfg {
	return &cfg{
		Debug:          false,
		Listen:         8000,
		ReadTimeout:    time.Second * 3,
		WriteTimeout:   time.Second * 3,
		IdleTimeout:    time.Second * 3,
		MaxHeaderBytes: 1 << 20,
	}
}
