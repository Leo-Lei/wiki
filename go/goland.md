---
layout: post
title: GoLand
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# 快捷键

|            Key                 |                     Desc                                |
| ------------------------------ | ------------------------------------------------------- |
| `Command` + `N`                | Open Type                                               |
| `Command` + `Shift` + `N`      | Open File                                               |
| `Command` + `B`                | Declaration                                             |
| `Control` + `Option` + `H`     | Call Hierarchy                                          |


# 保存Go文件时自动格式化

GoLand --> Preferences -->  Tools --> File Watchers
添加一个`go fmt`即可。


# GoLand设置Project级别的GOPATH

针对每个项目配置一个Project GOPATH。
global GOPATH中不要放项目代码。

github上项目地址：
github.com/${github_name}/{project_name}

项目使用的go package为：foo/bar

将项目clone到下面的目录：
* /home/admin/projectA1/src/foo/bar/{project_name}
* /home/admin/projectA2/src/foo/bar/{project_name}


在GoLand中分别为项目设置project级别的GOPATH：
* /home/admin/projectA1/src/foo/bar/{project_name}：/home/admin/projectA1
* /home/admin/projectA2/src/foo/bar/{project_name}：/home/admin/projectA2

注意：
* Project级别GOPATH是GoLand的概念，GO原生是不支持的。Go原生只有GOPATH的概念。GoLand中设置了Project级别的GOPATH，GoLand会自动帮我们把Project GoPath添加到全局GoPath中。类似于export GOPATH=$GLOBAL_GOPATH:$PROJECT_GOPATH。这个操作只有在使用GOLAND来触发时才生效，并且是在当前会话级别生效。不会影响环境变量中的GOPATH。
* 如果通过GoLand的按钮来Run，Debug，Project GOPATH会生效。如果是在shell中直接执行go build，Go是不认识Goland中设置的Project GOPATH的，只认识环境变量中的GOPATH。shell中执行go build时，需要手动export GOPATH=$GOPATH:/home/admin/projectA1
 

