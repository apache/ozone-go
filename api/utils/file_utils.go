package utils

import (
	"os"
)

// Exists TODO
// 判断所给路径文件/文件夹是否存在
func Exists(p string) bool {
	_, err := os.Stat(p) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir TODO
// 判断所给路径是否为文件夹
func IsDir(p string) bool {
	s, err := os.Stat(p)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile TODO
// 判断所给路径是否为文件
func IsFile(p string) bool {
	s, err := os.Stat(p)
	if err != nil {
		return false
	}
	return !s.IsDir()
}
