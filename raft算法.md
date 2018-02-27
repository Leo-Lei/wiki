---
layout: post
title: Raft算法
date: 2016-07-14 15:40:00
tags:
- Java
categories: Java
---

# Raft
Raft是分布式系统中的一致性算法。可以保证分布式系统中的数据的一致性。

# Raft使用场景
分布式系统中，消除单点故障。多个节点，有着一致的数据。有且仅有一个节点是Leader，来对外提供服务。当Leader宕机时，重新选举出一个Leader，继续对外提供服务。

# Raft中的角色
* Leader：处理客户端请求，读，写数据
* Follower：不处理客户端请求，只被动同步Leader的数据，即同步日志
* Candidate：当开始选举的时候，Follower会成为Candidate

# Raft中的术语
|         术语         |                   描述                       |
| ------------------- | -------------------------------------------- | 
| Leader              |                                              |
| Follower            |                                              |
| Candidate           |                                              |
| Heartbeat           |                                              |
| election timeout    |                                              |
| heartbeat timeout   |                                              |

# Raft中的选举
简单来说，如果某个节点收到了超过半数的选票，就成为Leader。    
但要想真正理解选举，里面的细节还是很多的。下面是一些关键点：
* Raft开始的时候，所有节点都是Follower
* Raft中有任期(Term)的概念。新的一轮选举开始的时候，Term加1。Term是递增的。
* 每个节点在本地都会保存一份当前Term的值。
* Raft开始后，经过一段时间，某个或某几个节点开始新一轮的选举。发起选举的节点，会把本地的Term加1，投一票给自己，然后通知其他，让他们来选自己。
* 接到投票请求的节点，如果节点本地的Term比投票请求中的小，并且该轮选举中，自己还没有投过票，就将票投给发起投票请求的节点。
* Raft开始后，各节点经过election timeout后开始发起新的一轮选举。每个节点的election timeout是随机的。所以，绝大部分情况下，会有一个节点最先发起新的一轮选举，它会给自己投一票，Term加1，然后通知其他节点给自己投票。在其他节点给它投了票后，如果超过了半数，它就成为了Leader。


# Raft心跳
选出一个Leader之后，Leader会定期的给Follower发送心跳，表示Leader还活着。只要当前Leader还活着，该节点就会一直是Leader。

# Raft算法中为什么推荐节点数是奇数
因为奇数节点与和其配对的偶数个节点相比(比如3 节点和4 节点相比)，容错能力相同，却可以少一个节点。








