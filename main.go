package main

import (
	"github.com/gahara/lite_backup/local"
	"os"
)

func main() {
	args := os.Args
	src := args[1]
	dst := args[2]
	message, err := local.CreateLocalCopy(src, dst)
	println(message, err)
}
