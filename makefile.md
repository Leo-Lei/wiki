---
layout: post
title: Makefile
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---

# Makefile格式
```bash
<target> : <prerequisites>
    <commands>
```

### target

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


### 前置条件

前置条件通常是一组文件名，之间用空格分隔。它指定了目标是否重新构建的判断标准：只要有一个前置文件不存在，或者有更新，目标就需要重新构建。

```bash
result.txt: source.txt
    cp source.txt result.txt
```
上面代码中，构建result.txt的前置条件是source.txt。如果当前目录中，source.txt存在，那么make result.txt可以正常运行。否则必须再写一个规则，来生成source.txt。


### 命令

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


# Makefile文件语法
### 注释
```bash
# 这是注释
result.txt: source.txt
    cp source.txt result.txt   # 这是注释
```

### 回声

正常情况下，make会打印每条命令，然后再执行，叫着回声

```bash
test: 
    # 这是测试
```

```bash
$ make test
#这是测试
```
可以在命令前面加上@，就可以关闭回声。
```bash
test:
    @# 这是测试
```

### 通配符
Makefile的通配符和Bash一致，主要有星号* 问号？
```bash
clean:
    rm -rf *.txt
```

### 模式匹配

Make命令允许对文件名，进行类似正则运算的匹配，主要用到的匹配符是%。比如，假定当前目录下有 f1.c 和 f2.c 两个源码文件，需要将它们编译为对应的对象文件。
```bash
%.o: %.c
```
等同于下面的写法。
```bash
f1.o: f1.c
f2.o: f2.c
```
使用匹配符%，可以将大量同类型的文件，只用一条规则就完成构建。

### 变量和赋值

Makefile允许使用等号自定义变量。
```bash
txt = Hello World
test:
    echo $(txt)
```
上面代码中，变量txt等于Hello World。调用时，变量需要放在$()之中。    
调用Shell变量，需要在美元符号前，再加上一个美元符号，因为Make命令会对美元进行符号转义。
```bash
test:
    echo $$HOME
```

Makefile一共提供了4个赋值运算符:
```bash
VARIABLE = value
# 在执行时扩展，允许递归扩展。

VARIABLE := value
# 在定义时扩展。

VARIABLE ?= value
# 只有在该变量为空时才设置值。

VARIABLE += value
# 将值追加到变量的尾端。
```

```bash
HELLO = world
HELLO_WORLD = $(HELLO) world!

# This echoes "world world!"
echo $(HELLO_WORLD)

HELLO = hello

# This echoes "hello world!"
echo $(HELLO_WORLD)
```

```bash
HELLO = world
HELLO_WORLD := $(HELLO) world!

# This echoes "world world!"
echo $(HELLO_WORLD)

HELLO = hello

# Still echoes "world world!"
echo $(HELLO_WORLD)

HELLO_WORLD := $(HELLO) world!

# This echoes "hello world!"
echo $(HELLO_WORLD)
```

```bash
HELLO_WORLD = hello
HELLO_WORLD += world!

# This echoes "hello world!"
echo $(HELLO_WORLD)
```


### 内置变量

Make命令提供一系列内置变量，比如$(CC)指向当前使用的编译器，$(MAKE)指向当前使用的Make工具。

### 自动变量
$@:指代当前目标，比如make foo的$@就指代foo
```bash
a.txt:
    touch $@
```

```bash
a.txt:
    touch a.txt
```

$<: 指代第一个前置条件。比如规则是t: p1 p2,那么$<就指代p1。
```bash
a.txt: b.txt c.txt
    cp $< $@
```
等同于:
```bash
a.txt: b.txt c.txt
    cp b.txt a.txt
```

$? 指代比目标更新的所有前置条件，之间以空格分隔。比如，规则为 t: p1 p2，其中 p2 的时间戳比 t 新，$?就指代p2。

$^ 指代所有前置条件，之间以空格分隔。比如，规则为 t: p1 p2，那么 $^ 就指代 p1 p2 。

$* 指代匹配符 % 匹配的部分， 比如% 匹配 f1.txt 中的f1 ，$* 就表示 f1。

### 函数

Makefile还可以使用函数，格式如下:
```bash
$(function arguments)
或者
${function arguments}
```
