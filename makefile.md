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

**target**
通常是文件名，指明需要构建的对象，比如hello.txt。目标可以是一个文件名，也可以是多个文件名。用空格分隔。
target也可以是某个操作的名字，称为伪目标。
```bash
clean:
    rm *.txt
```
上面的目标时候clean，它不是文件名，只是一个操作名，属于伪目标。当然操作名可以随便取。

