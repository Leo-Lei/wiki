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
1. Name Server集群: 就是注册中心。RocketMQ在注册中心没有使用第三方中间价，而是自己写代码实现的，代码行数才1000行。Producer，Consumer，Broker在启动时都需要向NameServer进行注册，NameServer之间不通讯。互不可见。
2. Producer集群:拥有相同的ProducerGroup。一般情况下，Producer不必要有集群。
3. Broker集群:一个集群有相同的BrokerClusterName,默认是DefaultCluster。一个集群下有多个master，每个master有多个slave。master和slave是一组，有相同的brokerName，不同的brokerId。master的brokerId是0，slave的brokerId大于0.master和slave之间进行同步复制或异步复制。
4. Consumer集群:拥有相同的ConsumerGroupName。





# RocketMQ组件间通信关系
* Producer和Name Server：每一个Producer会和Name Server集群中的某一台机器建立TCP连接，会从这台NameServer上拉取路由信息。
* Producer和Broker：Producer和它要发送的topic相关的master类型的broker建立TCP连接。用于发送消息和心跳。Broker中会记录该Producer的信息，供查询使用。
* Broker和Name Server：Broker(不管是master还是slave)会和每一台Name Server建立TCP连接。Broker在启动的时候会注册自己配置的topic信息到Name Server集群的每一台机器中。即每一台Name Server都有该broker的topic配置信息。
* Consumer和Name Server：每一个Consumer会和Name Server集群中的某一台机器建立TCP连接，会从这台Name Server上拉取路由信息，进行负载均衡。
* Consumer和Broker：Consumer可以和master或者slave的broker建立TCP连接来进行消费消息。Consumer也会向它所消费的Broker发送心跳信息，供Broker记录。



# Name Server
和kafka，hadoop中的zookeeper的角色类似。提供topic的路由信息，路由信息存储在内存中。Broker会定时发送路由信息到NameServer集群的所有机器。NameServer是无状态和相互独立的。
```bash
[root@ip-172-31-19-214 bin] bash mqadmin topicRoute -n 172.31.19.214:9876 -t qibei_user_invite
{
	"brokerDatas":[
		{
			"brokerAddrs":{0:"172.31.10.10:10911"
			},
			"brokerName":"ip-172-31-10-10"
		}
	],
	"filterServerTable":{},
	"queueDatas":[
		{
			"brokerName":"ip-172-31-10-10",
			"perm":6,
			"readQueueNums":4,
			"topicSynFlag":0,
			"writeQueueNums":4
		}
	]
}
```

# Broker


# Producer


# Consumer



# 消息的存储
Topic是一类消息的统称，为了提高消息的写入和读取并发能力，将一个topic的消息进行拆分，可以分散到多个broker中，kafka上称为分区，RockertMQ称为队列。
