---
layout: post
title: Java Functional Programing
date: 2017-04-14 15:40:00
tags:
- Java
categories: Java
description: web-authentication
---

# FunctionalInterface
如果一个接口有且仅有一个接口方法，那么这个接口就是函数式接口。JVM会自动认为这样的方法是函数式接口，或者我们也可以显式的指定该接口是函数式接口。
```java
@FunctionalInterface
public interface Function<T, R> {
    R apply(T t);
}
```
如果一个函数式接口作为方法的参数，那么，在调用方法，给方法传参时，可以穿入一个lambda表达式。
