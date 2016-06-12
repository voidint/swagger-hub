#swagger-hub
为[swagger-ui](https://github.com/swagger-api/swagger-ui)提供动态渲染本地swagger文档服务。

# 安装
```
// 使用golang的工具链下载源代码并编译安装
go get -u github.com/voidint/swagger-hub/...

// 给二进制文件建立软连接
ln -s $GOPATH/bin/doc-server /usr/local/bin/doc-server
```

# 使用方法
```
$ doc-server -h
Usage of doc-server:
  -dir string
    	需要提供文件服务的目录路径
  -domain string
    	HTTP服务域名 (default "apihub.idcos.net")
  -log string
    	日志打印全路径(包含日志文件名称) (default "doc-server.log")
  -port uint
    	服务端口号 (default 80)
```

## 参数说明
- `dir`: 需要提供文件服务的目录路径。具体指的就是源代码中名为`web`的那个目录路径。
- `domain`: HTTP服务域名。如果是本地运行，那么可以指定为`localhost`，如果是部署在服务器上以供他人访问，那么可以指定那台服务器的`IP`或者`域名`。
- `log`: 日志打印全路径(包含日志文件名称)。
- `port`: 服务端口号。

## 示例
1. 编写符合[swagger specification](http://swagger.io/specification/)规范的文档。
1. 将文档（`YAML`格式或者`JSON`格式）放入`dir`参数所指向的名为`web`目录下的`api`目录中。
1. 启动服务`doc-server --dir $GOPATH/src/github.com/voidint/swagger-hub/web --domain localhost --log /tmp/doc-server.log --port 8090`。
1. 通过浏览器访问`http://localhost:8090`。
