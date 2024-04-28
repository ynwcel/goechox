package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/ynwcel/goxbase/internal/cmd"
)

var (
	appVersion = "0.0.1"
)

func main() {
	if err := cmd.New(appVersion).Run(); err != nil {
		panic(err)
	}
}
