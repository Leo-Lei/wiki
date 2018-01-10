---
layout: post
title: Spring Handler Interceptor
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
---

# 执行顺序

假设添加了5个HandlerInterceptor，1，2，3，4，5.

# 正常流程

* 1.preHandle
* 2.preHandle
* 3.preHandle
* 4.preHandle
* 5.preHandle
* real method
* 5.postHandle
* 4.postHandle
* 3.postHandle
* 2.postHandle
* 1.postHandle
* 5.afterCompletion
* 4.afterCompletion
* 3.afterCompletion
* 2.afterCompletion
* 1.afterCompletion


# 中断流程

* 1.preHandle
* 2.preHandle
* 3.preHandle   返回 false，中断流程,下一次执行2.afterCompletion
* 4.preHandle
* 5.preHandle
* real method
* 5.postHandle
* 4.postHandle
* 3.postHandle
* 2.postHandle
* 1.postHandle
* 5.afterCompletion
* 4.afterCompletion
* 3.afterCompletion
* 2.afterCompletion
* 1.afterCompletion








