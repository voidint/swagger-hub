package util

import "strings"

// ReplaceStr 将字符串line中按照pairs的key替换成对应的value并返回。
func ReplaceStr(line string, pairs map[string]string) string {
	if len(line) <= 0 {
		return line
	}
	for k, v := range pairs {
		line = strings.Replace(line, k, v, -1)
	}
	return line
}
