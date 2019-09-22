package util

import (
	"os"
	"path/filepath"
)

// 获取可执行文件所在的根目录
func GetRootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}
