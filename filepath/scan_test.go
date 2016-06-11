package filepath

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScanSwaggerDocs(t *testing.T) {
	var dir string

	dir = "../web/api"
	Convey(fmt.Sprintf("扫描路径（%s）下的JSON和YAML文件路径", dir), t, func() {
		paths, err := ScanSwaggerDocs(dir)
		So(err, ShouldBeNil)
		t.Logf("%s-->%v\n", dir, paths)
		So(paths, ShouldNotBeEmpty)
	})

	dir = "../web/lib"
	Convey(fmt.Sprintf("扫描路径（%s）下的JSON和YAML文件路径", dir), t, func() {
		paths, err := ScanSwaggerDocs(dir)
		So(err, ShouldBeNil)
		So(paths, ShouldBeEmpty)
	})

	dir = "../web/lib2"
	Convey(fmt.Sprintf("扫描无效路径（%s）下的JSON和YAML文件路径", dir), t, func() {
		paths, err := ScanSwaggerDocs(dir)
		So(err, ShouldNotBeNil)
		So(paths, ShouldBeEmpty)
	})

}
