package ghttpx

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/ynwcel/goxbase/internal/svcx"
)

func Start() error {
	ctx := gctx.GetInitCtx()
	svcx.Log().Info(ctx, "start-ghttpx")
	fmt.Println(svcx.Viper().GetInt("ghttpx.listen"))
	fmt.Println(svcx.DB().Query(ctx, "show tables"))
	fmt.Println(svcx.Redis().Keys(ctx, "*"))
	//svcx.Debug()
	return nil
}
