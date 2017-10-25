---
layout: post
title: Kubernetes Lifecycle
date: 2017-07-08 11:25:00
tags:
- docker
categories: Java
description: docker
---


# Lifecycle

|     Event      |  Desc                    |
| -------------- | ------------------------ |
| `postStart`    | 容器创建之后               |
| `preStop`      | 容器停止之前               |

# lifecycle的使用场景
1. 容器优雅停止

# Demo
定义了一个容器，会监听Pod的一些事件，然后执行相应的命令        
***lifecycle-events.yaml:***
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: lifecycle-demo
spec:
  containers:
  - name: lifecycle-demo-container
    image: nginx
    lifecycle:
      postStart:
        exec:
          command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
      preStop:
        exec:
          command: ["/usr/sbin/nginx","-s","quit"]
```

