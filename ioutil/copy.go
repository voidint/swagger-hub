package ioutil

import (
	"io"
	"os"
)

// CopyFile 将原文件逐字节拷贝至目标文件
func CopyFile(srcFile, destFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dest, err := os.Create(destFile)
	if err != nil {
		return 0, err
	}
	defer dest.Close()
	return io.Copy(dest, src)
}
