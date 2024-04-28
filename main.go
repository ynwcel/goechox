package main

import "github.com/ynwcel/goxbase/internal/cmd"

var (
	appVersion = "0.0.1"
)

func main() {
	if err := cmd.New(appVersion).Run(); err != nil {
		panic(err)
	}
}
