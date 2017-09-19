---
title: Apollo
date: 2017-06-09 10:22:23
categories:
- Music
tags:
- Music
---

# Apollo
[Apollo](https://github.com/ctripcorp/apollo)是携程开发的分布式配置管理平台，集中管理多个环境，不同集群的配置，配置修改后能够实时推送到应用端。

# 架构模块
![apollo-overall-architecture.png](http://ohaq3i4w3.bkt.clouddn.com/apollo-overall-architecture.png)






|   network       |          ip:port            |            jar              |        service              |
| --------------- | --------------------------- | --------------------------- | --------------------------- |
| local           | 192.168.4.100:8080          | config-service.jar          | Config Service: 8080        | 
| local           | 192.168.4.100:8090          | admin-service.jar           | Admin Service: 8090         | 
| aliyun vpc      | 172.168.128.100:8080        | config-service.jar          | Config Service: 8080        | 
| aliyun vpc      | 172.168.128.100:8090        | admin-service.jar           | Admin Service: 8090         | 
| local           | 192.168.4.101:8070          | portal.jar                  | Portal Service: 8070        | 


|                |                         url                                           |   user    |   password   |
| -------------- | --------------------------------------------------------------------- | --------- | ------------ |  
| local          | jdbc:mysql://localhost:3306/ApolloConfigDB?characterEncoding=utf8     |           |              |
| vpc            | jdbc:mysql://aliyun.xxx:3306/ApolloConfigDB?characterEncoding=utf8    |           |              |
| local          | jdbc:mysql://localhost:3306/ApolloPortalDB?characterEncoding=utf8     |           |              |
