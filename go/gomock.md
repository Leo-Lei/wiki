---
layout: post
title: GoMock
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---


```bash
go get github.com/golang/mock/gomock
```

```bash
cd $GOPATH/src/github.com/golang/mock/mockgen
go install
```

```bash
mockgen
```



新建一个接口：
`GOPATH/go-tutorial/repo/repository.go`
```go
package repo

type User struct {
	Id   int
	Name string
	Age  int
}

type Repository interface {
	Create(user User)
	Get(id int) User
	Update(user User)
	Delete(id int)
}

```


```bash
mockgen go-tutorial/repo Repository
```

注意： 第一个参数是文件基于GOPATH的相对路径，第二个参数是要代码生成的接口，可以是多个。多个interface之间用逗号分隔，不能有空格。

mockgen命令支持如下选项:
* -source: 一个文件，包含要代码生成的interface
* -destination: 存放生成的代码文件路径。如果没有设置，代码将被打印到标准输出
* -package: 指定mock类源文件的包名。如果没有设置该选项，包名由`mock_`和输入文件的包名组成


```go
import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestMockRepository(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	mockRepo.EXPECT().Get(gomock.Any()).Return(User{Id: 1, Name: "tom", Age: 20})
	mockRepo.EXPECT().Get(gomock.Any()).Return(User{Id: 2, Name: "jerry", Age: 30})
	mockRepo.EXPECT().Get(gomock.Any()).Return(User{Id: 3, Name: "john", Age: 25})

	fmt.Println(mockRepo.Get(0)) // {1 tom 20}
	fmt.Println(mockRepo.Get(0)) // {2 jerry 30}
	fmt.Println(mockRepo.Get(0)) // {3 john 25}
	//fmt.Println(mockRepo.Get(0)) // exception
}

```
