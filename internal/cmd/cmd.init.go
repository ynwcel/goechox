package cmd

import (
	"io/fs"
	"path/filepath"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ynwcel/goxbase/public"
)

func (cx *cmdx) init_handler() error {
	if !cx.init {
		return nil
	}
	var (
		rootfs     = public.FS(true)
		files      = get_public_files(rootfs)
		target_dir = "./public"
	)
	if err := gfile.Mkdir(target_dir); err != nil {
		return err
	}
	for _, f := range files {
		var (
			target   = filepath.Join(target_dir, f)
			finfo, _ = fs.Stat(rootfs, f)
		)
		if finfo.IsDir() {
			if err := gfile.Mkdir(target); err != nil {
				return err
			}
		} else if content, err := fs.ReadFile(rootfs, f); err != nil {
			return err
		} else if err := gfile.PutBytes(target, content); err != nil {
			return err
		}
	}

	return nil
}
