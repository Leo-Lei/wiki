---
layout: post
title: iptables规则管理
date: 2016-07-20 13:50:00
tags:
- Linux
categories: Linux
---


|                               Command                                   |                                                    |
| ----------------------------------------------------------------------- | -------------------------------------------------- |
| `iptables -F INPUT`                                                     | 清空filter表的INPUT链中的规则                        |
| `iptables -t filter -I INPUT -s 192.168.1.146 -j DROP`                  | 向filter表的INPUT链中插入一个规则。报文的源地址是192.168.1.146时，执行DROP。   |          



