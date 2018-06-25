---
layout: post
title: Netty Decoder
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---


# Decoder
Decoder就是解码器。解码器就是讲一个字节序列转换成另一种格式(一个消息或另一个字节序列)。
Decoder是用于处理入站数据的。Netty中的Decoder都继承于ChannelInboundHandler。





# LengthFieldBasedFrameDecoder
在RPC协议中，大的数据会分成多个包来发送，一般的协议会在消息头中定义长度字段来标识消息的总长度。这时候就可以使用LengthFieldBasedFrameDecoder来
进行合并，直到读到了完整的消息，才交给后面的handler处理。同时，这个decoder也可以对消息进行一些处理，比如丢弃掉其中的一些字节。
LengthFieldBasedFrameDecoder的作用:
1. 合并多个包的数据为一个完整的数据
2. 可以丢弃掉一些字节

|             arg            |                         desc                          | 
| -------------------------- | ----------------------------------------------------- |
| `lengthFieldOffset`        | 长度字段的偏移量                                         |
| `lengthFieldLength`        | 长度字段占的字节数                                       |
| `lengthAdjustment`         |                                                        |
| `initialBytesToStrip`      |                                                       |



