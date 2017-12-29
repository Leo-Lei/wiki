---
layout: post
title: Kubernetes Resources Limit
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---

# 资源限制
当定义一个Pod的时候，可以指定每个容器可以使用多少的cpu和内存。
使用场景:    
1. 可以让scheduler更好的调度Pod到合适的Node节点。
2. 反正某个容器消耗太多资源，影响Node上其他的容器。

可以添加如下的限制:

* spec.containers[].resources.limits.cpu: cpu上限，可以短暂超过，容器也不会被停止
* spec.containers[].resources.limits.memory: 内存上限，不可以超过；如果超过，容器可能会被终止或调度到其他资源充足的机器上
* spec.containers[].resources.requests.cpu :最低CPU
* spec.containers[].resources.requests.memory :最低内存
其中:
* request是Pod要求的最小的资源
* limits是Pod能使用的最大的资源

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

# cpu和内存单位说明

|                      |                                                             |
| -------------------- | ----------------------------------------------------------- |
| 1 kb(Kilobyte)       | 1 kb(Kilobyte) = 1000 bytes = 8000 bits                     |
| 1 Kib(kibibyte)      | 1 Kib(kibibyte) = 2^10 bytes = 1024 bytes = 8192 bits       |
