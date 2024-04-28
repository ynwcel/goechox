package svcx

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func DB(group ...string) gdb.DB {
	var (
		groupName    = gdb.DefaultGroupName
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	instance_key = fmt.Sprintf("svcx.database.%s", groupName)
	db, ok := instances.Load(instance_key)
	if !ok {
		cfg_any := Viper().Get(instance_key)
		vdb := new_gf_gdb(cfg_any)
		vdb.SetLogger(Log(fmt.Sprintf("db.%s", groupName)))
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
	for idx, val := range cfg {
		if val_map, ok := val.(map[string]any); ok {
			if _, ok := val_map["debug"]; !ok {
				val_map["debug"] = true
			}
			cfg[idx] = val_map
		}
	}

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
	if _, ok := cfg["debug"]; !ok {
		cfg["debug"] = true
	}

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
