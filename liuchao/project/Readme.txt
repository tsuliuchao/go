此时运行项目，不能像之前简单的使用go run main.go，因为包main包含main.go和router.go的文件，因此需要运行go run *.go命令编译运行。
如果是最终编译二进制项目，则运行go build -o app
