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


# Zookeeper的Standalone部署
### 1. 配置conf文件
进入`<zookeeper/conf>`目录, 有一个`zoo_sample.cfg` 文件. 将它重命名成`zoo.cfg`.        
```text
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/opt/data/zookeeper
clientPort=2181
```
### 2. 启动Zookeeper
1. 进入`<zookeeper>/bin`文件夹.    
2. 运行`zkServer.sh start`启动zookeeper服务端。    
3. run `zkCli.sh -server 127.0.0.1:2181`启动zookeeper客户端。

# Zookeeper的集群部署
Zookeeper集群的节点数量为奇数比较合适。
假设我们要搭建一个3个节点的集群。有3台机器：
* 192.168.1.81     该机器上需要开放3个端口:2181,2881,3881
* 192.168.1.82     该机器上需要开放3个端口:2182,2882,3882
* 192.168.1.83     该机器上需要开放3个端口:2183,2883,3883
每台机器上都安装zookeeper。

### 1.配置zoo.cfg
将conf/zoo_sample.cfg重命名成conf/zoo.cfg。
机器`192.168.1.81`的配置文件`conf/zoo.cfg`    
```conf
tickTime=2000
initLimit=10
syncLimit=5

dataDir=/opt/data/zookeeper
dataLogDir=/opt/data/zookeeper/logs

clientPort=2181

server.1=192.168.1.81:2881:3881
server.2=192.168.1.82:2882:3882
server.3=192.168.1.83:2883:3883
```

机器`192.168.1.82`的配置文件`conf/zoo.cfg`    
```conf
tickTime=2000
initLimit=10
syncLimit=5

dataDir=/opt/data/zookeeper
dataLogDir=/opt/data/zookeeper/logs

clientPort=2182

server.1=192.168.1.81:2881:3881
server.2=192.168.1.82:2882:3882
server.3=192.168.1.83:2883:3883
```

机器`192.168.1.83`的配置文件`conf/zoo.cfg`    
```conf
tickTime=2000
initLimit=10
syncLimit=5

dataDir=/opt/data/zookeeper
dataLogDir=/opt/data/zookeeper/logs

clientPort=2183

server.1=192.168.1.81:2881:3881
server.2=192.168.1.82:2882:3882
server.3=192.168.1.83:2883:3883
```








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

