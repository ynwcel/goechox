package svcx

import (
	"fmt"

	"github.com/gogf/gf/v2/os/glog"
)

func Log(group ...string) glog.ILogger {
	var (
		groupName    = glog.DefaultName
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	instance_key = fmt.Sprintf("svcx.logger.%s", groupName)
	log, ok := instances.Load(instance_key)
	if !ok {
		cfg_map := Viper().GetStringMap(instance_key)
		vlog := new_gf_glog(groupName, cfg_map)
		instances.Store(instance_key, vlog)
		log = vlog
	}
	return log.(glog.ILogger)
}

func new_gf_glog(group string, cfg_map map[string]any) glog.ILogger {
	log := glog.New()
	log.SetConfig(gfx_glog_default_cfg(group))
	if len(cfg_map) > 0 {
		if err := log.SetConfigWithMap(cfg_map); err != nil {
			panic(err)
		}
	}
	log = log.Line(true)
	log.SetHandlers(glog.HandlerJson)
	return log
}

func gfx_glog_default_cfg(group string) glog.Config {
	def := glog.DefaultConfig()
	def.Path = "./runtime"
	def.File = fmt.Sprintf("%s-{Y-m-d}.log", group)
	def.TimeFormat = "2006-01-02T15:04:05.999"
	def.StdoutPrint = false
	def.Level = glog.LEVEL_ALL
	def.Flags = glog.F_FILE_SHORT | glog.F_CALLER_FN | glog.F_TIME_DATE | glog.F_TIME_MILLI
	return def
}
