---
layout: post
title: Java Synchronized
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---

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
