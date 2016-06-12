package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	shfilepath "github.com/voidint/swagger-hub/filepath"
	shiosutil "github.com/voidint/swagger-hub/osutil"
)

const (
	maxPort = 65535
)

var (
	// ErrPort 非法的服务端口号
	ErrPort = errors.New("invalid port")
	// ErrDir 无效的目录路径
	ErrDir = errors.New("invalid directory")
)

// Options 命令行参数
type Options struct {
	Port    uint
	Domain  string
	Dir     string
	LogFile string
}

// Validate 校验命令行参数是否合法
func (opts *Options) Validate() error {
	if opts.Port > maxPort {
		return ErrPort
	}

	if !shiosutil.DirExisted(opts.Dir) {
		return ErrDir
	}
	return nil
}

func main() {
	var opts Options
	flag.UintVar(&opts.Port, "port", 80, "服务端口号")
	flag.StringVar(&opts.Domain, "domain", "apihub.idcos.net", "HTTP服务域名")
	flag.StringVar(&opts.Dir, "dir", "", "需要提供文件服务的目录路径")
	flag.StringVar(&opts.LogFile, "log", "doc-server.log", "日志打印全路径(包含日志文件名称)")
	flag.Parse()

	var err error
	if err = opts.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var logger *log.Logger
	if logger, err = initLog(opts.LogFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if err = Run(opts, logger); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
}

// Run 运行服务
func Run(opts Options, logger *log.Logger) (err error) {
	if err = genIndexHTML(opts, logger); err != nil {
		logger.Println(err)
		return err
	}
	logger.Printf("Start doc service(port=%d, dir=%s, log=%s)\n", opts.Port, opts.Dir, opts.LogFile)

	http.Handle("/", http.FileServer(http.Dir(opts.Dir)))
	if err = http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), nil); err != nil {
		logger.Println(err)
	}
	return err
}

func initLog(file string) (logger *log.Logger, err error) {
	logfile, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return log.New(logfile, "", log.Llongfile|log.LstdFlags), nil
}

// 在指定目录下通过模板生成index.html文件
func genIndexHTML(opts Options, logger *log.Logger) (err error) { // TODO 通过golang的template生成HTML
	indexHTML := filepath.Join(opts.Dir, "index.html")
	indexTPL := filepath.Join(opts.Dir, "index.tpl")

	tplData, err := ioutil.ReadFile(indexTPL)
	if err != nil {
		logger.Println(err)
		return err
	}

	apiBasePath := filepath.Join(opts.Dir, "api")
	paths, err := shfilepath.ScanSwaggerDocs(apiBasePath)
	if err != nil {
		logger.Println(err)
		return err
	}

	logger.Printf("Find docs: %v\n", paths)

	html := string(tplData)
	// html = strings.Replace(html, "${domain}", opts.Domain, -1)
	// html = strings.Replace(html, "${port}", fmt.Sprintf("%d", opts.Port), -1)
	html = strings.Replace(html, "${baseURLs}", genSelectHTML(opts, logger, paths), -1)
	return ioutil.WriteFile(indexHTML, []byte(html), 0666)
}

func genSelectHTML(opts Options, logger *log.Logger, paths []string) string {
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
