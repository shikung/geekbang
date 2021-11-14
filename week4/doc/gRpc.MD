1. 安装 protobuf
下载 protoc-3.5.1-win32.zip

把解压后的 bin 目录下的 protoc.exe 放入到 GOBIN (go env 查看) 中

2. 安装 grpc
Git clone https://github.com/grpc/grpc-go

将grpc-go文件夹重命名为grpc，放入到google.golang.org中，完整路径如下

C:\GOPATH\src\google.golang.org\grpc

3. 安装 genproto
Git clone https://github.com/google/go-genproto

将go-genproto文件夹重命名为genproto，放到google.golang.org下，完整路径如下

C:\GOPATH\src\google.golang.org\genproto

4. 安装 proto
go get -u github.com/golang/protobuf/proto

5. 安装 protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go

go install github.com/golang/protobuf/protoc-gen-go/

go run google.golang.org/grpc/examples/helloworld/greeter_server/main.go

6. 下载依赖包
创建目录 C:\GOPATH\src\golang.org\x
进入 x 目录
git clone https://github.com/golang/net.git
git clone https://github.com/golang/text.git

7. 使用
cd 到 src\google.golang.org\grpc\examples\helloworld目录
执行命令生成代码 helloworld.pb.go
protoc -I ./helloworld --go_out=plugins=grpc:./helloworld ./helloworld\helloworld.proto

protoc -I ./helloworld --go_out=plugins=grpc:./helloworld ./helloworld\helloworld.proto

protoc -I ./api --go-out=plugins=grpc:./api ./api/user.proto

protoc --go_out=plugins=grpc:. *.proto
protoc --go_out=plugins=grpc: *.proto


protoc -I --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/user.proto
protoc --go_out=. --go_opt=plugins=grpc: --go-grpc_out=. --go-grpc_opt=plugins=grpc: user.proto


protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_rative api/user.proto
protoc --go_out=. --go-grpc_opt=paths=source_rative api/user.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/user.proto
protoc --go_out=. --go_opt=paths=source_relative  api/user.proto


作者：谷前雨
链接：https://www.jianshu.com/p/be8fc6ab43d5
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

进入 \src\google.golang.org\grpc\examples\helloworld\greeter_server 目录
执行 go run main.go
进入 \src\google.golang.org\grpc\examples\helloworld\greeter_client 目录
执行 go run main.go yourname

作者：butters001
链接：https://www.jianshu.com/p/9d123f2a1043
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。