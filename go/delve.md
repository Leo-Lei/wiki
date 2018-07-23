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

# Delve调试之Hello World
编写一个如下的hello.go文件：
```go
func main()  {
	s := "Hello"
	s += " world"
	println(s)
}
```
文件很简单，只有一个main方法，打印一条"Hello World"。下面来使用Delve启动程序
```bash
dlv debug hello.go
```
运行这个命令，dlv会去编译hello.go代码，然后传一些参数给编译器，好让编译器编译出来更加方便调试的可执行文件，然后启动了你的程序，并且attach上去，这样，我们的终端或命令行就会停留在debug模式，应用只是启动了，但还没有开始执行，下面就可以进行调试了。
```bash
$ dlv debug hello.go
Type 'help' for list of commands.
(dlv) help

```
在main函数上设置一个断点:
```bash
(dlv) break main.main
Breakpoint 1 set at 0x1050123 for main.main() ./hello.go:3
```
输出信息里有设置的断点的信息，有断点位置，函数名，文件名和所在行数。用continue命令让程序运行到我们设置的断点位置:
```bash
> main.main() ./hello.go:3 (hits goroutine(1):1 total:1) (PC: 0x1050123)
     1: package main
     2: 
=>   3: func main()  {
     4:         s := "Hello"
     5:         s += " world"
     6:         println(s)
     7: }
     8: 

```
continue命令可以让程序运行到下一个断点的位置。
接下来可以使用next命令让程序运行到下一句话。就是单步测试了。如果想继续向下，可以再执行next或直接按回车。如果按回车，Delve会重复执行上一条命令。
```bash
(dlv) next
> main.main() ./hello.go:5 (PC: 0x1050146)
     1: package main
     2: 
     3: func main()  {
     4:         s := "Hello"
=>   5:         s += " world"
     6:         println(s)
     7: }
     8: 
     9: 
(dlv) print s
"Hello"
(dlv) 

```
可以使用print把变量的值打印出来:
```bash
(dlv) print s
"Hello"
(dlv)
```





