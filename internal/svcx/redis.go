package svcx

import (
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func Redis(group ...string) *gredis.Redis {
	var (
		groupName    = gredis.DefaultGroupName
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("redis.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	redis, ok := instances.Load(instance_key)
	if !ok {
		cfg_map := Viper().GetStringMap(cfg_key)
		vredis := new_gf_greids(cfg_map)
		instances.Store(instance_key, vredis)
		redis = vredis
	}
	return redis.(*gredis.Redis)
}

func new_gf_greids(cfg map[string]any) *gredis.Redis {
	var (
		redis    *gredis.Redis
		cfg_node = new(gredis.Config)
		err      error
	)
	if err = gconv.Scan(cfg, &cfg_node); err != nil {
		panic(gerror.Wrap(err, "parse-redis-config-failed"))
	}
	if redis, err = gredis.New(cfg_node); err != nil {
		panic(gerror.Wrapf(err, "new-redis-instance-failed"))
	}
	return redis
}
