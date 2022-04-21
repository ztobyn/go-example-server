# go-demo

Go语言项目流水线demo

## 项目初始化
*go.mod*中模块名为demo模块名，首次初始化项目，先删除go.mod文件，然后执行
 
> go mod init 模块名

进行项目初始化

## 构建命令
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-s -w -X go-demo/version.Commit=$(git rev-parse --short HEAD) -X go-demo/version.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S')" -o go-demo

## 运行命令
export SERVER_PORT=8080; ./go-demo
