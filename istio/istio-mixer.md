---
layout: post
title: Istio Mixer
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---




Istio中每个请求，每个Envoy会调用两次Mixer:
1. 转发前，调用Mixer，进行前置检查
2. 转发后，调用Mixer，上报日志和监控数据








# Templates
Envoy -> Mixer -> adapter。不同的adapter接收不同类型的input数据来进行处理。一个logging adapter需要一个log数据，一个metric adapter需要一个metric数据。Istio使用Mixer template来描述adapter需要的具体数据。        



