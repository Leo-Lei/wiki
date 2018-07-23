---
layout: post
title: Go Debugger Delve
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---


# Go的调试工具
在Go的官网，有一篇文章[https://golang.org/doc/gdb](https://golang.org/doc/gdb)说明了如何在Go中进行调试:

> Note that Delve is a better alternative to GDB when debugging Go programs built with the standard toolchain. It understands the Go runtime, data structures, and expressions better than GDB. Delve currently supports Linux, OSX, and Windows on amd64. For the most up-to-date list of supported platforms, please see the Delve documentation.

> GDB does not understand Go programs well. The stack management, threading, and runtime contain aspects that differ enough from the execution model GDB expects that they can confuse the debugger and cause incorrect results even when the program is compiled with gccgo. As a consequence, although GDB can be useful in some situations (e.g., debugging Cgo code, or debugging the runtime itself), it is not a reliable debugger for Go programs, particularly heavily concurrent ones. Moreover, it is not a priority for the Go project to address these issues, which are difficult.

总结下来就是说:
1. Go的debug工具有GDB，但是对Go的支持不好。
2. Delve是一个比GDB更好的调试go应用的工具。

GDB不是一款专门用于调试Go的工具，它还可以调试C++代码，它对Go中的协程等支持不好，所以基本认为使用GDB来调试Go代码不可用。截止到本文编写的时候(2018年7月23日)，Delve是调试Go应用的最好的工具:
1. Delve对Go协程的调试支持很好
2. Go官方文档中推荐了Delve
3. GoLand这款IDE也使用Delve来进行Go应用的调试(比如远程调试，attach进程调试等)。

# 安装Delve

使用go get来安装Delve:
```bash
go get github.com/derekparker/delve/cmd/dlv
```
注：本文章使用的Go版本为Go1.8。








