---
layout: post
title: Netty ChannelHandler
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---



# ChannelHandler

|          handler                    |                      Desc                    |
| ----------------------------------- | -------------------------------------------- |
| ChannelInboundHandler               | 拦截和处理入站事件                              |
| ChannelOutboundHandler              | 拦截和处理出站事件                              |
| ChannelInboundHandlerAdaptor        |                                              |
| SimpleChannelInboundHandler         |                                              |


# ChannelHandler生命周期

|         Handler       |                     Desc                        |
| --------------------- | ----------------------------------------------- |
| handlerAdd            | 把ChannelHandler添加到ChannelPipeline中时被调用    |
| handlerRemoved        | 把ChannelHandler从ChannelPipeline中移除时被调用    |
| exceptionCaught       | 当处理过程中在ChannelPipeline中有错误时发生          |











# 编码器






# 解码器







# SimpleChannelInboundHandler


