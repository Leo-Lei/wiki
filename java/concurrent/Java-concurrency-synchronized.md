---
layout: post
title: Java Synchronized
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---

# Synchronized的作用
Synchronized有3个作用：
1. 线程之间的互斥性
2. 保证共享变量的可见性
3. 解决重排序问题
大多数人对线程间的互斥性比较了解，但其实Synchronized还有确保可见性和有序性的作用。

# Synchronized实现原理

Sycncronized是基于Object的monitor的。

|  synchronize作用范围  |                   详情                  |
| -------------------- | -------------------------------------- |
| 非静态方法             | 获取对象的monitor                       |
| 静态方法               | 获取类的monitor,即类对应的Class的monitor |
| 对象                  | 获取对象的monitor                       |
| 类                    | 获取类的monitor                         |


如下是当2个线程同时访问一个方法时，这两个方法是否会阻塞的情况：

|  Thread1             |           Thread2           |         Result       |
| -------------------- | --------------------------- | -------------------- |
| synchronized方法1     | synchronized方法1           | 阻塞                  |
| synchronized方法1     | 非synchronized方法2         | 不阻塞                 |
| synchronized方法1     | synchronized方法2           | 阻塞                 |

# Synchronized的性能开销
Synchronized是一个比较重量级的操作。对系统的性能有比较大影响。所以，如果有其他解决方案，通常都避免使用Synchronized。

