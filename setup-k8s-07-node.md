---
layout: post
title: Setup Kubernetes Cluster - Node
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---

# 部署Node节点
kubernetes master 节点包含的组件：
* flanneld
* docker
* kubelet
* kube-proxy


# 变量
```bash
$ # 替换为 kubernetes master 集群任一机器 IP
$ export MASTER_IP=10.64.3.7
$ export KUBE_APISERVER="https://${MASTER_IP}:6443"
$ # 当前部署的节点 IP
$ export NODE_IP=10.64.3.7
$ # 导入用到的其它全局变量：ETCD_ENDPOINTS、FLANNEL_ETCD_PREFIX、CLUSTER_CIDR、CLUSTER_DNS_SVC_IP、CLUSTER_DNS_DOMAIN、SERVICE_CIDR
$ source /root/local/bin/environment.sh
$
```
