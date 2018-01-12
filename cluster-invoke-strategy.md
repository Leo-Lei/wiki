---
layout: post
title: Mongo
date: 2017-04-11 13:05:00
tags:
- Java
categories: Java
---


# 调用策略和负载均衡的区别
假设有10个服务端组成的集群，那么
* 负载均衡：在这10个服务端中，选择哪个来调用，比如有：随机，轮询等
* 调用策略：如果本次调用失败了，接下来该如何处理。

|            Strategy          |                              Desc                            |
| ---------------------------- | ------------------------------------------------------------ |
| FailfaseClusterInvoker       | 只调用一次，失败立即返回错误                                      |
| FailoverClusterInvoker       | 重试n次，每次更新服务端列表，保证失败后切换到另一个服务端              |
| FailbackClusterInvoker       | select后发起一次调用，若失败则将invoker加入失败列表，定期重试        |
| FailsafeClusterInvoker       | select后发起一次调用，若失败忽略异常，返回一个空Result               |
| AvailableClusterInvoker      | 不进行select，进选取第一个available的Invoker                     |
| BroadCastClusterInvoker      | 不进行select，每个Invoker都调用一次                               |
| ForkingClusterInvoker        | 不进行select，对所有分组并行发起调用，最快的调用完成后，返回结果        |
