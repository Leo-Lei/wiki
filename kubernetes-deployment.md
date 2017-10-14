---
layout: post
title: Kubernetes Deployment
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---

# Deployment

Deployment是在Replication Controller之后引入的概念，可以把它看成是RC的一次升级。两者相识度超过90%。         
### Deployment的一些特点
Deployment内部会创建Replica Set来创建和管理Pod         
### Deployment相对于RC的升级:       
1. 可以看到Pod部署的进度。


# Deployment 使用场景
1. 创建一个Deployment来生成对应的Replica Set并完成Pod副本的创建过程
2. 检查Deployment的状态来看部署动作是否完成。
3. 更新Deployment以创建新的Pod来进行扩容
4. 更新Deploymnet以更新Pod来达到滚动升级
5. 如果当前Deployment不稳定，可以回滚到一个早先的Deployment版本
6. 挂起或恢复一个Deployment

# Deployment定义

Deployment的定义和RC，Replica Set定义很类似，除了API声明与king等有所区别         
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      tier: frontend
    matchExpressions:
      - {key: tier, operator: In, values:[frontend]}
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
          ports:
          - containerPort: 8080

```

```bash
$ kubectl create -f /opt/tomcat-deployment.yaml
deployment "tomcat-deployment" created
```

```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80

```

# 创建Deployment

```bash
$ kubectl create -f docs/user-guide/nginx-deployment.yaml --record
deployment "nginx-deployment" created
```

# 查看Deployment
```bash
$ kubectl get deployments
```
