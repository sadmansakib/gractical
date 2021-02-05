package common

import (
	"github.com/otiai10/copy"
	"strings"
)

//CopyDir copies folder content
func CopyDir(src, dst string) error {
	options := copy.Options{
		Sync:          true,
		PreserveTimes: true,
	}
	err := copy.Copy(src, dst, options)

	return err
}

//CopyDirExcludingTest copies folder contest except test files with _test substring
func CopyDirExcludingTest(src, dst string) error {
	options := copy.Options{
		Skip: func(src string) (bool, error) {
			return strings.Contains(src, "_test"), nil
		},
		Sync:          true,
		PreserveTimes: true,
	}
	err := copy.Copy(src, dst, options)

	return err
}
