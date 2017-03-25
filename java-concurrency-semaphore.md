---
layout: post
title: Java Semaphore
date: 2015-08-11 17:15:00
tags:
- Java
categories: Java
description: The tutoria will describe the useage of wait and notify.
---

# 1.Some sample of Java Semaphore

```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Semaphore;

public class SemaphoreTest {
    public static void main(String[] args) {
        // 线程池
        ExecutorService exec = Executors.newCachedThreadPool();
        // 只能5个线程同时访问
        final Semaphore semp = new Semaphore(5);
        // 模拟20个客户端访问
        for (int index = 0; index < 50; index++) {
            final int NO = index;
            Runnable run = new Runnable() {
                public void run() {
                    try {
                        // 获取许可
                        semp.acquire();
                        System.out.println("Accessing: " + NO);
                        Thread.sleep((long) (Math.random() * 10000));
                        // 访问完后，释放
                        semp.release();
                        //availablePermits()指的是当前信号灯库中有多少个可以被使用
                        System.out.println("-----------------" + semp.availablePermits());
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            };
            exec.execute(run);
        }
        // 退出线程池
        exec.shutdown();
    }
}

```
