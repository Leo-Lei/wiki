---
layout: post
title: Kibana
date: 2016-11-12 14:05:00
tags:
- Java
categories: Java
---

# kafka是什么 

* 消息中间件
* 分布式的
* 依赖zookeeper来meta信息，保证系统的可用性

![hh](http://ohaq3i4w3.bkt.clouddn.com/kafka001.jpg)

# kafka相关概念

|      name       |                                   Desc                                 |
| --------------- | ---------------------------------------------------------------------- |
| producer        | 消息生产者，发布消息到kafka。                                              |
| broker          | kafka集群中的一个server就是一个broker                                     |
| topic           | 每条发布到kafka的消息属于的类别                                             |
| partition       | partition是物理上的概念。每个topic包含一个或多个partition。                  |
| consumer        | 从kafka中消费消息。                                                       |
| consumer group  | 每个consumer属于一个consumer group。每条消息只能被consumer group中的一个Consumer消费，但能被多个consumer group消费|
| replica         | partition的副本                                                          |
| leader          | replica中的一个角色，某一个partition的所有的repica中，有且仅有一个leader,producer和consumer只跟leader打交道  |
| follower        | replica中的一个角色，从leader中同步数据                                      |
| controller      | kafka集群中的某一个服务器，用来进行leader selection以及failover               |
| zookeeper       | kafka通过zookeeper存储集群的meta信息                                        |

# kfaka中的partition
一个topic分成多个partition。partition是物理上的概念。topic中的每一条新消息，kafka会根据一定的算法（比如轮训，hash等）来计算消息应该属于哪个partition。一条message属于且只能属于一个partition。但是一个partition可以有几个备份，以提高可用性。

# producer发布消息
