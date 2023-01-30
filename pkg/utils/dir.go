package utils

import (
	"github.com/pkg/errors"
	"os"
)

func PathExits(path string) (bool, error) {
	file, err := os.Stat(path)
	if err == nil {
		if file.IsDir() {
			return true, nil
		}
		return false, errors.New("文件已存在")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
