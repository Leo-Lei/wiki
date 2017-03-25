---
layout: post
title: Dubbo
date: 2016-09-20 15:40:00
tags:
- Java
categories: Java
description: Dubbo
---

# Overview               
![Dubbo架构图](http://dubbo.io/dubbo-architecture.jpg-version=1&modificationDate=1330892870000.jpg)    
**节点角色**      
* Provider: 暴露服务的服务提供方。
* Consumer: 调用远程服务的服务消费方。
* Registry: 服务注册与发现的注册中心。
* Monitor: 统计服务的调用次调和调用时间的监控中心。
* Container: 服务运行容器。


# Dubbo 支持的协议               
|      Protocol     |        传输协议         | 序列化方式   |     跨语言       |       性能        |
| ----------------- | ----------------------- | ------------ | ---------------- | ----------------- |
| hession           |   http                  |     二进制   |                  |                   |


