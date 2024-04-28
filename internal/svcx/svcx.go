package svcx

import (
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/joho/godotenv"
)

var (
	instances = new(sync.Map)
)

func Init() (err error) {
	if _, ok := instances.Load("svcx.inited"); !ok {
		defer func() {
			if exception := recover(); exception != nil {
				if exception_err, ok := exception.(error); ok {
					err = exception_err
				} else {
					err = fmt.Errorf("%v", exception)
				}
			}
		}()
		if gfile.Exists(".env") {
			if err = godotenv.Load(); err != nil {
				return
			}
		}
		_ = Viper()
	}
	return
}
