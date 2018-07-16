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

