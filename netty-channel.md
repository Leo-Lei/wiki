---
layout: post
title: Netty BootStrap
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
---



# Handler
### 获取状态

|          method                     |                      Desc                    |
| ----------------------------------- | -------------------------------------------- |
| boolean isOpen()                    | 是否打开                                      |
| boolean isRegistered()              | 是否注册到一个EventLoop                        |
| boolean isActive()                  | 是否激活                                      |
| boolean isWritable()                | 是否可写                                      |

### getter方法

|          method                     |                      Desc                    |
| ----------------------------------- | -------------------------------------------- |
| EventLoop eventLoop()               | 获取注册到的EventLoop                          |
| Channel parent()                    | 父类channel                                   |




