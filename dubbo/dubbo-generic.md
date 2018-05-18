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
