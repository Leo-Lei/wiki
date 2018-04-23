---
layout: post
title: Dubbo的SPI扩展机制
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---


Dubbo的扩展机制




1. 可扩展的几种解决方案
    工厂
    Java SPI
    Spring 容器（依赖Spring）
    其他IoC容器（依赖第三方库
2. Java SPI机制
    对SPI的概念有个了解。可以把API和SPI做个对比。
3. Java的SPI机制的缺陷
4. dubbo SPI机制 
    Dubbo SPI机制对Java SPI的优化
5. Dubbo Extension Loader
    ExtentionLoader源码解读
6. Dubbo的LoadBalance扩展点解读
    Dubbo中的LoadBalance也是一个SPI，结合源码，分析LoadBalance是如何被加载的
7. 自定义一个LoadBalance扩展
    演示如何自己实现一个LoadBbalance，在不改变dubbo源码的情况下，让Dubbo使用我们自定义的LoadBalance实现
8. Dubbo SPI高级用法之IoC
   AdaptiveInstance
9. Dubbo SPI高级用法之AoP
   wrapper
10. Dubbo SPI核心剥离：https://github.com/alibaba/cooma

