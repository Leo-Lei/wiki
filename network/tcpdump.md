---
layout: post
title: tcpdump
date: 2017-08-16 14:30:00
tags:
- docker
categories: Java
---


|                              命令                                   |                          说明                        |
| ------------------------------------------------------------------- | ---------------------------------------------------- |
| `sudo tcpdump -D`                                                   | 查看本地网络接口列表                                   |
| `sudo tcpdump -A -i lo0 '( (tcp) and (port 8088) )'`                | 监听本地localhost回环地址网卡，tcp协议，端口8088的流量  |

