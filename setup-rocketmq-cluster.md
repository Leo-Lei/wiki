---
layout: post
title: Setup RocketMQ Cluster
date: 2017-07-07 20:05:00
tags:
- docker
categories: Java
---

# RocketMQ版本
版本为4.1.0    
官方下载页面为:    
[http://rocketmq.apache.org/dowloading/releases/](http://rocketmq.apache.org/dowloading/releases/)    
二进制文件为:    
[rocketmq-all-4.1.0-incubating-bin-release.zip](https://www.apache.org/dyn/closer.cgi?path=incubator/rocketmq/4.1.0-incubating/rocketmq-all-4.1.0-incubating-bin-release.zip)            
对应的RocketMQ的client为:        
```bash
compile 'org.apache.rocketmq:rocketmq-client:4.1.0-incubating'
```


# 集群配置        
* nameserver-01
* nameserver-02
* broker-a-master
* broker-a-slave
* broker-b-master
* broker-b-slave


# Name Server
|      Name      |            ip           |
| -------------- | ----------------------- |
| NameServer1    | 192.168.10.101          |
| NameServer2    | 192.168.10.102          |

分别启动NameServer        
```bash
nohup sh mqnameserver &
# 查看日志
tail -f -n 500 $ROCKETMQ_HOME/logs/rocketmqlogs/namesrv.log
```

> rocketMQ默认运行nameserver和broker的jvm的内存都设置的很大，是8G。所以，如果我们的机器内存没有8G，那么请修改`runserver.sh`和`runbroker.sh`中的jvm内存参数。不然无法启动rocketmq

# broker-a-master
```bash
vim /opt/rocketmq/conf/2m-2s-sync/broker-a.properties

brokerClusterName=DefaultCluster
namesrvAddr=192.168.10.101:9876;192.168.10.102:9876
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
启动broker
```bash
cd /opt/rocketmq/bin
nohup sh mqbroker -c /opt/rocketmq/conf/2m-2s-sync/broker-a.properties &
```
> 注意，线上环境请一定设置autoCreateTopicEnable=false和autoCreateSubscriptionGroup=false。autoCreateTopicEnable默认值是true。如果为true，当发送一个topic=mytopic的消息时，如果broker中没有这个topic，则rockerMQ会自动创建topic，让消息可以正常发送到broker中并被消费。当rocketMQ只会将这个topic路由到集群中的某一个broker，后续的所有的该topic的消息都只会发送到该一个broker中，达不到负载均衡和fail over的效果。所以，线上环境我们必须要通过mqadmin工具手动的创建topic到集群中。只有这样才能达到负载均衡和高可用。不然集群是形同虚设的!!所以一定要将autoCreateTopicEnable设置为false。

<!-- more -->

# broker-a-slave
```bash
vim /opt/rocketmq/conf/2m-2s-sync/broker-a-s.properties

brokerClusterName=DefaultCluster
namesrvAddr=192.168.10.101:9876;192.168.10.102:9876
brokerName=broker-a
brokerId=1


#是否允许 Broker 自动创建Topic，建议线下开启，线上关闭
autoCreateTopicEnable=false
#是否允许 Broker 自动创建订阅组，建议线下开启，线上关闭
autoCreateSubscriptionGroup=false

deleteWhen=04
fileReservedTime=48


brokerRole=SLAVE
flushDiskType=ASYNC_FLUSH
```
启动broker
```bash
cd /opt/rocketmq/bin
nohup sh mqbroker -c /opt/rocketmq/conf/2m-2s-sync/broker-a-s.properties &
```

# broker-b-master
```bash
vim /opt/rocketmq/conf/2m-2s-sync/broker-b.properties

brokerClusterName=DefaultCluster
namesrvAddr=192.168.10.101:9876;192.168.10.102:9876
brokerName=broker-b
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
启动broker
```bash
cd /opt/rocketmq/bin
nohup sh mqbroker -c /opt/rocketmq/conf/2m-2s-sync/broker-b.properties &
```


# broker-b-slave
```bash
vim /opt/rocketmq/conf/2m-2s-sync/broker-b-s.properties

brokerClusterName=DefaultCluster
namesrvAddr=192.168.10.101:9876;192.168.10.102:9876
brokerName=broker-b
brokerId=1


#是否允许 Broker 自动创建Topic，建议线下开启，线上关闭
autoCreateTopicEnable=false
#是否允许 Broker 自动创建订阅组，建议线下开启，线上关闭
autoCreateSubscriptionGroup=false

deleteWhen=04
fileReservedTime=48


brokerRole=SLAVE
flushDiskType=ASYNC_FLUSH
```
启动broker
```bash
cd /opt/rocketmq/bin
nohup sh mqbroker -c /opt/rocketmq/conf/2m-2s-sync/broker-b-s.properties &
```
# 创建topic
> 创建topic这个步骤在线上环境是必须的
```bash
sh mqadmin updateTopic -n 192.168.10.101:9876 -c DefaultCluster -t mytopic
sh mqadmin updateTopic -n 192.168.10.102:9876 -c DefaultCluster -t mytopic
```

# 创建consumer group
> 创建consumer group的步骤在线上环境是必须的
```bash
sh mqadmin updateSubGroup -g mygroup -c DefaultCluster -n 192.168.10.101:9876
sh mqadmin updateSubGroup -g mygroup -c DefaultCluster -n 192.168.10.102:9876
```
# 测试集群的高可用
将某一个broker停掉，只要有一个broker在运行，应用依然可以正常的发送和消费消息。更细节的高可用测试还待测试。。。


# Broker
```bash
#所属集群名字
brokerClusterName=DefaultCluster
#broker名字，注意此处不同的配置文件填写的不一样
brokerName=broker-a
#0 表示 Master，>0 表示 Slave
brokerId=0
#nameServer地址，分号分割
namesrvAddr=192.168.10.101:9876;192.168.10.102:9876
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


