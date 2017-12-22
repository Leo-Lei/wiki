---
layout: post
title: Java Concurrency (3) Atomic operation part 2
date: 2015-06-30 19:50:00
tags:
- Java
- Concurrency
categories:
- Java
- Concurrency
---

# 1. AtomicIntegerArray/AtomicLongArray

The common API for AtomicIntegerArray is as below. The API is very similar to AtomicInteger.
The name of method and argument are self-explanatory, so no more description is required.

| method                                            | description    |
| :------------------------------------------------ | :------------- |
| int get(int i)                                    |                |
| void set(int i,int newValue)                      |                |
| void getAndSet(int i,int newValue)                |                |
| void getAndIncrement(int i,int newValue)          |                |
| void getAndDecrement(int i,int newValue)          |                |
