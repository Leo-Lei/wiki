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





假设有两台机器：
* 192.168.1.100(以下简称100机器)
* 192.168.1.200（以下简称200机器）
接下来，我们会在192.168.1.100上ping机器192.168.1.200。并在192.168.1.200上添加一些iptables规则，来进行操作演示。

1. 在192.168。1.100上ping机器192.168.1.200。ping命令可以得到响应。说明ping命令发送的报文已正常到达了192.168.1.200机器。
2. 在200机器上添加拒绝所有报文的请求。
在200机器上添加一条规则，拒绝100机器上所有的报文。
```bash
iptables -t filter -I INPUT -s 192.168.1.100 -j DROP
```
* -t：使用filter表。
* -I：插入规则。
* -j：指定动作为DROP。
* -s：当source IP为192.168.1.100时匹配。
此时，在100机器上去ping机器200。看看能否ping通。发现ping一直没有回应。看来iptables规则已经生效了。
3. 再来做一个实验，目前INPUT链中已经有了一条规则。拒绝所有的来自100机器的报文。如果在这个规则之后，再添加一个规则，接收所有来自100机器的报文，会发生什么呢？
```bash
iptables -t filter -A INPUT -s 192.168.1.100 -j ACCEPT
iptables -nvL INPUT
Chain INPUT (policy ACCEPT)
target     prot   opt   in    out    source         destination
DROP       all    --    *     *      192.168.1.100  0.0.0.0/0
DROP       all    --    *     *      192.168.1.100  0.0.0.0/0
```
* -A: append的意思，在当前链的尾部追加一个规则。此时，在100机器上ping机器200。发现任然ping不通。
