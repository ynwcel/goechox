package cmd

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ynwcel/goxbase/internal/gcronx"
	"github.com/ynwcel/goxbase/internal/ghttpx"
	"github.com/ynwcel/goxbase/internal/grpcx"
	"github.com/ynwcel/goxbase/internal/svcx"
	"golang.org/x/sync/errgroup"
)

func (cx *cmdx) start_handler() error {
	var (
		dirs = []string{
			"./runtimes",
			"./storages",
		}
		errGroup = new(errgroup.Group)
	)
	for _, d := range dirs {
		if err := gfile.Mkdir(d); err != nil {
			return err
		}
	}
	if err := svcx.Init(); err != nil {
		return err
	}
	if cx.ghttpx {
		errGroup.Go(func() error {
			return ghttpx.Start()
		})
	}
	if cx.gcronx {
		errGroup.Go(func() error {
			return gcronx.Start()
		})
	}
	if cx.grpcx {
		errGroup.Go(func() error {
			return grpcx.Start()
		})
	}
	return errGroup.Wait()
}
