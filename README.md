# todos-react-golang

## go 项目初始化

```shell
$ mkdir app
$ cd app/
$ GO111MODULE=on
$ go mod init github.com/yuchou87/todos-react-golang
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/jinzhu/gorm
```

## 接口操作

```shell
$ curl -XGET http://localhost:8080/api/v1.0/todos/      # 获取所有的todos
$ curl -XPOST http://localhost:8080/api/v1.0/todos/     # 创建一个todo
$ curl -XPUT http://localhost:8080/api/v1.0/todos/1     # 更新一个todo
$ curl -XDELETE http://localhost:8080/api/v1.0/todos/1  # 删除一个todo
```
