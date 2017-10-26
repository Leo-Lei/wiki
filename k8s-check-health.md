---
layout: post
title: Kubernetes Check Health
date: 2017-07-08 11:20:00
tags:
- docker
categories: Java
description: docker
---


# 容器健康检查

Pod通过LivenessProbe和ReadinessProbe两种探针来检查容器的健康状态：

* LivenessProbe 用于判断容器是否健康，如果LivenessProbe探测到容器不健康，kubelet将删除该容器并根据容器的重启策略做相应的处理。如果容器不包含LivenessProbe，则kubelet认为该容器的LivenessProbe探针永远返回success。
* ReadinessProbe 用于判断容器是否启动完成且准备接收请求。如果该探针探测到失败，则Endpoint Controller将会从Service的Endpoint中删除包含该容器Pod的条目。
Kubelet定期调用容器中的LivenessProbe针来检查容器的健康状态。

# Liveness Probes


# Readiness Probes

