GRPC-Gateway 模板demo

流程：
````
1、安装buf：https://github.com/bufbuild/buf/releases
2、创建buf配置文件：buf.gen.yaml, buf.work.yaml, proto/buf.yml
3、创建proto文件：proto/user/user.proto
4、安装buf依赖：cd proto && buf update
5、构建pb文件: buf generate
6、编写两个server代码
````

Demo测试：
````
启动服务：
go run .\server\grpc\main.go
go run .\server\http\main.go

测试：
8080端口：GRPC服务
8081端口：HTTP接口(请求参数只支持json格式)
````