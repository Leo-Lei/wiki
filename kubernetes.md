---
layout: post
title: Kubernetes
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---



# kubernetes命令

|              Command                                               |                   Desc                            |
| ------------------------------------------------------------------ | ------------------------------------------------- |
| `kubectl get nodes`                                                | 显示所有nodes                                      |
| `kubectl version`                                                  | 显示版本                                           |
| `kubectl get deployments`                                          | 显示所有的deployments                              |
| `kubectl get pods`                                                 | 获取pod                                            |
| `kubectl get pods -o wide`                                         | 获取pod                                            |
| `kubectl describe pods`                                            | 查看pod详情                                         |
| `kubectl run my-app --image=my-app:v1 --port=8080`                 | 运行容器                                            |
| `kubectl exec my_pod my_container env`                             | 在容器中执行命令                                     |
| `kubectl exec -it my_pod my_container bash`                        | 在容器中开始一个bash session                         |
| `kubectl get pods -l 'env in (production,qa),tier in (frontend)'`  | 根据Label选择Pod                                    |
| `kubectl delete -f /opt/pod.json`                                  | 删除kubernetes对象                                  |

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


# Pod

> A Pod is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), IP address and information about how to run them.











# Reference
[kubernetes-interactive-tutorials/kubernetes-basics/deploy-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/deploy-intro/)
