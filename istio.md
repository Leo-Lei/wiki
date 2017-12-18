---
layout: post
title: Istio
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
description: Service Mesh Istio
---



|     Component     |                           Desc                         |
| ----------------- | ------------------------------------------------------ |
| Pilot             |                                                        |
| Mixer             |                                                        |
| Istio-Auth        |                                                        |
| Envoy             |                                                        |




# Envory
Envoy
> Istio 使用Envoy代理的扩展版本，Envoy是以C++开发的高性能代理，用于调解服务网格中所有服务的所有入站和出站流量。
> 
> Istio利用了Envoy的许多内置功能，例如动态服务发现，负载均衡，TLS termination，HTTP/2&gRPC代理，熔断器，健康检查，基于百分比流量拆分的分段推出，故障注入和丰富的metrics。
> 
> Envoy实现了过滤和路由、服务发现、健康检查，提供了具有弹性的负载均衡。它在安全上支持TLS，在通信方面支持gRPC。




# Pilot
> Pilot负责收集和验证配置并将其传播到各种Istio组件。它从Mixer和Envoy中抽取环境特定的实现细节，为他们提供独立于底层平台的用户服务的抽象表示。此外，流量管理规则（即通用4层规则和7层HTTP/gRPC路由规则）可以在运行时通过Pilot进行编程。
> 
> 每个Envoy实例根据其从Pilot获得的信息以及其负载均衡池中的其他实例的定期健康检查来维护 负载均衡信息，从而允许其在目标实例之间智能分配流量，同时遵循其指定的路由规则。
> 
> Pilot负责在Istio服务网格中部署的Envoy实例的生命周期。


# Mixer
Mixer 提供三个核心功能：
* 前提条件检查。允许服务在响应来自服务消费者的传入请求之前验证一些前提条件。前提条件包括认证，黑白名单，ACL检查等等。
* 配额管理。使服务能够在多个维度上分配和释放配额。典型例子如限速。
* 遥测报告。使服务能够上报日志和监控。

在Istio内，Envoy重度依赖Mixer。










# Reference
* https://doczhcn.gitbooks.io/istio/
* https://istio.io/

