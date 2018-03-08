---
layout: post
title: CAS
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---

# CAS
CAS全称是Compare And Swap。如果目前的值与预期值相匹配，那么处理器就会将该值更新为新值。否则处理器不做任何处理。CAS是一个原子操作。

> AtomicInteger, AtomicBoolean, AtomicLong等类就是用CAS实现的，在性能上比Synchronized更高。

