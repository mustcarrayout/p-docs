

- ## build
	- 交叉编译
		- ```shell
		  # 是否开启cgo
		  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
		  ```
		- 参数
		  
		  GOOS=linux|windows|darwin(mac)|android|freebsd
		  GOARCH=amd64|arm|386
# 参考
- [go cmd](https://blog.csdn.net/wohu1104/article/details/106295007