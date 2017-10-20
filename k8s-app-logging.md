---
layout: post
title: Kubernetes Process Application Logs
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---


# Kubernetes集群中的日志收集解决方案

|       |            方案                    |                优点               |                  缺点                 |
| ----- | --------------------------------- | --------------------------------- | ------------------------------------ |
|  1    | 每个app的镜像中都集成日志收集组件      | 部署方便，kubernetes的yaml文件无须特别配置，可以为每个app自定义日志收集配置 | 强耦合，不方便应用和日志收集组件升级和维护且会导致镜像过大 |  
|  2    | 单独创建一个日志收集组件跟app的容器一起运行在同一个pod中      | 低耦合，扩展性强，方便维护和升级 |  需要对kubernetes的yaml文件进行单独配置，略显繁琐   |
|  3    | 将所有的Pod的日志都挂载到宿主机上，每台主机上单独起一个日志收集Pod | 完全解耦，性能最高，管理起来最方便    |   需要统一日志收集规则，目录和输出方式   |


# Kubernetes处理容器内应用日志

* 每个Pod中，容器中的java应用通过log4j或logback将日志输出到目录`/opt/logs`
* 每个Pod中，配置Volume，将宿主机器的`/opt/logs`目录挂载到容器的目录`/opt/logs`,这样容器中的java应用在写日志的时候，会写入到宿主机器的`/opt/logs`目录
* 每个宿主机器，即每个k8s的Node上，运行一个Filebeat容器，用来收集宿主机器上的`/opt/logs`目录的日志。并发送到ElasticSearch中。
* Kubernetes中有DaemonSet，可以在每个Node上运行一个Filebeat的Pod。
* Elastic Search和Kibana也是以容器的方式，运行在k8s集群中


![k8s-logging-solution](http://ohaq3i4w3.bkt.clouddn.com/k8s-logging.png)

# 为什么是Filebeat，不是logstash
* 运行Logstash会消耗500M内存
* 运行Filebeat只消耗12M内存

