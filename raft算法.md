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

# Raft中的选举

