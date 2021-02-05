package common

import "github.com/otiai10/copy"

//CopyDir copies folder content
func CopyDir(src, dst string) error {
	options := copy.Options{
		Sync:          true,
		PreserveTimes: true,
	}
	err := copy.Copy(src, dst, options)

	return err
}
