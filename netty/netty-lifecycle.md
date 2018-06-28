---
layout: post
title: Netty生命周期
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---

Netty中的关键对象有：
* Channel
* Pipeline
* ChannelHandlerContext
* Handler
* Bootstrap
* ServerBootstrap

这节要讨论的问题主要是Netty中这些关键对象的生命周期。以及这些对象是什么级别的，或者这样问，一个连接中只有一个响应的实例，还是所有连接中只有一个实例，还是每来一个消息都会生成一个相应的实例。如果这些问题不弄清楚，对开发会有影响。

> 如果Netty的底层实现没弄清楚，是不太敢用Netty的。


# Channel

|           Status        |                      desc                     |
| ----------------------- | --------------------------------------------- |
| channelUnregistered     | channel已创建但没注册到一个EventLoop            |












