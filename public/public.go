package public

import (
	"embed"
	"io/fs"
	"os"
)

var (
	//go:embed all:**
	rootfs embed.FS
	realfs fs.FS = os.DirFS("./public")
)

func FS(embed bool) fs.FS {
	if embed {
		return rootfs
	} else {
		return realfs
	}
}

func Stat(name string) (fs.FileInfo, error) {
	if info, err := fs.Stat(realfs, name); err == nil {
		return info, nil
	} else {
		return fs.Stat(rootfs, name)
	}
}

func ReadDir(dirname string) ([]fs.DirEntry, error) {
	if dirs, err := fs.ReadDir(realfs, dirname); err == nil {
		return dirs, nil
	} else {
		return fs.ReadDir(rootfs, dirname)
	}
}

func ReadFile(filename string) ([]byte, error) {
	if content, err := fs.ReadFile(realfs, filename); err == nil {
		return content, nil
	} else {
		return fs.ReadFile(rootfs, filename)
	}
}
