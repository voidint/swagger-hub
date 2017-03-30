package cli

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli"
	"github.com/voidint/swagger-hub/util"
)

const (
	maxPort = 65535
)

var (
	// ErrPort 非法的服务端口号
	ErrPort = errors.New("invalid port")
	// ErrDir 无效的目录路径
	ErrDir = errors.New("invalid UI project directory")
)

// RunOptions Run子命令参数
type RunOptions struct {
	Domain string
	Port   uint
	Dir    string
}

// Validate 校验命令行参数是否合法
func (opts *RunOptions) Validate() error {
	if opts.Port > maxPort {
		return ErrPort
	}

	if !util.DirExisted(opts.Dir) {
		return ErrDir
	}
	return nil
}

func run(ctx *cli.Context) (err error) {
	opts := RunOptions{
		Domain: ctx.String("domain"),
		Port:   ctx.Uint("port"),
		Dir:    ctx.String("dir"),
	}

	if err = opts.Validate(); err != nil {
		return cli.NewExitError(err, 1)
	}

	// 扫描文档目录及其子目录中所有swagger文档并生成index.html内容
	if err = genIndexHTML(&opts, logger); err != nil {
		return cli.NewExitError(err, 1)
	}

	done := make(chan struct{})

	// 监视API文档目录，若发生变动，则立即更新index.html
	apiBasePath := filepath.Join(opts.Dir, "api")
	go util.Watch(logger, done, apiBasePath, func(event fsnotify.Event) {
		var opDesc string
		switch event.Op {
		case fsnotify.Chmod:
			opDesc = "Chmod"
		case fsnotify.Create:
			opDesc = "Create"
		case fsnotify.Remove:
			opDesc = "Remove"
		case fsnotify.Rename:
			opDesc = "Rename"
		case fsnotify.Write:
			opDesc = "Write"
		}
		logger.Printf("%s --> %s\n", event.Name, opDesc)
		if event.Op == fsnotify.Create ||
			event.Op == fsnotify.Remove ||
			event.Op == fsnotify.Rename {
			genIndexHTML(&opts, logger)
		}
	})

	defer func() {
		logger.Println("write data to done channel")
		done <- struct{}{}
	}()

	logger.Printf("Serving on http://%s:%d\n", opts.Domain, opts.Port)

	http.Handle("/", http.FileServer(http.Dir(opts.Dir)))
	if err = http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), nil); err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

// 在指定目录下通过模板生成index.html文件
func genIndexHTML(opts *RunOptions, logger *log.Logger) (err error) { // TODO 通过golang的template生成HTML
	indexHTML := filepath.Join(opts.Dir, "index.html")
	indexTPL := filepath.Join(opts.Dir, "index.tpl")

	tplData, err := ioutil.ReadFile(indexTPL)
	if err != nil {
		logger.Println(err)
		return err
	}

	apiBasePath := filepath.Join(opts.Dir, "api")
	paths, err := util.ScanSwaggerDocs(apiBasePath)
	if err != nil {
		logger.Println(err)
		return err
	}

	logger.Printf("Find docs: %v\n", paths)

	html := string(tplData)
	html = strings.Replace(html, "${baseURLs}", genSelectHTML(opts, logger, paths), -1)
	return ioutil.WriteFile(indexHTML, []byte(html), 0666)
}

func genSelectHTML(opts *RunOptions, logger *log.Logger, paths []string) string {
	apiBasePath := filepath.Join(opts.Dir, "api")
	baseURI := fmt.Sprintf("http://%s:%d/api", opts.Domain, opts.Port)

	var buf bytes.Buffer
	buf.WriteString(`<select id="input_baseUrl" name="baseUrl">`)
	for _, path := range paths {
		val := strings.Replace(path, apiBasePath, baseURI, -1)
		buf.WriteString(fmt.Sprintf(`<option value="%s">%s</option>`, val, val))
	}
	buf.WriteString(`</select>`)
	return buf.String()
}
