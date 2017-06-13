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

### Name Server
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
