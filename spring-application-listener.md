---
layout: post
title: Spring Application Listener
date: 2017-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---

# ApplicationListener定义
```java
@FunctionalInterface
public interface ApplicationListener<E extends ApplicationEvent> extends EventListener {
    void onApplicationEvent(E var1);
}
```
