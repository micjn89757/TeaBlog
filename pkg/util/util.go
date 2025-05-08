package util

import (
	"path/filepath"
	"runtime"
)

// 获取当前函数所在路径和行号
func GetCallerInfo() (string, int) {
	_, filename, line, _ := runtime.Caller(1) // 0表示调用信息为runtime.Caller， 1表示调用信息为更上一层
	return filename, line
}


func GetRootPath() string {
	path, _ := GetCallerInfo()
	return filepath.Join(filepath.Dir(path), "/../../")
}