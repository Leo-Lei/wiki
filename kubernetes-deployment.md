---
layout: post
title: Kubernetes Deployment
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
banner: http://ohaq3i4w3.bkt.clouddn.com/docker-01.png
---

# Deployment


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


```bash
kubectl create -f docs/user-guide/nginx-deployment.yaml --record
```

```bash
deployment "nginx-deployment" created
```


* apiVersion: 使用的Kubernetes API
* kind: 创建的Object的类型
* metadata: Object的唯一标识，包含name，UID，namespace等


