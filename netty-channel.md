---
layout: post
title: Netty Channel
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
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

