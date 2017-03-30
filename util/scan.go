package util

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	scanPath = "/api/"
)

const (
	// JSONExt JSON格式文件扩展名
	JSONExt = ".json"
	// YAMLExt YAML格式文件扩展名
	YAMLExt = ".yaml"
)

// ScanSwaggerDocs 扫描指定目录及其子目录下的swagger文档。
func ScanSwaggerDocs(rootDir string) (paths []string, err error) {
	// ch = make(chan string)
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, ierr error) error {
		if ierr != nil {
			return ierr
		}
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(info.Name()))
		if ext != JSONExt && ext != YAMLExt {
			return nil
		}
		// ch <- path
		paths = append(paths, path)
		return nil
	})
	// return ch
	return paths, err
}
