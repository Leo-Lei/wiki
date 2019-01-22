---
layout: post
title: Java Lombok
date: 2018-06-30 12:20:00
tags:
- Java
categories: Java
---

```java
@Builder
@Data
public class MyClass{

    @Singular(value = "addParam")
    private List<String> params;
}
```

```java

MyClass.builder()
.addParam("hello")
.addParam("world")
.build();

MyClass.builder()
.params(Arrays.asList("hello","world"))
.build();

```
