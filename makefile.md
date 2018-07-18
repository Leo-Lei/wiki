---
layout: post
title: Makefile
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---

# Makefile格式:
```bash
<target> : <prerequisites>
    <commands>
```

**target:**

通常是文件名，指明需要构建的对象，比如hello.txt。目标可以是一个文件名，也可以是多个文件名。用空格分隔。
target也可以是某个操作的名字，称为伪目标。
```bash
clean:
    rm *.txt
```
上面的目标时候clean，它不是文件名，只是一个操作名，属于伪目标。当然操作名可以随便取。
但是，如果当前文件中，恰好有这么一个文件clean，那么这个命令就不会执行。因为Make发现clean文件已经存在了，就认为没有必要重新构建了，就不会执行指定的rm命令了。        
为了避免这种情况，可以明确声明clean是伪目标，写法如下:

```bash
.PHONY: clean
clean:
    rm o.txt
```
声明clean是伪目标后，make就不会检查是否存在clean文件，而是每次都会执行对应的命令。    
如果Make命令运行时没有指定目标，默认会执行Makefile文件的第一个目标。


**前置条件**

前置条件通常是一组文件名，之间用空格分隔。它指定了目标是否重新构建的判断标准：只要有一个前置文件不存在，或者有更新，目标就需要重新构建。

```bash
result.txt: source.txt
    cp source.txt result.txt
```
上面代码
