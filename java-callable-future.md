---
layout: post
title: Java Concurrency Basic
date: 2017-06-28 14:40:00
tags:
- Java
- Concurrency
categories:
- Java
- Concurrency
description: The series tutorial talk about the Java Concurrency programing model
---

# Callable Future

|          name       | Desc                                    |
| ------------------- | --------------------------------------- |
| Runable             | 只有一个方法 void run()                   |
| Callable<T>         | 只有一个方法 V call()                     | 



```java
public class CallableAndFuture {
    public static void main(String[] args) {
        Callable<Integer> callable = new Callable<Integer>() {
            public Integer call() throws Exception {
                return new Random().nextInt(100);
            }
        };
        FutureTask<Integer> future = new FutureTask<Integer>(callable);
        new Thread(future).start();
        try {
            Thread.sleep(5000);// 可能做一些事情
            System.out.println(future.get());
        } catch (InterruptedException e) {
            e.printStackTrace();
        } catch (ExecutionException e) {
            e.printStackTrace();
        }
    }
}
```

