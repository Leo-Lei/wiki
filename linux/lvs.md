---
layout: post
title: LVS
date: 2018-10-30 11:50:00
tags:
- Java
categories: Linux
---


# LVS
LVS是Linux Virtual Server的缩写，是Linux内核中一个四层的负载均衡组件。
使用LVS是一个实现负载均衡的软件方案。
如果要实现负载均衡，有软件方案和硬件方案。硬件方案的性能好，但是成本比较高。软件方案的性能比硬件方案低，但是成本低，且运维简单，扩展能力强。
而实现负载均衡的软件也不只有LVS一种，也可以通过nginx，haproxy等软件实现负载均衡的目的。

LVS有多种模式，在不同模式下，LVS实现负载均衡的方式不同。先来了解一下LVS的NAT模型。此模型比较容易理解。

# LVS的NAT模型

![LVS的NAT模型](http://www.zsythink.net/wp-content/uploads/2017/07/070617_0124_1.png)



http://www.zsythink.net/wp-content/uploads/2017/07/070617_0124_3.png




# 参考资料
* [http://www.zsythink.net/archives/2134](http://www.zsythink.net/archives/2134)
