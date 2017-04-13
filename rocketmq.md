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
1. Name Server集群: 提供topic的路由信息,即topic的消息保存在哪几个broker的哪些queue中。
2. Producer集群:拥有相同的ProducerGroup。一般情况下，Producer不必要有集群。
3. Broker集群:一个集群有相同的BrokerClusterName,默认是DefaultCluster。一个集群下有多个master，每个master有多个slave。master和slave是一组，有相同的brokerName，不同的brokerId。master的brokerId是0，slave的brokerId大于0.master和slave之间进行同步复制或异步复制。
4. Consumer集群:拥有相同的ConsumerGroupName。






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
