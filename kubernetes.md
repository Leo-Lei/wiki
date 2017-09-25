---
layout: post
title: Kubernetes
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
banner: http://ohaq3i4w3.bkt.clouddn.com/docker-01.png
---

# kubernetes
[kubernetes](https://kubernetes.io/)

# Kubernetes一些概念

|            Term           |  Desc                                             |
| ------------------------- | ------------------------------------------------- |
| kubectl                   | Kubernetes命令行工具                                |
| heapster                  | kubernetes默认监控系统                              |
| kubernetes-dashboard      |                                                   |
| monitoring-grafana        |                                                   |
| monitoring-influxdb       |                                                   |


|           Term            |                                 Desc                                |
| ------------------------- | ------------------------------------------------------------------- |
| `Pod`                     | 代表一个组，包含一个或多个容器，和这些容器共享的资源(volume,IP等)            |

# kubernetes命令

|              Command                                 |                   Desc                              |
| ---------------------------------------------------- | --------------------------------------------------- |
| `kubectl get nodes`                                  | 显示所有nodes                                        |
| `kubectl version`                                    | 显示版本                                             |
| `kubectl get deployments`                            | 显示所有的deployments                                |
| `kubectl run my-app --image=my-app:v1 --port=8080`   | 运行容器                                             |





# Pod

> A Pod is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), IP address and information about how to run them.











# Reference
[kubernetes-interactive-tutorials/kubernetes-basics/deploy-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/deploy-intro/)
