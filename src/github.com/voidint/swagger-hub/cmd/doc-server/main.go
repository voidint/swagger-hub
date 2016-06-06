package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	shioutil "github.com/voidint/swagger-hub/ioutil"
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
func genIndexHTML(opts Options, logger *log.Logger) (err error) {
	indexHTML := filepath.Join(opts.Dir, "index.html")
	indexTPL := filepath.Join(opts.Dir, "index.tpl")

	// _ = os.Remove(indexHTML)

	if _, err = shioutil.CopyFile(indexTPL, indexHTML); err != nil {
		logger.Println(err)
		return err
	}

	// f, err := os.OpenFile(indexHTML, os.O_RDWR, 0)
	// if err != nil {
	// 	logger.Println(err)
	// 	return
	// }
	// defer f.Close()

	return shioutil.ReplaceFileContent(indexHTML, map[string]string{
		"${domain}": opts.Domain,
		"${port}":   fmt.Sprintf("%d", opts.Port),
	})
}
