---
layout: post
title: Kubernetes Node
date: 2017-07-08 11:20:00
tags:
- docker
categories: Java
---

# Node

> A node is a worker machine in Kubernetes and may be a VM or physical machine, depending on the cluster. Multiple Pods can run on one Node.


1. Node是Kubernetes中的一个worker machine
2. Node可以是一个物理机，也可以是一个虚拟机，或一个云服务商的机器
3. Node被Master管理
4. Node上运行Pod。

# Node overview
![Node](https://d33wubrfki0l68.cloudfront.net/5cb72d407cbe2755e581b6de757e0d81760d5b86/a9df9/docs/tutorials/kubernetes-basics/public/images/module_03_nodes.svg)

# Address
* HostName: 机器的hostname
* ExternalIP: 集群外部可访问的IP
* InternalIP: 只能集群内访问的IP



# 创建Node
```json
{
  "kind": "Node",
  "apiVersion": "v1",
  "metadata": {
    "name": "10.240.79.157",
    "labels": {
      "name": "my-first-k8s-node"
    }
  }
}
```






# Reference
[https://kubernetes.io/docs/concepts/architecture/nodes/](https://kubernetes.io/docs/concepts/architecture/nodes/)        
[kubernetes-interactive-tutorials/kubernetes-basics/explore-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/explore-intro/)
