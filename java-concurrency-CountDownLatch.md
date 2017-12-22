---
layout: post
title: Java Concurrency (10) CountDownLatch
date: 2015-07-19 14:40:00
tags:
- Java
- Concurrency
categories:
- Java
- Concurrency
---

# 1. The API of CountDownLatch
The common API for CountDownLatch is as below.

| method                                        | description                                                           |
| :-------------------------------------------- | :-------------------------------------------------------------------- |
| void await()                                  | Calling this method will block current thread, until the counter is 0.|
| boolean await(long timeout,TimeUnit unit)     |                                                                       |
| void countDown()                              | Reduce the counter.                                                   |
| long getCount()                               | Get the count. This method is often used for debugging.               |

# 2. Sample of CountDownLatch

```java
package com.leo.gemini.javaconcurrency;

import java.util.concurrent.CountDownLatch;

public class CountDownLatchDemo {

    public long performanceTest(int times,Runnable task) throws InterruptedException{

        if(times <= 0) throw new IllegalArgumentException();

        final CountDownLatch startLatch = new CountDownLatch(1);
        final CountDownLatch overLatch = new CountDownLatch(times);

        for (int i = 0;i < times; i++){
            new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        startLatch.await();
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                    }finally {
                        overLatch.countDown();
                    }
                }
            }).start();
        }

        // Do some preparation work
        long start = System.nanoTime();
        startLatch.countDown();
        overLatch.await();
        return System.nanoTime() - start;
    }
}
```
