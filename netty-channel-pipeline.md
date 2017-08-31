---
layout: post
title: Netty Channel Pipeline
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
---



# ChannelHandler

|          handler                    |                      Desc                    |
| ----------------------------------- | -------------------------------------------- |
| ChannelInboundHandler               | 拦截和处理入站事件                              |
| ChannelOutboundHandler              | 拦截和处理出站事件                              |







# 事件在ChannelPipeline中流动
```java
ChannelPipeline p = ...;
p.addLast("1", new InboundHandlerA());
p.addLast("2", new InboundHandlerB());
p.addLast("3", new OutboundHandlerA());
p.addLast("4", new OutboundHandlerB());
p.addLast("5", new InboundOutboundHandlerX());
```
对于入站事件，处理顺序为: 1 -> 2 -> 5。对于出站事件，顺序是: 5 -> 4 -> 3。    
事件不会在ChannelPipeline中自动流动，完全由用户控制
```java
public class InboundHandlerA implements ChannelInboundHandler {
        @Override
        public void channelActive(ChannelHandlerContext ctx) {
            System.out.println("Connected!");    // 用户自定义处理逻辑
            ctx.fireChannelActive();             // 将channelActive事件传播到InboundHandlerB
        }
    }

    public class OutboundHandlerB extends ChannelOutboundHandler{
        @Override
        public void close(ChannelHandlerContext ctx, ChannelPromise promise) {
            System.out.println("Closing ..");   // 用户自定义处理逻辑
            ctx.close(promise);                 // 将close事件传播到OutboundHandlerA
        }
    }
```


# 入站事件

入站事件一般由I／O线程触发。

|             event             |                                           |
| ----------------------------- | ----------------------------------------- |
| ChannelRegistered()           | channel注册到EventLoop                     |
| ChannelActive()               | channel激活                                |
| ChannelRead(Object)           | channel读取到数据                           |
| ChannelReadComplete           | channel读取数据完毕                         |
| ExceptionCaught(Throwable)    | 捕获到异常                                  |
| UserEventTriggered(Object)    | 用户自定义事件                               |
| ChannelWritablilityChanged()  | channel可写性改变，由写高低位控制              |
| ChannelInactive()             | Channel不再激活                             |
| ChannelUnregistered()         | channel从EventLoop中注销                    |


# 出站事件

出站事件一般由用户触发

|                          event                          |                                           |
| ------------------------------------------------------- | ----------------------------------------- |
| bind(SocketAddress, ChannelPromise)                     | 绑定到本地地址                              |
| connect(SocketAddress, SocketAddress, ChannelPromise)   | 连接一个远程机器                             |
| write(Object, ChannelPromise)                           | 写数据，写到Netty出站缓冲区                   |
| flush()                                                 | flush数据，实际执行底层写                     |
| read()                                                  | 读数据                                      |
| disconnect(ChannelPromise)                              | 断开连接                                     |
| close(ChannelPromise)                                   | 关闭channel                                 |
| deregister(ChannelPromise)                              | 从EventLoop注销channel                      |
 


