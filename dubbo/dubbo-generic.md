---
layout: post
title: Dubbo的泛化调用
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---

# 什么是泛化
泛化，简单而言就是consumer在调用provider时，不需要依赖provider提供的接口api的Jar包。

# 为什么需要泛化
想到以下几种场景：

1. Provider的接口和接口方法变化很快，包括新增，修改等。这时候，如果采用传统方式，consumer就需要不断同步更新接口的Jar包。在维护上会比较困难。
2. Consumer不想强依赖Provider的接口Jar。可以降低耦合。
