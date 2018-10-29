---
layout: post
title: Linux IP tables
date: 2016-07-20 13:50:00
tags:
- Linux
categories: Linux
---


# 链的概念

链可以比喻成一个关卡，数据在传送的过程中会经过不同的关卡，比如INPUT，FORWARD，OUTPUT。这些关卡在iptables中就叫做链。

Iptables中包含了一些链，如下所示：
![iptables](http://www.zsythink.net/wp-content/uploads/2017/02/021217_0051_2.png)
根据上图，在某些常用的场景下，报文的流向是：
* 到本机某进程的报文：PREROUTING --> INPUT
* 由本机转发的报文：PREROUTING --> FORWARD --> POSTROUTING
* 由本机的某进程发出报文：OUTPUT --> POSTROUTING

每个链下面，都会有多个规则。这些规则串在一起，就形成了一个链。每个经过这个关卡的报文，都要将这条链上的所有规则匹配一遍，如果有符合条件的规则，则执行对应的动作。

iptables中已经内置了一些链，可可以自定义一些链。

# 表的概念
在每个链上添加了一些规则，但是有些规则是很类似的，比如，A类规则都是对IP或端口进行过滤，B类规则则是修改报文。那么这个时候，可以把实现类似功能的规则放在一起。

具有类似功能的规则的集合叫做表。不同功能的规则，可以放置在不同的表中进行管理。iptables已经为我们定义了4种表：
|      表名        |                         功能                           |        内核模块       |
| ---------------- | ------------------------------------------------------ | -------------------- |
|   filter表       | 过滤。                                                  | iptables_filter     |




