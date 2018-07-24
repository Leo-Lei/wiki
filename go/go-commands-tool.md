---
layout: post
title: Go Command tool
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# 通过os库获取命令行参数
```go
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    fmt.Println(os.Args)
}
```

编译后执行
```bash
./cmd -user="tom"
[./cmd -user=tom]
```

# 通过flag库获取命令行参数
新建文件`hello.go`:
```go
func main()  {

	ok := flag.Bool("ok", false, "is ok")
	id := flag.Int("id", 0, "id")
	port := flag.String("port", ":8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "123", "name")

	flag.Parse()

	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)
}
```
编译hello.go:
```go
go build hello.go
```
运行hello:
```go
./hello --id=3 --name="golang"
```
使用flag操作命令行参数，支持格式如下:
```go
-id=1
--id=1
-id 1
--id 1
```
