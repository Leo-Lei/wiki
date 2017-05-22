---
layout: post
title: Zookeeper
date: 2016-08-11 17:10:00
tags:
- Java
categories: Java
description: web-authentication
---

# Overview               

# 安装 Zookeeper
Zookeeper目前还没有提供yum安装的方式。所以，要下载zookeeper的二进制文件安装。


# zookeeper的Standalone部署
### 配置conf文件
进入`<zookeeper/conf>`目录, 有一个`zoo_sample.cfg` 文件. 将它重命名成`zoo.cfg`.        
```text
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/opt/data/zookeeper
clientPort=2181
```
### 启动Zookeeper
1. 进入`<zookeeper>/bin`文件夹.    
2. 运行`zkServer.sh start`启动zookeeper服务端。    
3. run `zkCli.sh -server 127.0.0.1:2181`启动zookeeper客户端。

# Zookeeper命令行接口
```bash
[zk: localhost:2181(CONNECTED) 0] ls /
[zookeeper]
[zk: localhost:2181(CONNECTED) 1] create /app app
Created /app
[zk: localhost:2181(CONNECTED) 2] ls /
[app, zookeeper]
[zk: localhost:2181(CONNECTED) 3] create /app/app1 app1
Created app/app1
[zk: localhost:2181(CONNECTED) 4] rmr /app
```

# Resources

[https://www.ibm.com/developerworks/cn/opensource/os-cn-zookeeper](https://www.ibm.com/developerworks/cn/opensource/os-cn-zookeeper)

