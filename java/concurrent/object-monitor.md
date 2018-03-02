---
layout: post
title: Object的monitor
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---



每个对象有一个监视器锁（monitor）。当monitor被占用时就会处于锁定状态。JVM中有两个指令来获取和释放对象的monitor。
* monitorenter: 获取monitor
* monitorexit: 释放monitor

线程执行monitorenter指令时尝试获取monitor的所有权，过程如下：
1. 如果monitor的进入数为0，则该线程进入monitor，然后将进入数设置为1，该线程即为monitor的所有者。
2. 如果线程已经占有该monitor，只是重新进入，则进入monitor的进入数加1.
3. 如果其他线程已经占用了monitor，则该线程进入阻塞状态，直到monitor的进入数为0，再重新尝试获取monitor的所有权。

执行monitorexit的线程必须是objectref所对应的monitor的所有者。

指令执行时，monitor的进入数减1，如果减1后进入数为0，那线程退出monitor，不再是这个monitor的所有者。其他被这个monitor阻塞的线程可以尝试去获取这个 monitor 的所有权。 
