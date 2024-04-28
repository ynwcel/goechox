package ghttpx

import (
	"github.com/labstack/echo/v4"
	"github.com/ynwcel/goxbase/internal/svcx"
	"github.com/ynwcel/goxbase/pkg/pechox"
)

func Start() error {
	var (
		server = pechox.New(svcx.Viper().GetStringMap("ghttpx"))
		engine = server.Engine()
	)
	engine.GET("/", func(ectx echo.Context) error {
		return ectx.String(200, "running~")
	})
	return server.Run()
}
