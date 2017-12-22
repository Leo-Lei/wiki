---
layout: post
title: Caffeine
date: 2017-07-05 17:10:00
tags:
- Java
categories: Java
---

# Overview               
一个Java的内存数据库


```java
Cache<String,String> cache = Caffeine.newBuilder()
                .maximumSize(3)
                .build();
        cache.put("A","100");
        cache.put("B","200");
        cache.put("C","300");
        cache.put("D","400");
        cache.put("E","500");
        cache.put("F","600");
        cache.put("G","700");
        cache.put("H","800");
        cache.put("I","900");

        // 当超过maximumSize时，Caffeine会执行清理工作，但是异步的，所以不是严格的保证任何时候的size都不超过maximumSize
        // Caffeine可以保证当数量超过maximumSize时，Caffeine可以尽快的执行清理工作
        // 所以这里让线程sleep3秒钟，让清理工作完成
        Thread.sleep(3 * 1000);     
        
        System.out.println(cache.estimatedSize());
```

