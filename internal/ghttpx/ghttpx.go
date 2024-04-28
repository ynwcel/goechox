package ghttpx

import (
	"fmt"

	"github.com/ynwcel/goxbase/internal/svcx"
)

func Start() error {
	fmt.Println(svcx.Viper().GetInt("ghttpx.listen"))
	return nil
}
