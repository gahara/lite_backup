package lite_backup

import "os"

func testRun() {
	args := os.Args
	src := args[1]
	dst := args[2]
	message, err := CreateLocalCopy(src, dst)
	println(message, err)
}
