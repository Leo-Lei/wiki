---
layout: post
title: Java ThreadLocal
date: 2015-08-18 16:40:00
tags:
- Java
categories: Java
---


# 1. Thread Local Example

```java

public class ThreadLocalSample {

    @Test
    public void Test() throws InterruptedException {

        MyRunnable work = new MyRunnable();

        Thread thread1 = new Thread(work);
        Thread thread2 = new Thread(work);
        Thread thread3 = new Thread(work);

        thread1.start();
        thread2.start();
        thread3.start();

        Thread.sleep(1000*10);
    }
}

class MyRunnable implements Runnable{

    private Integer local = (int) (Math.random() * 100);
    private ThreadLocal<Integer> threadLocal = new ThreadLocal<>();

    @Override
    public void run() {
        threadLocal.set((int) (Math.random() * 100));

        try {
            Thread.sleep(2000);
        }catch (Exception e){

        }

        System.out.println("Local:" + local);
        System.out.println("ThreadLocal:" + threadLocal.get());
    }
}

// Local:60
// Local:60
// ThreadLocal:82
// Local:60
// ThreadLocal:45
// ThreadLocal:99

```

**Initial ThreadLocal Value**

```java

package com.pfs.ip.tools.common.thread;

import org.junit.Test;

public class ThreadLocalSample {

    class MyRunnable implements Runnable{

        private ThreadLocal<Integer> threadLocal = new ThreadLocal<Integer>(){
            @Override protected Integer initialValue() {
                return 18;
            }
        };

        @Override
        public void run() {
            System.out.println("ThreadLocal:" + threadLocal.get());
        }
    }

    @Test
    public void test_init_value() throws InterruptedException {
        MyRunnable work = new MyRunnable();

        Thread thread1 = new Thread(work);
        Thread thread2 = new Thread(work);

        thread1.start();
        thread2.start();

        Thread.sleep(1000 * 10);
    }
}

```

