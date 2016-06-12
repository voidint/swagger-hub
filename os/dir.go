package os

import "os"

// DirExisted 返回指定路径的目录是否存在的布尔值
func DirExisted(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.IsDir()
}
