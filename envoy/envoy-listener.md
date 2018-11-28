---
layout: post
title: Envoy Listener
date: 2017-06-16 14:35:00
tags:
- docker
categories: Java
---

通常在一台机器上只启动一个Envoy。一个Envoy支持配置多个listener。

每个Listener可以配置一些Filter，这些Filter分为两类：
1. Listener filters
2. Network filters

一个Listener一般都会配置Network filter，对应Listener配置中的FilterChains。可以包含多个FilterChain，每个FilterChain可以包含多个Filter。
Listener filter是可选的。



# 参考文献
* [https://www.envoyproxy.io/docs/envoy/latest/configuration/network_filters/network_filters](https://www.envoyproxy.io/docs/envoy/latest/configuration/network_filters/network_filters)
