---
layout: post
title: Service Mesh
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
description: What is API and SPI.
---

# 什么是Service Mesh?
* Service Mesh是专用的基础设施层
* 轻量级高性能网络代理
* 提供安全的，快速的，可靠的服务间通讯
* 与实际应用部署一起，但对应用是透明的

# Service Mesh能做什么？
* 提供熔断机制（circuit-breaking）
* 提供感知延迟的负载均衡（latency-awareload balancing）
* 最终一致的服务发现（service discovery)
* 连接重试（retries）及终止（deadlines
* 管理微服务和云原生应用通讯的复杂性，确保可靠地交付应用请求


# Service Mesh是必要的吗？
这可能没有一个绝对的答案，但是:
* Service Mesh可使得快速转向微服务或者云原生应用。
* Service Mesh以一种自然的机制扩展应用负载，解决分布式系统不可避免的部分失败，捕捉高度动态分布式系统的变化。
完全解耦于应用。

# 业界有哪些Service Mesh产品？
* Buoyant的linkerd，基于Twitter的Fingle，长期的实际产线运行经验及验证，支持Kubernetes、DC/OS容器管理平台，也是CNCF官方支持的项目之一。
* Lyft的Envoy，7层代理及通信总线，支持7层HTTP路由、TLS、gRPC、服务发现以及健康监测等。
* IBM、Google、Lyft支持的Istio，一个开源的微服务连接、管理平台以及给微服务提供安全管理，支持Kubernetes、Mesos等容器管理工具，其底层依赖于Envoy。





