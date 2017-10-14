---
layout: post
title: Kubernetes Namespaces
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---

# Namespaces
Kubernetes在集群启动后，会创建一个default的Namespace。如果不特别指明，用户创建的Pod，RC，Service等都会被系统创建到这个默认的default的Namespace。


