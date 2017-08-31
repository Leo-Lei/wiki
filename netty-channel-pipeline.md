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

