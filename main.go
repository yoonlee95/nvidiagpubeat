package main

import (
	"os"

	"github.com/yoonlee95/nvidiagpubeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
