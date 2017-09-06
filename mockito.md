---
layout: post
title: Mockito
date: 2017-08-16 14:30:00
tags:
- docker
categories: Java
description: mockito
---


# 创建Mock对象
```java
//基于接口创建mock对象
IHelloService mockService = mock(IHelloService.class);
//基于类创建mock对象
User mockUser = mock(User.class);
```



