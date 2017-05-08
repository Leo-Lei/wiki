---
layout: post
title: Introduction to JVM
date: 2016-01-12 12:20:00
tags:
- Java
categories: Java
description: The tutoria will describe the useage of Linux.
---


# java 内存分布
1. 程序计数器：每一条java线程都有一个独立的程序计数器。当前线程字节码的执行行号指示器。如果执行的是java方法，计数器记录的就是正在执行的虚拟机字节码的指令地址。
2. 虚拟机栈： 也是线程私有的。java虚拟机栈描述的是java方法执行的内存模型，每个方法被执行的时候都会同时创建一个栈帧（Stack Frame）用于存储局部变量表，操作数栈，动态链路，方法出口等信息。每一个方法被调用直至执行完成的过程，就对应着一个栈帧在虚拟机栈中从入栈到出栈的过程。
3. java堆：堆分为新生代和老年代。新生代又分为：Eden，From Survivor，To Survivor。
4. 方法区：是各个线程共享的。保存了类信息，常量，静态变量，编译器编译后的字节码等数据。

# jmap
```bash
jmap -histo 4939
```
> 注意:使用jmap命令时，必须`su - user_name`切换到进程属于的用户。


```bash
jmap -dump:format=b,file=/opt/app/heap.bin 12080
```

# jstat

```bash
jstat -gc 4399
```

# JVM参数

|      option                  |              Desc                  |           Remark                             |
| ---------------------------- | ---------------------------------- | ------------------------------------------------- |
| `-Xms1024m`                  | 初始堆大小                           |  默认空余堆内存小于40%时，JVM会增大heap直到-Xmx的最大值  |
| `-Xmx2048m`                  | 最大堆大小                           | 默认空余堆内存大于70%时，JVM会减少堆直到-Xms的最小值     |
| `-Xmn600m`                   | 年轻代大小(1.4 or later)             | 是eden + 2 survior space的大小。该参数同事设置了年轻代的初始大小和最大值，即固定了年轻代的大小。增大年轻代后，将会减小老年代大小。Sun官方推荐配置为整个堆的3/8 |
| `-XX:NewSize`                | 年轻代初始化大小(for 1.3/1.4)         | 版本比较低的jvm使用该参数，不建议高版本使用。高版本jvm使用`-Xmn`参数即可 |
| `-XX:-UseAdaptiveSizePolicy` | 关闭自适应调整Eden和Survior空间的比例   | 如果使用的是ParallelGC作为GC算法，`-XX:SurvivorRatio`参数将不起作用。可以使用该参数将Eden和Surivor区的比例固定。    |                     
| `-XX:SurvivorRatio=8`        | Eden和surivor的大小比例              | `-XX:SurvivorRatio=8`，表示Eden:From Surivor:To Surivor=8:1:1 |                                        














