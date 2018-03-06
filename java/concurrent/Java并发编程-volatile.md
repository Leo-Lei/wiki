---
layout: post
title: Java Synchronized
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---

# Volatile的作用
Volatile有2个作用：
1. 保证共享变量的可见性
2. 解决重排序问题

# Java内存模型
假设有变量count=10, 线程a对其进行读写
### 没有volatile修饰
1. 线程a从主存中读取count=10，将count=10保存到线程a的工作内存中。
2. 线程a在工作内存中进行操作count=count+1，所以，工作内存中count=11。
3. 这个时候，主存中的count还是10。工作内存中的count=11还没有同步到主存中。
4. 经过了一段不确定的时间，线程a的工作内存中的count=11同步到了主存中。
5. 当线程退出的时候，也会把count=11同步到主存中。

### 有volatile修饰
count=10，线程a和b对其进行读写
1. 线程a从主存中读取count=10，保存到线程a的工作内存中
2. 线程b从主存中读取count=10，保存到线程b的工作内存中
3. 线程a将count自增，count=11
4. 由于count是valatile的，当线程a将count=11写入线程a的工作线程中时，count=11立即被同步到了主存中
5. 同时，CPU中的所有线程的工作缓存中的count都被设置为无效了。
6. 当线程b需要读取count的值时，因为线程b的工作线程中的count已无效了。线程b从主存中去读取，count=11。是最新的值

# Volatile实现可见性
可见性问题是指一个线程修改了共享变量值，而另一个线程却看不到，这是由于JVM的内存模型决定的。每个线程在CPU中都有自己的一个高速缓存区--线程工作内存。volatile可以解决这个问题。
先看下面的代码：
```java
boolean stop = false;

// 线程1
while(!stop){
    doSomething();
}

// 线程2
stop = true;
```
这是一段很典型的代码，在线程2中将stop设置为true，来中断线程1的工作。但这段代码，不一定能正常工作。
