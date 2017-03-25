---
layout: post
title: Java Concurrency (11) CyclicBarrier
date: 2015-07-19 14:40:00
tags:
- Java
- Concurrency
categories:
- Java
- Concurrency
description: The series tutorial talk about the Java Concurrency programing model
---

# 1. The API of CyclicBarrier
The common API for CyclicBarrier is as below.

| method                                        | description                                                           |
| :-------------------------------------------- | :-------------------------------------------------------------------- |
| CyclicBarrier(int parties)                    | Create a new CyclicBarrier.                                           |
| int await()                                   |                                                                       |

## 1.1 await method
The await method will block current thread, until all the thread are arriving the barrier. This method return a int value, the int value is the index of the thread. For example, there are total 10 threads, the first arriving thread which first call the await method, will return 10 - 1 = 9. The second will return 8, the last one will return 0.

# 2. Sample of CyclicBarrier

```java
package com.leo.gemini.javaconcurrency;

import java.util.Random;
import java.util.concurrent.BrokenBarrierException;
import java.util.concurrent.CyclicBarrier;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class CyclicBarrierDemo2 {

    public void main(String[] args){
        CyclicBarrier barrier = new CyclicBarrier(3);

        ExecutorService executor = Executors.newFixedThreadPool(3);
        executor.submit(new Thread(new Runner(barrier,"No 1")));
        executor.submit(new Thread(new Runner(barrier,"No 2")));
        executor.submit(new Thread(new Runner(barrier,"No 3")));

        executor.shutdown();
    }


    private class Runner implements Runnable{

        private CyclicBarrier barrier;

        private String name;

        public Runner(CyclicBarrier barrier,String name){
            this.barrier = barrier;
            this.name = name;
        }

        @Override
        public void run() {

            try {
                Thread.sleep(1000 * (new Random().nextInt(8)));
                System.out.println(name + " is ready.");
                barrier.await();
            } catch (InterruptedException e) {
                e.printStackTrace();
            } catch (BrokenBarrierException e) {
                e.printStackTrace();
            }
            System.out.println(name + " run.");
        }
    }
}
```
