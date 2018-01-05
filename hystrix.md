---
layout: post
title: Hystrix
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# 服务雪崩效应
分布式系统中某个基础服务不可用，最后导致整个系统都不可用，叫雪崩效应。

# 服务雪崩效应的定义
服务雪崩效应是因为**服务提供者**的不可用，导致**服务调用者**的不可用，并**逐渐放大**的过程。        
![服务雪崩](https://segmentfault.com/img/bVziad)






