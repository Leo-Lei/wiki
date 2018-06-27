---
layout: post
title: Netty 线程模型
date: 2016-11-16 10:20:00
tags:
- Netty
categories: Netty
---



|    Name          |  boss group   |    worker group    |                                                                               code                                                               |
| ---------------- | ------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| 单线程           |       1       | 和boss group同一个  | EventLoopGroup bossGroup = new NioEventLoopGroup(1); <br> ServerBootStrap b = new ServerBootstrap();<br> b.group(bossGroup)                      |         
| 多线程           |       1       | 多个                | EventLoopGroup bossGroup = new NioEventLoopGroup(1); <br> EventLoopGroup workerGroup = new NioEventLoopGroup(); <br> ServerBootStrap b = new ServerBootstrap();<br> b.group(bossGroup, workerGroup)  |         
