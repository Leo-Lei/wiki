---
layout: post
title: Kubernetes Pod
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---

# Pod

> A Pod is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), IP address and information about how to run them.



# Pods overview
![Pod](https://d33wubrfki0l68.cloudfront.net/fe03f68d8ede9815184852ca2a4fd30325e5d15a/98064/docs/tutorials/kubernetes-basics/public/images/module_03_pods.svg)



```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myweb
  labels:
    name: myweb
spec:
  containers:
  - name: myweb
    image: kuberguide/tomcat-app:v1
    ports:
    - containerPort: 8080
    env:
    - name: MYSQL_SERVICE_HOST
      value: 'mysql'
    - nmae: MYSQL_SERVICE_PORT
      value: '3306' 
```




# Pod特性
1. 每个Pod包含一个或多个container
2. 每个Node可以有一个或多个Pod
3. 一个Kubernetes集群中，每个Pod有一个唯一的IP地址，甚至相同Node上的不同Pod


# Reference
[https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/](https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/)
[kubernetes-interactive-tutorials/kubernetes-basics/explore-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/explore-intro/)
