---
layout: post
title: Netty 线程模型
date: 2016-11-16 10:20:00
tags:
- Netty
categories: Netty
---



|    Name          |  boss group |    worker group    |                                                                               code                                                               |
| ---------------- | ----------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| 单线程           |       1     | 和boss group同一个  | EventLoopGroup bossGroup = new NioEventLoopGroup(1); ServerBootStrap b = new ServerBootstrap();<br> b.group(bossGroup)                           |              
