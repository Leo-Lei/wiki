---
layout: post
title: Protobuff
date: 2017-06-30 16:30:00
tags:
- docker
categories: Java
---

# Protobuf
Protobuf是google开发的一个序列化格式，和XML，JSON类似。
Protobuf的一些优势：
1. 数据体积小，编解码速度快
2. 平台无关，语言无关。目前支持大多数主流语言。
Protobuf的一些劣势：
1. 数据没有可读性，不像XML和JSON。
2. 改动协议字段，需要重新生成文件。
3. 使用不方便，要想使用Protobuf，必须先定义IDL文件，安装Protobuf编译器，生成对应语言的代码文件。


# 安装
在go中使用protobuf，有两个可选用的包goprotobuf（go官方出品）和gogoprotobuf。
gogoprotobuf完全兼容google protobuf，它生成的代码质量和编解码性能均比goprotobuf高一些。
两者的代码仓库如下:
* go: github.com/golang/protobuf/
* gogo: github.com/gogo/protobuf/

### 安装protoc
从[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)下载protobuf的编译器protoc。然后将可自行文件拷贝到$GOPATH的bin目录下。

### golang/protobuf    
```bash
go get github.com/golang/protobuf/proto
```

golang/protobuf
安装插件
```bash
go get github.com/golang/protobuf/protoc-gen-go
```
生成go文件
```bash
# 编译当前目录的hello.proto文件，会生成一个hello.proto.go文件，生成的文件在当前目录
protoc --go_out=. hello.proto
```
### gogo/protobuf
1. 安装gogo/protobuf的库文件
```bash
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto  //这个不装也没关系
```
2. 安装gogo/protobuf的插件
有两个插件可以用：
* protoc-gen-gogo:和protoc-gen-go生成的文件差不多，性能也几乎一样(稍微快一点点)
* proto-gen-gofast:生成的文件更复杂，性能也更高(快5-7倍)
```bash
// protoc-gen-gogo
go get github.com/gogo/protobuf/protoc-gen-gogo
// proto-gen-gofast
go get github.com/gogo/protobuf/protoc-gen-gofast
```
3. 生成go文件
```bash
//gogo
protoc --gogo_out=. *.proto
//gofast
protoc --gofast_out=. *.proto
```

# Java中使用protobuf
```bash
protoc --java_out=/Users/leiwei/workspace/vangoleo/java-tutorial/protobuf student.proto
```

