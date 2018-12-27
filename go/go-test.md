---
layout: post
title: Go Slice
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# Go单元测试

* 文件名格式：`hello_test.go`
* 测试方法的签名:
```go
func Test_hello_world(t *testing.T)  {
    // .....
}
```




# 单元测试覆盖率

GOPATH/go-tutorial/testcover/foo.go

```go
package testcover

import "fmt"

func Foo(tag int) {
	switch tag {
	case 1:
		fmt.Println("Android")
	case 2:
		fmt.Println("Go")
	case 3:
		fmt.Println("Java")
	default:
		fmt.Println("C")

	}
}
```

GOPATH/go-tutorial/testcover/foo_test.go
```go
package testcover

import "testing"

func TestFoo(t *testing.T) {

	Foo(1)
	Foo(2)
	Foo(3)
}
```
```bash
cd GOPATH/go-tutorial
go test -v -coverprofile=c.out ./testcover
// use tool to generate html file from c.out
go tool cover -html=c.out -o=tag.html
```

在浏览器中打开生成的tag.html，可以更直观的看到testcovere包在每个文件的测试覆盖情况。哪一行被覆盖了。
