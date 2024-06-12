package lite_backup

import (
	"io"
	"os"
)

func CreateLocalCopy(src, dst string) (string, error) {
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return "Could not open source file", err
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			return
		}
	}(sourceFile)

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return "Could not create destination file", err
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			return
		}
	}(destFile)

	// Copy the contents from the source to the destination
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return "Could not copy file", err
	}

	// Flush the file system buffers to ensure all data is written
	err = destFile.Sync()
	if err != nil {
		return "Could not finish copy", err
	}

	return "Copy successfully created", nil
}
