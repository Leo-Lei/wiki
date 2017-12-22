---
layout: post
title: Kubernetes Network
date: 2017-07-08 11:20:00
tags:
- docker
categories: Java
---

# Kubernetes里的三种IP
* Node IP：Node节点的IP地址
* Pod IP：Pod的IP地址
* Cluster IP：Service的IP地址

### Node IP
Node IP是kubernetes中每个节点的物理网卡的IP地址，是一个真实存在的网络。kubernetes之外的节点访问kubernetes集群内的某个节点或者TCP IP服务时，必须通过Node Ip进行通信。

### Pod IP
Pod IP是Docker Engine根据docker0网桥的IP地址段进行分配的，通常是一个虚拟的二层网络，比如Flanneld来创建的网络。Kubernetes里Pod里的容器访问另外一个Pod容器，
都是通过Pod所在的虚拟二层网络进行通信的。

### Service IP
也是一个虚拟IP，但更像一个伪造的IP网络    
* Cluster IP无法被ping，因为没有一个实体网络对象来响应





