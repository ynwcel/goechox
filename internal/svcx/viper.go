package svcx

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/spf13/viper"
	"github.com/ynwcel/goxbase/public"
)

func Viper() *viper.Viper {
	instance_key := "svcx.viper"
	cfg, ok := instances.Load(instance_key)
	if !ok {
		var (
			cfg_file    = "./config.yaml"
			cfg_content []byte
			err         error
		)
		if env_cfg_file := os.Getenv("CFG_FILE"); len(env_cfg_file) > 0 {
			cfg_file = fmt.Sprintf("./%s.yaml", strings.TrimSuffix(env_cfg_file, ".yaml"))
		}
		if gfile.Exists(cfg_file) {
			cfg_content, err = os.ReadFile(cfg_file)
		} else {
			cfg_content, err = public.ReadFile(filepath.Base(cfg_file))
		}
		if err != nil {
			panic(fmt.Errorf("load-cfg-file-faild:%w", err))
		}
		v := viper.New()
		v.SetConfigType("yaml")
		if err := v.ReadConfig(bytes.NewReader(cfg_content)); err != nil {
			panic(fmt.Errorf("parse-cfg-file-content-faild:%w", err))
		}
		instances.Store(instance_key, v)
		cfg = v
	}
	return cfg.(*viper.Viper)
}
