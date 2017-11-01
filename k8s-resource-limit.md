---
layout: post
title: Kubernetes Resources Limit
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---

# 给Pod添加资源限制
当定义一个Pod的时候，可以指定每个容器可以使用多少的cpu和内存。
使用场景:    
1. 可以让scheduler更好的调度Pod到合适的Node节点。
2. 反正某个容器消耗太多资源，影响Node上其他的容器。

可以添加如下的限制:

* spec.containers[].resources.limits.cpu
* spec.containers[].resources.limits.memory
* spec.containers[].resources.requests.cpu
* spec.containers[].resources.requests.memory



```yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend
spec:
  containers:
  - name: db
    image: mysql
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "password"
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  - name: wp
    image: wordpress
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
```

