---
layout: post
title: Java
date: 2017-04-21 15:40:00
tags:
- Java
categories: Java
description: web-authentication
---

# java 内存分布
1. 程序计数器：每一条java线程都有一个独立的程序计数器。当前线程字节码的执行行号指示器。如果执行的是java方法，计数器记录的就是正在执行的虚拟机字节码的指令地址。
2. 虚拟机栈： 也是线程私有的。java虚拟机栈描述的是java方法执行的内存模型，每个方法被执行的时候都会同时创建一个栈帧（Stack Frame）用于存储局部变量表，操作数栈，动态链路，方法出口等信息。每一个方法被调用直至执行完成的过程，就对应着一个栈帧在虚拟机栈中从入栈到出栈的过程。
3. java堆：堆分为新生代和老年代。新生代又分为：Eden，From Survivor，To Survivor。
4. 方法区：是各个线程共享的。保存了类信息，常量，静态变量，编译器编译后的字节码等数据。





# jmap
```
jmap -histo 4939
```
> 使用jmap命令时，必须切换到进程属于的用户。
