---
layout: post
title: Kubernetes Service
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---

# Service

> A Kubernetes Service is an abstraction layer which defines a logical set of Pods and enables external traffic exposure, load balancing and service discovery for those Pods.

![Kubernetes-service](https://d33wubrfki0l68.cloudfront.net/cc38b0f3c0fd94e66495e3a4198f2096cdecd3d5/ace10/docs/tutorials/kubernetes-basics/public/images/module_04_services.svg)

![Kubernetes-label](https://d33wubrfki0l68.cloudfront.net/b964c59cdc1979dd4e1904c25f43745564ef6bee/f3351/docs/tutorials/kubernetes-basics/public/images/module_04_labels.svg)



# Service特点
1. 每个Service都有一个唯一的Cluster IP，以及唯一的名字。IP是kubernetes自动分配的，名字是开发者自己定义的。
2. Kubernetes提供了DNS插件，安装好DNS插件后，可以通过Service的名字找到Service的Cluster IP。

# 定义Service
### 一个端口的Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: tomcat-service
spec:
  ports:
  - port: 8080
  selector:
    tier: frontend
```
其中：
* ports中表示，service将哪个端口的请求转发到Pod中的哪个端口。需要定义2个端口，port和targetPort，如果没有指定，默认targetPort和port相同。

### 多个端口的Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: tomcat-service
spec:
  ports:
  - port: 8080
    name: service-port
  - port: 8005
    name: shutdown-port
  selector:
    tier: frontend
```
注意：
* 多端口的服务中，给每个端口命名，是为了Kubernetes的服务发现机制。


# 创建Service
```bash
kubectl create -f /some/path/some-service.yaml
```

# 查看Service
```bash
$ kubectl get svc tomcat-service -o yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: 2016-07-21T17:05:52z
  name: tomcat-service
  namespace: default
  resourceVersion: "23964"
  selfLink: /api/v1/namespaces/default/services/tomcat-service
  uid: 61987d3c- 4f65- 11e6- a9d8- 000c29ed42c1
spec:
  clusterIP: 169.169.65.227
  ports:
  - port: 8080
    portocol: TCP
    targetPort: 8080
  selector:
    tier: frontend
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer:{}
```
其中：
* targetPort：将Service的请求，转发到Pod的对应端口上。提供该服务的容器需要暴露该接口。



# 外部系统访问Service
Cluster IP是属于Kubernetes集群内部的地址，无法再集群外部直接使用这个地址。有些时候，Kubernetes中的服务是要暴露给Kubernetes集群外部的应用或用户来访问的，比如web服务。有以下几种方式来访问Service：    
1. 使用Node iP + NodePort
2. 使用Load Balancer


### Node IP + Node Port
指定Service的spec.type为NodePort，手动指定nodePort的值。
```yaml
apiVersion: v1
kind: Service
metadata:
  name: tomcat-service
spec:
  type: NodePort
  ports:
  - port: 8080
    nodePort: 31002
  selector:
    tier: frontend
```
然后通过<NodeIP>:31002来访问tomcat服务。    
NodePort的实现


# Reference
[kubernetes-interactive-tutorials/kubernetes-basics/expose-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/explore-intro/)
