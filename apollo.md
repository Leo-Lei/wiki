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

Apollo的整体架构包含：
* Config Service
* Admin Service
* Meta Server
* Eureka
* Portal
* Client

### Config Service
* 提供读配置的接口
* 接口服务对象为Apollo Client
* 每个环境都需要部署一个Config Service
* 每个环境要根据自己的配置来打包一个config-service.jar,并部署
* 会读ApolloConfigDB数据库

### Admin Service
* 提供读/写配置的接口
* 接口服务对象为Portal
* 会读/写ApolloConfigDB
* 每个环境都需要部署一个Admin Service
* 每个环境都要根据自己的配置来打包一个admin-service.jar，并部署

### Meta Server
* Portal访问Meta Server来获取Admin Service服务列表
* Client访问Meta Server来获取Config Service服务列表
* Meta Server从Eureka获取Config Service和Admin Service的服务信息，相当于是一个Eureka Client
* 增加一个Meta Server的角色是为了封装服务端发现的细节，对Portal和Client而言，永远通过一个Http接口获取Admin Service和Config Service的服务信息，而不需要关心背后实际的服务注册和发现组件
* Meta Server是一个逻辑角色，在部署时是和Config Service在一个JVM中的
* config-service.jar中已包含了Meta Server

### Eureka
* 基于Eureka和Spring Cloud Netflix提供服务注册和发现
* Config Service和Admin Service会向Eureka注册服务，并保持心跳
* 目前Eureka和Config Service在同一个JVM中。
* config-service.jar中已包含了Eureka

### Portal
* 提供WEB界面供用户管理配置
* 通过Meta Service获取Admin Service服务列表通过IP+Port访问服务
* 在Portal侧做load balance，fail over
* 只需要打包一个porta.jar，只需要在某一个环节部署一个portal.jar。

### Client
* 客户端程序，可能是JAVA的一个jar包，.NET的一个DLL文件
* 通过Meta Server获取Config Service服务列表，通过IP+Port访问服务
* 在Client侧做load balance，fail over
* 只需要打包一个client的jar，并且需要把这个jar上传到Maven私服中


# 集群配置
如下是Apollo集群的部署情况：
|    env    |  private ip:port      |  public ip:port     |          jar          |
| --------- | --------------------- | ------------------- | --------------------- |
| dev       | 192.168.4.100:8080    |                     | config-service.jar    |
| dev       | 192.168.4.100:8090    |                     | admin-service.jar     |
| pro       | 172.16.128.100:8080   | 11.22.33.44:1111    | config-service.jar    |
| pro       | 172.16.128.100:8090   | 11.22.33.44:2222    | admin-service.jar     |
| dev       | 192.168.4.200:8070    |                     | portal.jar            |























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
