---
layout: post
title: Zookeeper
date: 2017-07-05 17:10:00
tags:
- Java
categories: Java
description: web-authentication
---

# Overview               

# 安装 Zookeeper
Zookeeper目前还没有提供yum安装的方式。所以，要下载zookeeper的二进制文件安装。安装过程就是将zookeeper的二进制文件解压到某个目录即可。


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
* 192.168.1.100     该机器上需要开放3个端口:2181,2881,3881
* 192.168.1.101     该机器上需要开放3个端口:2182,2881,3881
* 192.168.1.102     该机器上需要开放3个端口:2183,2881,3881
每台机器上都安装zookeeper。

### 1.配置zoo.cfg
将conf/zoo_sample.cfg重命名成conf/zoo.cfg。
机器`192.168.1.100`的配置文件`conf/zoo.cfg`    
```conf
tickTime=2000
initLimit=10
syncLimit=5

dataDir=/opt/data/zookeeper/data
dataLogDir=/opt/data/zookeeper/log

clientPort=2181

server.1=192.168.1.100:2881:3881
server.2=192.168.1.101:2881:3881
server.3=192.168.1.102:2881:3881
```

机器`192.168.1.101`的配置文件`conf/zoo.cfg`    
```conf
tickTime=2000
initLimit=10
syncLimit=5

dataDir=/opt/data/zookeeper/data
dataLogDir=/opt/data/zookeeper/log

clientPort=2182

server.1=192.168.1.100:2881:3881
server.2=192.168.1.101:2881:3881
server.3=192.168.1.102:2881:3881
```

机器`192.168.1.102`的配置文件`conf/zoo.cfg`    
```conf
tickTime=2000
initLimit=10
syncLimit=5

dataDir=/opt/data/zookeeper/data
dataLogDir=/opt/data/zookeeper/log

clientPort=2183

server.1=192.168.1.100:2881:3881
server.2=192.168.1.101:2881:3881
server.3=192.168.1.102:2881:3881
```

参数说明:
* tickTime=2000
tickTime这个时间是作为Zookeeper服务器之间或客户端与服务器之间维持心跳的时间间隔,也就是每个tickTime时间就会发送一个心跳。
* initLimit=10
initLimit这个配置项是用来配置Zookeeper接受客户端（这里所说的客户端不是用户连接Zookeeper服务器的客户端,而是Zookeeper服务器集群中连接到Leader的Follower 服务器）初始化连接时最长能忍受多少个心跳时间间隔数。当已经超过10个心跳的时间（也就是tickTime）长度后 Zookeeper 服务器还没有收到客户端的返回信息,那么表明这个客户端连接失败。总的时间长度就是 10*2000=20 秒。
* syncLimit=5
syncLimit这个配置项标识Leader与Follower之间发送消息,请求和应答时间长度,最长不能超过多少个tickTime的时间长度,总的时间长度就是5*2000=10秒。
* dataDir=/home/wusc/zookeeper/node-01/data
dataDir顾名思义就是Zookeeper保存数据的目录,默认情况下Zookeeper将写数据的日志文件也保存在这个目录里。
* clientPort=2181
clientPort这个端口就是客户端（应用程序）连接Zookeeper服务器的端口,Zookeeper会监听这个端口接受客户端的访问请求。
* server.A=B：C：D    
server.1=192.168.1.100:2881:3881    
server.2=192.168.1.101:2881:3881    
server.3=192.168.1.102:2881:3881    
A是一个数字,表示这个是第几号服务器；    
B是这个服务器的IP地址（或者是与IP地址做了映射的主机名）；    
C第一个端口用来集群成员的信息交换,表示这个服务器与集群中的Leader服务器交换信息的端口；    
D是在leader挂掉时专门用来进行选举leader所用的端口。    

### 2. 在dataDir=/opt/data/zookeeper下创建myid文件
`192.168.1.81`
```bash
echo 1 >> /opt/data/zookeeper/data/myid 
```
`192.168.1.82`
```bash
echo 2 >> /opt/data/zookeeper/data/myid 
```
`192.168.1.83`
```bash
echo 3 >> /opt/data/zookeeper/data/myid 
```
### 3. 启动zookeeper
启动`192.168.1.100`上的zookeeper
```bash
/opt/zookeeper/bin/zkServer.sh start
```
启动`192.168.1.101`上的zookeeper
```bash
/opt/zookeeper/bin/zkServer.sh start
```
启动`192.168.1.102`上的zookeeper
```bash
/opt/zookeeper/bin/zkServer.sh start
```

### 4. 查看zookeeper的状态
3台机器上，随便选一台，然后执行下面的命令:
```bash
/opt/zookeeper/bin/zkServer.sh status
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
# zookeeper集群中的角色
|             | 写?      | 读        | 投票?      | 备注                |
| ----------- | -------- | -------- | ---------- | ------------------ |
| leader      | Y        | Y        | Y          |                    |
| follower    | N        | Y        | Y          |                    |
| observer    | N        | Y        | N          | 当集群中的节点增加到一定程度，由于投票的压力增大而使得吞吐量降低，增加了Observer|

# zookeeper 选举算法

为什么需要leader？    
如果有两个客户端同时向zookeeper节点A和B发起，请求。client1的请求是`set name=john`。client2的请求是'set name=tom'。在分布式的系统中，该如何处理这2个请求呢？这时候就需要有一个leader来进行决策。

# zookeeper的高可用
只要获得超过最大节点数的半数，就是leader。所以：        
* 2个节点的集群可以启动，当一台挂机后，不可用。
* 3个节点的集群可以启动，当一台挂机后，仍然可用。
* 5个几点的集群可以运行有2台挂机。2n+1台的集群，允许有n台挂机。



# Curator
Apache Curator是一个Zookeeper的客户端，和另一个客户端zkclient比起来，Curator要更优秀。

# Resources

[https://www.ibm.com/developerworks/cn/opensource/os-cn-zookeeper](https://www.ibm.com/developerworks/cn/opensource/os-cn-zookeeper)
[http://www.cnblogs.com/zhangs1986/p/6564839.html?from=timeline](http://www.cnblogs.com/zhangs1986/p/6564839.html?from=timeline)
[http://curator.apache.org](http://curator.apache.org)
