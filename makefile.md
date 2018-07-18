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
上面代码中，构建result.txt的前置条件是source.txt。如果当前目录中，source.txt存在，那么make result.txt可以正常运行。否则必须再写一个规则，来生成source.txt。


**命令**

命令表示如何更新目标文件。由一行或多行的Shell命令组成。它是构建目标的具体命令。它的运行结果通常就是生成目标文件。    
每行命令之前必须有一个Tab键。如果想用其他键，可以用内置变量.RECIPEPREFIX声明。
```bash
.RECIPEPREFIX = >
someTarget:
> echo Hello, world
```
 注意，每行命令在一个单独的shell中执行。这些shell之间没有继承关系。
 ```bash
 someTarget:
     export foo=bar
     echo "foo=[$$foo]"
 ```
上面的代码，make someTarget，取不到foo的值。因为两个命令在两个不同的进程中运行。一个解决办法是将两行命令写在一起
```bash
 someTarget:
     export foo=bar; echo "foo=[$$foo]"
```
另一个解决办法是在换行符前加反斜杠转义
```bash
 someTarget:
     export foo=bar; \
     echo "foo=[$$foo]"
 ```





