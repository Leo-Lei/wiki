---
layout: post
title: Kubernetes Labels and Selectors
date: 2017-07-09 11:10:00
tags:
- docker
categories: Java
description: docker
banner: http://ohaq3i4w3.bkt.clouddn.com/docker-01.png
---

# Labels and Selectors

> A Pod is a group of one or more application containers (such as Docker or rkt) and includes shared storage (volumes), IP address and information about how to run them.


# Label

1. Label就是一些key/value对
2. Label可以附加到Kubernetes的一些对象上，比如Pod


```json
"release" : "stable", "release" : "canary"
"environment" : "dev", "environment" : "qa", "environment" : "production"
"tier" : "frontend", "tier" : "backend", "tier" : "cache"
"partition" : "customerA", "partition" : "customerB"
"track" : "daily", "track" : "weekly"
```








# Reference
[kubernetes-interactive-tutorials/kubernetes-basics/explore-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/explore-intro/)


