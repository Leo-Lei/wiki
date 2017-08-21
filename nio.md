---
layout: post
title: Java NIO
date: 2016-11-16 11:50:00
tags:
- Java
categories: Java
description: Java NIO
---

# 基本概念    
* Channel: Channel像流，可以从Channel读到Buffer，也可以从Buffer写到Channel中。
* Buffer
* Selector: 允许单线程处理多个Channel。

# Channel
|          Channel         |             Description                                            |  
| ------------------------ | ------------------------------------------------------------------ |
| `FileChannel`            | 从文件读写数据。                                                      |               
| `DatagramChannel`        | 通过UDP读写网络中的数据。                                              |               
| `Sockethannel`           | 通过TCP读写网络中的数据。                                              |               
| `ServerSockethannel`     | 监听TCP连接，像Web服务器一样，对每一个新进来的连接都会创建一个SocketChanel。 |               

# Selector
通过Selector可以在一个线程中处理多个Channel，且是非阻塞的。    
向Selector注册一个或多个Channel，注册的时候，提供一些Selector需要监听的事件。




# Buffer

