package ioutil

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/voidint/swagger-hub/strutil"
)

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

		if _, err = buf.WriteString(strutil.ReplaceStr(line, pairs)); err != nil {
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
