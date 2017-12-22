---
layout: post
title: Kubernetes Replication Controller
date: 2017-07-08 11:20:00
tags:
- docker
categories: Java
---

# Replication Controller

Replication Controller(简称RC)是定义了一个期望的场景，声明某个Pod的副本数量在任意时刻都符合某个预期值。RC的定义包括如下部分：
* 用于筛选目标Pod的Label Selector
* Pod期待的副本数
* 当Pod副本数量小于预期数量时，用于创建新Pod的Pod模板

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: frontend
spec:
  replicas: 1
  selector:
    tier: frontend
  template:
    metadata:
      labels:
        app: app-demo
        tier: frontend
    spec:
      containers:
      - name: tomcat-demo
        image: tomcat
        imagePullPolicy: IfNotPresent
        env:
        - name: GET_HOSTS_FROM
          value: dns
        ports:
        - containerPort: 80
```


# RC实现Pod动态扩容

```bash
$ kubectl scale rc redis-slave --replicas=3
scaled
```

# RC实现滚动升级
通过改变RC里Pod模板中的镜像版本，实现Pod的滚动升级

# Replication Controller的替换物
Replication Controller是老版本的kubernetes使用的，推荐使用更强大的：
* Replica Set       支持label selector的集合匹配语法
* Deployment        可以查看Pod的部署进度,内部会创建Replica Set


