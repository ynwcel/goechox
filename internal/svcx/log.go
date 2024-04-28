package svcx

import (
	"fmt"

	"github.com/gogf/gf/v2/os/glog"
)

func Log(group ...string) glog.ILogger {
	var (
		groupName    = glog.DefaultName
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("logger.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	log, ok := instances.Load(instance_key)
	if !ok {
		cfg_map := Viper().GetStringMap(cfg_key)
		vlog := new_gf_glog(cfg_map)
		instances.Store(instance_key, vlog)
		log = vlog
	}
	return log.(glog.ILogger)
}

func new_gf_glog(cfg_map map[string]any) glog.ILogger {
	log := glog.New()
	log.SetConfig(gfx_glog_default_cfg())
	if len(cfg_map) > 0 {
		if err := log.SetConfigWithMap(cfg_map); err != nil {
			panic(err)
		}
	}
	log = log.Line(true)
	log.SetHandlers(glog.HandlerJson)
	return log
}

func gfx_glog_default_cfg() glog.Config {
	def := glog.DefaultConfig()
	def.Path = "./runtimes"
	def.File = "{Y-m-d}.log"
	def.TimeFormat = "2006-01-02T15:04:05.999"
	def.StdoutPrint = false
	def.Level = glog.LEVEL_ALL
	def.Flags = glog.F_FILE_SHORT | glog.F_CALLER_FN | glog.F_TIME_DATE | glog.F_TIME_MILLI
	return def
}
