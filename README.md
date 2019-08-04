# todos-react-golang

## go 项目

### 初始化

```shell
$ mkdir app
$ cd app/
$ GO111MODULE=on
$ go mod init github.com/yuchou87/todos-react-golang
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/jinzhu/gorm
```

### 安装依赖

```shell
$ go get ./...      # 查找并下载依赖
$ go mod vendor     # 复制第三方依赖到vendor目录
```

### 打包部署

```shell
# 只有一个main.go的时候，执行下面命令
$ go run main.go
$ go build main.go

# 有多个go文件的时候，执行下面命令
$ go run .
$ go build

# 打包加上时间和hash值
$ go build -ldflags "-X main.buildstamp=`date +%Y-%m-%d_%H:%M:%S` -X main.githash=`git rev-parse HEAD`" .
```

## 接口操作

```shell
$ curl -XGET http://localhost:8080/api/v1.0/todos/      # 获取所有的todos
$ curl -XPOST http://localhost:8080/api/v1.0/todos/     # 创建一个todo
$ curl -XPUT http://localhost:8080/api/v1.0/todos/1     # 更新一个todo
$ curl -XDELETE http://localhost:8080/api/v1.0/todos/1  # 删除一个todo
```

## 跨域

### 后端配置

```go
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
```

### 前端配置

```js
useEffect(() => {
  fetch("http://10.0.0.147:8080/api/v1.0/todos/", {
    method: "GET",
    mode: "cors",
    credentials: "include",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
    }
  })
    .then(response => response.json())
    .then(data => {
      setTodos(data.data);
    });
}, []);
```
