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

# Label

1. Label就是一些key/value对
2. Label可以附加到Kubernetes的一些对象上，比如Pod
3. 每个Object不能有相同Key的Label
4. 不同的Object可以有相同Key的Label


```json
"release" : "stable", "release" : "canary"
"environment" : "dev", "environment" : "qa", "environment" : "production"
"tier" : "frontend", "tier" : "backend", "tier" : "cache"
"partition" : "customerA", "partition" : "customerB"
"track" : "daily", "track" : "weekly"
```


# Selector
通过Selector来选择Label            
### Equality-based Selector
```properties
environment = production # 有key=environment，且value=production
tier != frontend         # 有key=tier，且value！=frontend         或者没有key=tier的
environment=production,tier!=frontend     # ,隔开的是AND的关系，要同时满足这些条件
```
### Set-based Selector
```properties
environment in (production, qa)     # 有key=environment，且value等于production或qa
tier notin (frontend, backend)      # 有key=tier，且value不等于frontend或backend         或者没有key=tier
partition                           # 有key=partition
!partition                          # 没有key=partition
```

# Reference
[kubernetes-interactive-tutorials/kubernetes-basics/explore-intro/](https://kubernetes.io/docs/tutorials/kubernetes-basics/explore-intro/)


