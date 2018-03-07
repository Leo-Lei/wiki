---
layout: post
title: Java Thread
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---



# 线程状态
* New: 新建状态。当线程创建完成时为新建状态，即new Thread(...)，但还没有调用start方法时。
* Runnable：就绪状态。当调用线程的start方法后，线程进入就绪状态，等待CPU资源。
* Running：运行状态。就绪状态的线程获取到CPU执行权后进入运行状态，开始执行run方法。
* Blocked：阻塞状态。线程没有执行完，由于某种原因，如I/O操作，线程同步等。让出CPU执行权，自身进入阻塞状态。
* Dead：死亡状态。线程执行完成，或者执行过程中出现异常，线程进入死亡状态。

