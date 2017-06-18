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

# broker-a-master
```bash
brokerClusterName=DefaultCluster
namesrvAddr=172.31.24.81:9876;172.31.23.80:9876
brokerName=broker-a
brokerId=0

#是否允许 Broker 自动创建Topic，建议线下开启，线上关闭
autoCreateTopicEnable=false
#是否允许 Broker 自动创建订阅组，建议线下开启，线上关闭
autoCreateSubscriptionGroup=false

deleteWhen=04
fileReservedTime=48

brokerRole=SYNC_MASTER
flushDiskType=ASYNC_FLUSH
```





# Broker
```bash
#所属集群名字
brokerClusterName=DefaultCluster
#broker名字，注意此处不同的配置文件填写的不一样
brokerName=broker-a
#0 表示 Master，>0 表示 Slave
brokerId=0
#nameServer地址，分号分割
namesrvAddr=172.31.28.73:9876;172.31.19.73:9876
#在发送消息时，自动创建服务器不存在的topic，默认创建的队列数
#defaultTopicQueueNums=4
#是否允许 Broker 自动创建Topic，建议线下开启，线上关闭
#autoCreateTopicEnable=false
#是否允许 Broker 自动创建订阅组，建议线下开启，线上关闭
#autoCreateSubscriptionGroup=false
#Broker 对外服务的监听端口
#listenPort=10911
#删除文件时间点，默认凌晨 4点
deleteWhen=04
#文件保留时间，默认 48 小时
fileReservedTime=120
#commitLog每个文件的大小默认1G
#mapedFileSizeCommitLog=1073741824
#ConsumeQueue每个文件默认存30W条，根据业务情况调整
mapedFileSizeConsumeQueue=300000
#destroyMapedFileIntervalForcibly=120000
#redeleteHangedFileInterval=120000
#检测物理文件磁盘空间
#diskMaxUsedSpaceRatio=88
#存储路径
storePathRootDir=/opt/data/rocketmq/store
#commitLog 存储路径
#storePathCommitLog=/opt/data/rocketmq/store/commitlog
#消费队列存储路径存储路径
#storePathConsumeQueue=/opt/data/rocketmq/store/consumequeue
#消息索引存储路径
#storePathIndex=/opt/data/rocketmq/store/index
#checkpoint 文件存储路径
#storeCheckpoint=/opt/data/rocketmq/store/checkpoint
#abort 文件存储路径
#abortFile=/opt/data/rocketmq/store/abort
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
brokerRole=SYNC_MASTER

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
### 创建topic
```bash
sh mqadmin updateTopic -n 192.168.10.101:9876;192.168.10.102:9876 -c DefaultCluster -t mytopic
```






