package svcx

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/os/gcache"
)

func Cache(group ...string) *gcache.Cache {

	var (
		groupName    = "default"
		cfg_key      = ""
		instance_key = ""
	)
	if len(group) > 0 && len(group[0]) > 0 {
		groupName = group[0]
	}
	cfg_key = fmt.Sprintf("cache.%s", groupName)
	instance_key = fmt.Sprintf("svcx.%s", cfg_key)
	cache, ok := instances.Load(instance_key)
	if !ok {
		var (
			adapter          gcache.Adapter
			vcache           *gcache.Cache
			cfg_sub_type_key string = fmt.Sprintf("%s.type", cfg_key)
		)
		if cache_type := Viper().GetString(cfg_sub_type_key); strings.EqualFold(cache_type, "redis") {
			redis := new_gf_greids(Viper().GetStringMap(cfg_key))
			adapter = gcache.NewAdapterRedis(redis)
		}
		if adapter == nil {
			adapter = gcache.NewAdapterMemory()
		}
		vcache = gcache.NewWithAdapter(adapter)
		instances.Store(instance_key, vcache)
		cache = vcache
	}
	return cache.(*gcache.Cache)
}
