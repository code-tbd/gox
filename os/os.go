package os

import (
	"os"
	"path/filepath"
)

// Create x 传入完整绝对路径，不用担心路径不存在
func Create(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
