---
layout: post
title: 在Mac上安装Java
date: 2017-07-11 13:05:00
tags:
- Java
categories: Java
---

# 安装单个版本的Java
去oracle的官网下载对应版本的JDK。是一个Mac系统标准的dmg格式的文件。双击直接安装即可。

# 安装多版本的Java
在Mac上安装Java7和Java8。
1. 去oracle官网下载Java7和Java8。
{% note info %} 下载Java7时，需要先登录，这个比较坑 {% endnote %}
2. 安装Java7和Java8
3. 编辑`~/.bash_profile`文件。
```bash
# 设置 JDK 7
export JAVA_7_HOME=`/usr/libexec/java_home -v 1.7`
# 设置 JDK 8
export JAVA_8_HOME=`/usr/libexec/java_home -v 1.8`

export JAVA_HOME=$JAVA_8_HOME

# alias命令动态切换JDK版本
alias jdk6="export JAVA_HOME=$JAVA_6_HOME"
alias jdk7="export JAVA_HOME=$JAVA_7_HOME"
alias jdk8="export JAVA_HOME=$JAVA_8_HOME"
```

看看java程序究竟在哪里
```
which java
/usr/bin/java
ls -lh /usr/bin/java
lrwxr-xr-x  1 root  wheel    74B 11 11 17:25 /usr/bin/java -> /System/Library/Frameworks/JavaVM.framework/Versions/Current/Commands/java
```

使用jdk7或jdk8来切换java的版本。



