package util

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// DirExisted 返回指定路径的目录是否存在的布尔值
func DirExisted(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.IsDir()
}

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

// ReplaceFileContent 替换文件中的内容
func ReplaceFileContent(dstFile string, pairs map[string]string) (err error) {
	f, err := os.OpenFile(dstFile, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	var fi os.FileInfo
	if fi, err = f.Stat(); err != nil {
		return err
	}

	var buf bytes.Buffer
	rd := bufio.NewReader(f)

	for {
		var line string
		if line, err = rd.ReadString('\n'); err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if _, err = buf.WriteString(ReplaceStr(line, pairs)); err != nil {
			return nil
		}
	}
	// if err = f.Truncate(0); err != nil {
	// 	return err
	// }
	if err = os.Truncate(fi.Name(), 0); err != nil {
		return nil
	}

	_, err = buf.WriteTo(f)
	return err
}
