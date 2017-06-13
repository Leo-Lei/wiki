---
layout: post
title: RocketMQ
date: 2017-06-13 14:05:00
tags:
- docker
categories: Java
description: Rocket MQ
---

# 集群配置        
* 2台Name Server

# Name Server
|      Name      |            ip          |
| -------------- | ---------------------- |
| NameServer1    | 192.168.1.101          |
| NameServer2    | 192.168.1.102          |

分别启动NameServer        
```bash
nohup sh mqnameserver &
# 查看日志
tail -f -n 500 $ROCKETMQ_HOME/logs/rocketmqlogs/namesrv.log
```

# Broker
```bash
#所属集群名字
brokerClusterName=rocketmq-cluster
#broker名字
brokerName=broker-a
#0 表示 Master，>0 表示 Slave
brokerId=0
#nameServer地址，分号分割
namesrvAddr=192.168.1.101:9876;192.168.1.102:9876
#在发送消息时，自动创建服务器不存在的topic，默认创建的队列数
defaultTopicQueueNums=4
#是否允许 Broker 自动创建Topic，建议线下开启，线上关闭
autoCreateTopicEnable=true
#是否允许 Broker 自动创建订阅组，建议线下开启，线上关闭
autoCreateSubscriptionGroup=true
#Broker 对外服务的监听端口
listenPort=10911
#删除文件时间点，默认凌晨 4点
deleteWhen=04
#文件保留时间，默认 48 小时
fileReservedTime=120
#commitLog每个文件的大小默认1G
mapedFileSizeCommitLog=1073741824
#ConsumeQueue每个文件默认存30W条，根据业务情况调整
mapedFileSizeConsumeQueue=300000
#destroyMapedFileIntervalForcibly=120000
#redeleteHangedFileInterval=120000
#检测物理文件磁盘空间
diskMaxUsedSpaceRatio=88
#存储路径
storePathRootDir=/usr/local/alibaba-rocketmq/store
#commitLog 存储路径
storePathCommitLog=/usr/local/alibaba-rocketmq/store/commitlog
#消费队列存储路径存储路径
storePathConsumeQueue=/usr/local/alibaba-rocketmq/store/consumequeue
#消息索引存储路径
storePathIndex=/usr/local/alibaba-rocketmq/store/index
#checkpoint 文件存储路径
storeCheckpoint=/usr/local/alibaba-rocketmq/store/checkpoint
#abort 文件存储路径
abortFile=/usr/local/alibaba-rocketmq/store/abort
#限制的消息大小
maxMessageSize=65536

#flushCommitLogLeastPages=4
#flushConsumeQueueLeastPages=2
#flushCommitLogThoroughInterval=10000
#flushConsumeQueueThoroughInterval=60000
#Broker 的角色
#- ASYNC_MASTER 异步复制Master
#- SYNC_MASTER 同步双写Master
#- SLAVE
brokerRole=ASYNC_MASTER

#刷盘方式
#- ASYNC_FLUSH 异步刷盘
#- SYNC_FLUSH 同步刷盘
flushDiskType=ASYNC_FLUSH
#checkTransactionMessageEnable=false
#发消息线程池数量
#sendMessageThreadPoolNums=128
#拉消息线程池数量
#pullMessageThreadPoolNums=128
```
