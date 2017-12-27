---
layout: post
title: Kubernetes Ingress
date: 2017-07-08 11:25:00
tags:
- docker
categories: Java
---


# 概述
Kubernetes中的Service表现为IP:Port的形式，有以下几种表现形式:
* ClusterIP
* NodePort
* LoadBalance

ClusterIP: 是一个集群内部的IP，仅限于集群内通信。        
NodePort: 可以实现暴露服务的访问入口，但service的NodePort会作用于每个Node节点。        
LoadBalance: 通常需要第三方云提供商支持，比如AWS，Azure，阿里云等。有很大的约束性。

基于HTTP的服务，对于不同的URL，会对应后端不同的Service。需要一个规则，来把URL路由到Kubernetes集群中不同的Service。

# 什么是Ingress
> Ingress在英文中是入站的意思。Outgress是出站。
```text

                                         internet
                    
                    
          foo.com/some/path          foo.com/other/path              bar.com/some/path         
              
              
           [NodePort:8000]            [NodePort:8000]                  [NodePort:8000]
  ------------------------------------------------------------------------------------------------      
            [Service1]                  [Service2]                       [Service3]


Ingress实现集群内所有服务的入口，通过一系列规则来运行外部的访问。

                                         internet
                    
                    
          foo.com/some/path          foo.com/other/path              bar.com/some/path         
              
              
       [                                   Ingress                                            ]            
  ------------------------------------------------------------------------------------------------      
            [Service1]                  [Service2]                       [Service3]

```
在定义Ingress之前，需要先部署Ingress Controller，以实现为所有后端Service提供一个统一的入口。Ingress Controller需要实现基于不同HTTP URL 向后端转发的负载分发规则，通常应根据应用系统的需求进行自定义实现。在kubernetes中，Ingress Controller将以Pod的形式运行，监控Apiserver的/ingress接口后端的backend service，如果service发生变化，则Ingress Controll应自动更新其转发规则。             
Ingress Controller可以有多种实现方式，比如Nginx，traefik等。






