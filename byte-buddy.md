---
layout: post
title: Byte-Buddy
date: 2015-06-30 15:50:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
---

# 1. Byte-Buddy
[byte-buddy](https://github.com/raphw/byte-buddy)





# 创建一个接口
```java
public interface HelloService {
    String hello(String name);

    String hello(Person person);
}
```
```java
HelloService hello = new ByteBuddy()
        .subclass(HelloService.class)
        .method(isDeclaredBy(HelloService.class)).intercept(MethodDelegation.to(GeneralInterceptor.class))
        .make()
        .load(getClass().getClassLoader())
        .getLoaded()
        .newInstance();
hello.hello("world");
hello.hello(new Person("foo","bar"));
```

```java
package com.leibangzhu.iris.test.testbytebuddy;

import net.bytebuddy.implementation.bind.annotation.AllArguments;
import net.bytebuddy.implementation.bind.annotation.Origin;
import net.bytebuddy.implementation.bind.annotation.RuntimeType;

import java.lang.reflect.Method;

public class GeneralInterceptor {

    @RuntimeType
    public static String intercept(@AllArguments Object[] allArguments, @Origin Method method){

        String name = method.getDeclaringClass().getName();
        return "";
    }
}

```
