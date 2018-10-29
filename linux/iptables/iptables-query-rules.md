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
| `iptables -t filter -L INPUT`             | 查看filter表中INPUT链的规则                                         |
| `iptables -t filter -vL INPUT`            | 查看filter表中INPUT链的规则。`-v`查看更详细的信息                    |
| `iptables -t filter -nvL INPUT`           | 查看filter表中INPUT链的规则。`-n`不对IP地址进行名称反解，直接显示IP地址。比如不显示`anywhere`，显示`0.0.0.0/0`    |
| `iptables --line-number -L`               | 查看filter表的规则。`--linenumber`显示规则编号                       |                                                    |
| `iptables --line -L`                      | `--line`等价于`--line-number`                                      |

# 规则中的字段信息
使用如下命令来查看filter表中的规则：
```bash
iptables -t filter -vL INPUT 
CHAIN INPUT (policy ACCEPT 0 packets, 0 bytes)
pkts  bytes  target    prot   opt   in   out    source         destination
239   20792  ACCEPT    all    --    any  any    anywhers       anywhere       

```
* `-t`: 指定要操作的表。
* `-L`: 列出规则。
* `pkts`: 对应匹配到的报文的个数
* `bytes`: 对应匹配到的报文包的大小总和
* `target`: 规则对应的target，表示规则的“动作”，即规则匹配后需要采取的措施
* `port`: 规则对应的协议，是否只针对某些协议应用此规则。
* `opt`: 表示规则对应的选项
* `in`: 表示数据包由哪个接口(网卡)流入。可以设置通过哪块网卡流入的报文需要匹配当前规则。
* `out`: 表示数据包由哪个接口(网卡)流出。
* `source`: 表示规则对应的源地址，可以是一个IP，也可以是一个网段。
* `destination`: 表示规则对应的目标地址。可以是一个IP，也可以是一个网段。
* `POLICY ACCEPT`: 当前链的默认策略, 默认动作。
* `0 packets`: 当前链默认策略匹配到的包数量。
* `0 bytes`: 当前链默认策略匹配到的所有包的大小总和。


