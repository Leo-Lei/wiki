---
layout: post
title: Kubernetes Process Application Logs
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---



# Kubernetes处理容器内应用日志

* 每个Pod中，容器中的java应用通过log4j或logback将日志输出到目录`/opt/logs`
* 每个Pod中，配置Volume，将宿主机器的`/opt/logs`目录挂载到容器的目录`/opt/logs`,这样容器中的java应用在写日志的时候，会写入到宿主机器的`/opt/logs`目录
* 每个宿主机器，即每个k8s的Node上，运行一个Filebeat容器，用来收集宿主机器上的`/opt/logs`目录的日志。并发送到ElasticSearch中。
* Kubernetes中有DaemonSet，可以在每个Node上运行一个Filebeat的Pod。
* Elastic Search和Kibana也是以容器的方式，运行在k8s集群中


