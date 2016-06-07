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

## DIY
启动服务后并通过浏览器查看，可以发现默认是加载`web/api/cfg/swagger_v2.yaml`这个swagger文档。当然可以选择在这个文档中追加内容，但是更加合理的选择是按照自己的实际需求来安排具体放在哪个文档。那么，如何自定义默认加载哪个swagger文档呢？
1. 找到`web/index.tpl`模板文件。
1. 找到模板文件中的内容`url = "http://${domain}:${port}/api/cfg/swagger_v2.yaml";`
1. 自定义默认加载的swagger文档，将其改为`url = "http://${domain}:${port}/api/swagger.yaml";`。注意：只能更改`http://${domain}:${port}/api/`之后的内容，`api`对应的就是文件系统中`web/api`目录，换句话说，这里提供的是对本地文件系统`web/api`目录的映射功能。

