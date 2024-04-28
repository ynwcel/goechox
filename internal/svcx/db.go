package svcx

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func DB(group ...string) gdb.DB {
	var (
		groupName    = gdb.DefaultGroupName
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("database.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	db, ok := instances.Load(instance_key)
	if !ok {
		cfg_any := Viper().Get(cfg_key)
		vdb := new_gf_gdb(cfg_any)
		if log_cfg_map := Viper().GetStringMap(fmt.Sprintf("%s.logger", cfg_key)); len(log_cfg_map) > 0 {
			vdb.SetLogger(new_gf_glog(log_cfg_map))
		} else if group := Viper().GetString(fmt.Sprintf("%s.logger_group", cfg_key)); len(group) > 0 {
			vdb.SetLogger(Log(group))
		} else if run_mode := os.Getenv("RUN_MODE"); len(run_mode) > 0 && strings.ToLower(run_mode) == "debug" {
			vdb.SetDebug(true)
			vdb.SetLogger(Log())
		}
		instances.Store(instance_key, vdb)
		db = vdb
	}
	return db.(gdb.DB)
}

func new_gf_gdb(cfg any) gdb.DB {

	switch cfg.(type) {
	case []any:
		return new_gdb_by_slices(cfg.([]any))
	case map[string]any:
		return new_gdb_by_map(cfg.(map[string]any))
	default:
		panic("invalid-param-of-new-gdb")
	}
}

func new_gdb_by_slices(cfg []any) gdb.DB {
	var (
		groupName = fmt.Sprintf("pgdb_%d", time.Now().Unix())
		group     = gdb.ConfigGroup{}
		db        gdb.DB
		err       error
	)
	if err = gconv.Scan(cfg, &group); err != nil {
		panic(gerror.Wrap(err, "scan-gdb-config-node-failed"))
	}

	gdb.SetConfigGroup(groupName, group)
	if db, err = gdb.NewByGroup(groupName); err != nil {
		panic(gerror.Wrap(err, "new-gdb-failed"))
	}
	return db
}

func new_gdb_by_map(cfg map[string]any) gdb.DB {
	var (
		dbnode_cfg = gdb.ConfigNode{}
		db         gdb.DB
		err        error
	)
	if err = gconv.Scan(cfg, &dbnode_cfg); err != nil {
		panic(gerror.Wrap(err, "scan-gdb-config-node-failed"))
	}
	if db, err = gdb.New(dbnode_cfg); err != nil {
		panic(gerror.Wrap(err, "new-gdb-failed"))
	}
	return db
}
