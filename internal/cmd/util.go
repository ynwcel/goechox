package cmd

import (
	"fmt"
	"io/fs"
	"strings"
)

func get_public_files(rootfs fs.FS) []string {
	var (
		listfn func(path string) []string
	)
	listfn = func(path string) []string {
		var (
			result = make([]string, 0, 30)
			rule   = fmt.Sprintf("%s/*", strings.TrimRight(path, "/"))
		)
		if files, err := fs.Glob(rootfs, rule); err == nil {
			for _, f := range files {
				if strings.Contains(f, ".go") {
					continue
				}
				result = append(result, f)
				if finfo, err := fs.Stat(rootfs, f); err == nil && finfo.IsDir() {
					subresult := listfn(f)
					result = append(result, subresult...)
				}
			}
		}
		return result
	}
	return listfn(".")
}

func try(f func() error) (err error) {
	defer func() {
		if exception := recover(); exception != nil {
			if exception_err, ok := exception.(error); ok {
				err = exception_err
			} else {
				err = fmt.Errorf("%v", exception)
			}
		}
	}()
	return f()
}
