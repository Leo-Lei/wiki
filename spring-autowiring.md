---
layout: post
title: Spring Autowiring
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---



|      annotation     |                            desc                                                 |
| ------------------- | ------------------------------------------------------------------------------- |
| `@Primary`          | 在@Autowired的时候，如果一个type有多个实现，可以通过@Primary，告诉Spring优先选择该组件。  |
| `@Qualifier`        | 当@Autowired时，如果有多个bean，通过@Qualifier指定bean的id来注入                      |


# Autowire List,Set,Map
```java
interface StringCallable extends Callable<String> { }
 
@Component
class Third implements StringCallable {
    @Override
    public String call() {
        return "3";
    }
 
}
 
@Component
class Forth implements StringCallable {
    @Override
    public String call() {
        return "4";
    }
 
}
 
@Component
class Fifth implements StringCallable {
    @Override
    public String call() throws Exception {
        return "5";
    }
}
```



```java
@Component
class BootstrapTest {
 
    @Autowired
    List<StringCallable> list;     // 3,4,5
 
    @Autowired
    Set<StringCallable> set;      //3,4,5
 
    @Autowired
    Map<String, StringCallable> map;   // 3,4,5   key是bean name,即third,forth,five
 

```


