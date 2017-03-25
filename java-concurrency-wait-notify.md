---
layout: post
title: Java Wait and Notify
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
description: The tutoria will describe the useage of wait and notify.
---



# 1.Java Thread Synchronization Principle
In java, each object has an monitor. when multiple thrad are executing some synchronized methods or synchronized block, the monitor is responsible for coordinating the concurrent requests.
When a thread is entering a synchronized method of a object, JVM will check the monitor of this object. If the monitor is not used, the thread can get the monitor, and continue to execute the synchronized method. But if the monitor is owned by other thread, the thread will be hang up, untill the monitor is released.
When thread exit the synchronized method, it will release the monitor, this will let other waiting thread to get the monitor to continue execute method.

# 2.Wait and Notify Method
## 2.1 Wait 
`obj.wait()` method will hang up current method, and release object's monitor. When other thread call the obj.notify() or obj.notifyAll() method, the thread will be waked up and continue to execute.
## Notify
`obj.notify()` method will wake up a sleeping thread which is blocked on obj randomly. Then the waked up thread will get the monitor and continue to execute.

Pay attention to the notify method, after calling the notify method, the monitor will not be released immediately, the monitor will be released utill the thread exit from the synchronized method. Then the waked up sleeping thread can get the monitor.

## Waitign Thread Queue
Each obj has an waiting thread queue. The threads in the queue is waiting for the monitor to execute some synchronized method. There are 2 way to enter the queue: 
1. Other thread is owned the monitor, current thread will wait.
2. Call the obj.wait() method.

# 3.Some samples about Wait and Notify

```java
package com.pfs.ip.tools.common.thread;

public class MyThreadPrinter implements Runnable {

    private String name;
    private Object prev;
    private Object self;

    public MyThreadPrinter(String name, Object prev, Object self) {
        this.name = name;
        this.prev = prev;
        this.self = self;
    }

    @Override
    public void run() {
        int count = 5;
        while (count > 0) {
            synchronized (prev) {
                synchronized (self) {
                    System.out.println(name);

                    try{
                        Thread.sleep(1);
                    }
                    catch (InterruptedException e){
                        e.printStackTrace();
                    }

                    self.notify();
                    System.out.println( String.format("Thread %s: notify %s.",name,self));
                }
                if(count>1){
                    try {
                        System.out.println( String.format("Thread %s: wait for %s.",name,prev));
                        prev.wait();
                        System.out.println( String.format("Thread %s: %s is notified, continue to run.",name,prev));
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
            count--;
        }
        System.out.println( String.format("Thread %s: Exit.",name));
    }

    public static void main(String[] args) throws Exception {
        Object a = new String("A");
        Object b = new String("B");
        Object c = new String("C");
        MyThreadPrinter pa = new MyThreadPrinter("A", c, a);
        MyThreadPrinter pb = new MyThreadPrinter("B", a, b);
        MyThreadPrinter pc = new MyThreadPrinter("C", b, c);

        new Thread(pa).start();
        Thread.sleep(100);
        new Thread(pb).start();
        Thread.sleep(100);
        new Thread(pc).start();
        Thread.sleep(100);

        //Thread.sleep(1000 * 10);
        System.out.println("Finish!");
    }
}

```

