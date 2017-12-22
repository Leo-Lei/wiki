---
layout: post
title: Netty Channel
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---



# Channel

|          method                     |                      Desc                                 |
| ----------------------------------- | --------------------------------------------------------- |
| eventloop                           | 返回分配给Channel的EventLoop                                |
| pipeline                            | 返回分配给Channel的ChannelPipeline                          |
| isActive()                          | 是否激活                                                   |
| localAddress                        | 返回本地的SockerAddress                                     |
| remoteAddress                       | 返回远程的SockerAddress                                     |
| write                               | 将数据写到远程节点，数据将被传递给ChannelPipeline              |
| flush                               | 将之前写的数据flush到底层传输，如一个Socket                     |
| writeAndFlush                       |                                                           |


# Channel生命周期
|         Status        |                              Desc                            |
| --------------------- | ------------------------------------------------------------ |
| ChannelUnregistered   | Channel已被创建，但还没注册到EventLoop                           |
| ChannelRegistered     | Channel已被注册到EventLoop                                     |
| ChannelActive         | Channel处于活动状态(已连接到远程节点)。它现在可以接受和发送数据了      |
| ChannelInactive       | Channel还没有连接到远程节点                                      |


```text

  +---------------------+                   +---------------------+
  |  ChannelRegistered  |------------------>|   ChannelActive     |
  +---------------------+                   +---------------------+
                                                       |
                                                       |
  +---------------------+                   +---------------------+
  | ChannelUnregistered |<------------------|   ChannelInactive   |
  +---------------------+                   +---------------------+
```

当Channel的状态发生变化时，会生成对应的事件。这些事件会被转发给ChannelPipeline中的ChannelHandler。这些ChannelHandler可以做出响应。        



