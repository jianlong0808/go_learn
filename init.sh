#配置环境变量
## 配置 GOPROXY 环境变量, 否则下载三方包的时候会有问题
#export GOPROXY=https://goproxy.io,direct
## 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
#export GOPRIVATE=git.mycompany.com,github.com/my/private

cd go_learn
#生成go.mod文件
go mod init
# go mod vendor将依赖包复制到项目下的 vendor 目录。建议一些使用了被墙包的话可以这么处理，方便用户快速使用命令go build -mod=vendor编译。
# 但是官方不建议这么搞
go mod vendor
