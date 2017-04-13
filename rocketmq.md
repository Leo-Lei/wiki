---
layout: post
title: RocketMQ
date: 2017-04-01 11:10:00
tags:
- docker
categories: Java
description: Rocket MQ
---

# 安装RocketMQ

```bash
git clone https://github.com/apache/incubator-rocketmq.git
cd incubator-rocketmq
mvn clean package install -Prelease-all assembly:assembly -U
cd target/apache-rocketmq-all/
```
> 注意的地方：
> * rocketmq的启动脚本中默认分配的内存很大，4G和8G，如果电脑配置不够，需要将这些配置降低。
> * 如果机器的内存小，启动nameserver和broker时，可能不会报错无法分配内存，但是可能会影响rocketmq的正常启动。我遇到过broker可以启动，日志中也打印出来了boot success。但是看nameserver日志，有很多的unregister broker。消息无法正常发送和接收。

# RocketMQ 架构
![https://static.oschina.net/uploads/img/201609/28105945_t8eA.png](https://static.oschina.net/uploads/img/201609/28105945_t8eA.png)
