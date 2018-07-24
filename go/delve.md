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
这是一个很简单的Hello World的调试，通过delve启动程序，在命令行中通过dlv命令设置断点，进行单步调试，并可以print变量的值。


# Attach到进程调试
待添加......

# GoLand + Delve：Attach到进程调试(远程调试)
前面的例子有使用Delve启动应用或attach到应用进行调试，但添加断点，运行到下一步，查看变量值等操作都是在终端中，输入delve命令来进行的。这种调试方式也太古老了，效率低下。真实场景下，几乎不会使用这种方式来进行调试。我们还是希望借助IDE进行更高效的调试。    
下面就介绍如何在GoLand中配合delve进行调试。包括附加到进程调试和远程调试。其实附加到本地进程和远程调试原理是一样的，待调试的进程是通过delve启动的，delve会启动进程，并立即附加到进程，开启一个debug session。并且启动一个debug server，暴露某个端口，客户端IDE可以通过该端口连接到debug server进行调试。        
***准备条件:***
* 安装好了GoLand。
* 安装好了delve，并将dlv添加到$PATH中。
### 附加到进程调试
假设需要调试如下的应用,我们会在本地运行该应用，附加到进程调试:
`hello.go`:
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
1. 编译`hello.go`
```bash
go build -gcflags=all="-N -l" hello.go 
```
其中-gcflags='all -N -l'是告诉go编译器不要进行编译器优化，不然可能会导致调试不对。编译器优化后可能会对代码进行重排序等。所以，为了调试，编译时最好带上-gcflags='all -N -l'。
> 注意：本文使用Go版本是:go1.10.1
> ```bash
> $go version
> go version go1.10.1 darwin/amd64
> ```
> 在网上绝大部分的blog都说加上`go build -gcflags='all -N -l'`。进测试，这个参数会导致go编译器错误。而且GoLand的调试页面也提供的是这个错误的参数，不知道是什么情况，难道大家都不用go调试的吗。。。

2. 使用delve启动hello

```bash
dlv --listen=:2345 --headless=true --api-version=2 exec ./hello
```
使用dlv启动了hello进程，并立即attach到hello进程，同时开启了一个debug server，暴露了端口2345来进行远程调试。

3. 在GoLand中配置Remote Debug
Run -> Edit Configurations -> Add new Configuration -> Go Remote:    
![GoLand Remote Debug](https://raw.githubusercontent.com/vangoleo/file-repo/master/goland_remote_debug.png)

4. 在hello.go中设置断点
5. 点击GoLand中的debug按钮，开始调试

### 调试带命令行参数的进程
假设hello进程启动时有参数如下:
```bash
./hello --id=2 --name=tom
```
那么在用dlv调试该带参数的进程如何做呢?
```bash
dlv --listen=:2345 --headless=true --api-version=2 exec ./hello -- --id=2 --name=tom
```
使用`--`来分隔dlv命令和hello的命令行参数。可以看到`exec ./hello`和`--id=2 --name=tom`之间有一个` -- `进行分隔。


