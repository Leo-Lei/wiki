---
layout: post
title: Netty Channel Inbound Handler
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---



# ChannelHandler

|                      method                     |                      Desc                    |
| ----------------------------------------------- | -------------------------------------------- |
| handlerAdded(ChannelHandlerContext ctx)         | handler被添加到Pipeline中时，被调用            |
| handlerRemoved(ChannelHandlerContext ctx)       | handler从Pipeline中被移除时，被调用            |
| exceptionCaught(ChannelHandlerContext ctx)      | 有异常时被调用                                |


# ChannelInboundHandler

|                      method                     |                      Desc                    |
| ----------------------------------------------- | -------------------------------------------- |
| channelRegistered(ChannelHandlerContext ctx)         | handler被添加到Pipeline中时，被调用            |
