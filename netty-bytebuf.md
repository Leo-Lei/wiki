---
layout: post
title: Netty ByteBuf
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
---



# ByteBuf
### 获取状态

|          method                      |                      Desc                    |
| ------------------------------------ | -------------------------------------------- |
| ByteBuf markReaderIndex()            | 是否打开                                      |
| ByteBuf resetReaderIndex()           | 是否打开                                      |
| ByteBuf markWriterIndex()            | 是否打开                                      |
| ByteBuf resetWriterIndex()           | 是否打开                                      |



|          method                      |                      Desc                                                |
| ------------------------------------ | ------------------------------------------------------------------------ |
| boolean readBoolean()                | 从当前readerIndex读取一个字节，返回一个boolean。readerIndex会加1               |
| byte readByte()                      | 从当前readerIndex读取一个字节，返回一个byte。readerIndex会加1                  |


