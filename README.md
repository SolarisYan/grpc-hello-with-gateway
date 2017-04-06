# grpc-hello-with-gateway
example for grpc-hello-with-grpc-gateway use gb on win 10

# 1. 生成 gRPC golang stub 类
python -m grpc.tools.protoc
       -I..\..\protos  -I.
       -I%GOPATH%\src
       -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis
       --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:.
       ../../protos/helloworld.proto

It will generate a stub file path/to/your_service.pb.go

# 2. 生成反向代理代码
python -m grpc.tools.protoc
       -I..\..\protos  -I.
       -I%GOPATH%\src
       -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis
       --grpc-gateway_out=logtostderr=true:.
       ../../protos/helloworld.proto

It will generate a reverse proxy path/to/your_service.pb.gw.go

# 3. 编写 entrypoint 文件
src/helloworld_restful_swagger/main.go

# 4. 配置 go 代理
先设置代理:(前提是安装配置好 ss)
<br />
set http_proxy=http://127.0.0.1:1080/pac?t=20170405092729716
<br />
set https_proxy=http://127.0.0.1:1080/pac?t=20170405092729716

# 5. 获取 main.go 文件中使用到的依赖包
gb vendor fetch github.com/grpc-ecosystem/grpc-gateway/runtime
<br />
gb vendor fetch google.golang.org/genproto/googleapis/api/annotations

# 6. 生成可执行文件
gb build helloworld_restful_swagger

# 7. python.exe .\helloworld\greeter_server.py

# 8. .\bin\helloworld_restful_swagger.exe

# 9. 测试

curl -sSk http://localhost:8080/v1/example/echo -d '{"name": " solaris"}'
<br />
{"message":"Hello,  solaris!"}

