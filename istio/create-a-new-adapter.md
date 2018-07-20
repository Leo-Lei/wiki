---
layout: post
title: Create a New Adapter in Istio
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---



本文章会介绍如何在Istio中新建一个Adapter。该Adapter:
1. 支持Istio自带的metric模板
2. 对于每个请求，Adapter会把它接收到的数据打印到一个文件中


# 准备工作
### 安装go
安装go，配置`GOPATH`。
### 下载Istio源码

```bash
git clone git clone https://github.com/istio/istio
```
将istio源码复制到你的GOPATH中。比如`$GOPATH/src/istio.io/istio`。

### 安装protoc(protobuff编译器)
Istio中的template，adapter等代码都是根据proto文件来生成的，所以需要安装protoc。
> 注: 和protobuf的通常用法不一样，在Istio中，我们会先定义一个`template.proto`文件，是一个标准的protobuf格式的文件。然后Istio根据自己的脚本和protoc来生成对应的go代码。包括Template，Instance，Adapter，Handler等。而不只是用于序列化的代码。

从[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)下载protoc。并将protoc添加到PATH中。



