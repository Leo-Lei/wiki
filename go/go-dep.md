---
layout: post
title: Go dep
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---


Dep是Go语言官方推出的一个依赖管理工具。

Dep通过两个文件来管理依赖:
1. `Gopkg.toml`
2. `Gopkg.lock`

其中`Gopkg`通过命令生成，也可以被用户根据需要手动修改。`Gopkg.lock`是自动生成的，不可修改。
Gopkg.lock文件取决于源代码中的import语句和Gopkg.toml。

![](https://blog.boatswain.io/img/manage-go-dependencies-using-dep-01.png)

# 安装dep
截止到Go1.10.1，dep也没被包含在Go的工具包中，需要独立安装。可以执行以下命令来安装:
```bash
go get -u github.com/golang/dep/cmd/dep
```
安装完后，执行`dep version`验证是否安装成功。

> 注意：如果`$GOPATH/bin`不在`PATH`下，需要将生成的`dep`文件复制到`$GOBIN`下。

# Gopkg.toml语法
**required**
