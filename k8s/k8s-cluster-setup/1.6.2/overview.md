---
layout: post
title: Setup Kubernetes Cluster
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---

# 集群成员

|      IP         |           Role          |      OS        |
| --------------- | ----------------------- | -------------- |
| 192.168.5.100   | k8s-Master,etcd-0       | CentOS 7       |
| 192.168.5.101   | k8s-Node,  etcd-1       | CentOS 7       |
| 192.168.5.102   | k8s-Node,  etcd-2       | CentOS 7       |
| 192.168.5.103   | Docker Registry         | CentOS 7       |

# 集群组件和版本

|    Component     |       Version     |
| ---------------- | ----------------- |
| Kubernetes       | 1.6.2             |
| Docker           | 17.04.0-ce        |
| Etcd             | 3.1.6             |
| Flanneld         | 0.7.1             |

# 准备工作
### 关闭防火墙
```bash
systemctl disable firewalld
systemctl stop firewalld
```
### 使用全局变量

后续的部署过程中，会使用下面定义的全局环境变量，创建一个environment.sh文件，并将它拷贝到所有的Master和Node机器上。    
environment.sh
```bash
#!/usr/bin/bash

# TLS Bootstrapping 使用的 Token，可以使用命令 head -c 16 /dev/urandom | od -An -t x | tr -d ' ' 生成
BOOTSTRAP_TOKEN="41f7e4ba8b7be874fcff18bf5cf41a7c"

# 最好使用 主机未用的网段 来定义服务网段和 Pod 网段

# 服务网段 (Service CIDR），部署前路由不可达，部署后集群内使用IP:Port可达
SERVICE_CIDR="10.254.0.0/16"

# POD 网段 (Cluster CIDR），部署前路由不可达，**部署后**路由可达(flanneld保证)
CLUSTER_CIDR="172.30.0.0/16"

# 服务端口范围 (NodePort Range)
export NODE_PORT_RANGE="8400-9000"

# etcd 集群服务地址列表
export ETCD_ENDPOINTS="https://192.168.5.100:2379,https://192.168.5.101:2379,https://192.168.5.102:2379"

# flanneld 网络配置前缀
export FLANNEL_ETCD_PREFIX="/kubernetes/network"

# kubernetes 服务 IP (一般是 SERVICE_CIDR 中第一个IP)
export CLUSTER_KUBERNETES_SVC_IP="10.254.0.1"

# 集群 DNS 服务 IP (从 SERVICE_CIDR 中预分配)
export CLUSTER_DNS_SVC_IP="10.254.0.2"

# 集群 DNS 域名
export CLUSTER_DNS_DOMAIN="cluster.local."
```
