package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

type cmdx struct {
	pfset   *pflag.FlagSet
	init    bool
	ghttpx  bool
	gcronx  bool
	grpcx   bool
	help    bool
	version bool
}

func New(appVersion string) *cmdx {
	var (
		appName = filepath.Base(os.Args[0])
		cx      = &cmdx{}
	)
	cx.pfset = pflag.NewFlagSet(appName, pflag.ContinueOnError)
	cx.pfset.SortFlags = false
	cx.pfset.BoolVar(&cx.init, "init", false, "init in current folder")
	cx.pfset.BoolVar(&cx.ghttpx, "ghttpx", false, "start ghttpx process")
	cx.pfset.BoolVar(&cx.gcronx, "gcronx", false, "start gcronx process")
	cx.pfset.BoolVar(&cx.grpcx, "grpcx", false, "start grpcx process")
	cx.pfset.BoolVarP(&cx.help, "help", "h", false, "show help message")
	cx.pfset.BoolVarP(&cx.version, "version", "v", false, "show version")
	cx.pfset.Usage = func() {
		fmt.Println("Usage:")
		fmt.Printf("  %s [options]\n", appName)
		fmt.Println("Options:")
		cx.pfset.PrintDefaults()
		fmt.Println("Version:")
		fmt.Printf("  %s", appVersion)
		fmt.Println()
	}
	return cx
}

func (cx *cmdx) Run() error {
	var (
		args = os.Args[1:]
	)
	if err := cx.pfset.Parse(args); err != nil {
		return err
	}
	if len(args) <= 0 || cx.help || cx.version {
		cx.pfset.Usage()
		return nil
	}
	if err := cx.init_handler(); err != nil {
		return err
	}
	if err := cx.start_handler(); err != nil {
		return err
	}
	return nil
}
