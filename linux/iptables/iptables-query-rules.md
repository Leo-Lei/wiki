---
layout: post
title: Linux iptables规则查询
date: 2016-07-20 13:50:00
tags:
- Linux
categories: Linux
---

# 一些iptables的规则命令 

|            Command                        |                              Desc                                  |
| ----------------------------------------- | ------------------------------------------------------------------ |
| `iptables -t filter -L`                   | 查看filter表的规则                                                  |
| `iptables -t raw -L`                      | 查看raw表的规则                                                     |
| `iptables -L`                             | 查看filter表的规则。如果不指定-t，默认使用filter表                    |
| 





# 操作规则实例
使用如下命令来查看filter表中的规则：
```bash
iptables -t filter -L 
CHAIN INPUT (policy ACCEPT)
target    port   opt  source                 destination

CHAIN FORWARD (policy ACCEPT)
target    port   opt  source                 destination

CHAIN OUTPUT (policy ACCEPT)
target    port   opt  source                 destination
```
* `-t`: 指定要操作的表。
* `-L`: 列出规则。

