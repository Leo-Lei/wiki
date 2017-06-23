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



# wait
当在一个对象的实例上调用wait()方法后，当前线程会变成等待状态。一直等到别的线程调用了这个对象实例的notify()方法。比如，线程T1中调用obj.wait()方法，那么线程T1就会进入等待状态。一段时间后，线程T2中调用了obj.notify()方法，这样，T1线程又可以继续执行了。这时，obj对象就成为多个线程间通信的手段。     
关于wait的使用，注意以下几点:        
* 必须在synchronized语句块中使用wait方法
* wait方法内部会释放持有的obj的monitor，即释放obj的锁。


# notify




# 用wait和notify实现的生产者-消费者模式

```java

import java.util.LinkedList; 
import java.util.Queue; 
import java.util.Random; 
/** 
* Simple Java program to demonstrate How to use wait, notify and notifyAll() 
* method in Java by solving producer consumer problem.
* 
* @author Javin Paul 
*/
public class ProducerConsumerInJava { 
    public static void main(String args[]) { 
        System.out.println("How to use wait and notify method in Java"); 
        System.out.println("Solving Producer Consumper Problem"); 
        Queue&lt;Integer&gt; buffer = new LinkedList&lt;&gt;(); 
        int maxSize = 10; 
        Thread producer = new Producer(buffer, maxSize, "PRODUCER"); 
        Thread consumer = new Consumer(buffer, maxSize, "CONSUMER"); 
        producer.start(); consumer.start(); } 
    } 
    /** 
    * Producer Thread will keep producing values for Consumer 
    * to consumer. It will use wait() method when Queue is full 
    * and use notify() method to send notification to Consumer 
    * Thread. 
    * 
    * @author WINDOWS 8 
    * 
    */
    class Producer extends Thread 
    { private Queue&lt;Integer&gt; queue; 
        private int maxSize; 
        public Producer(Queue&lt;Integer&gt; queue, int maxSize, String name){ 
            super(name); this.queue = queue; this.maxSize = maxSize; 
        } 
        @Override public void run() 
        { 
            while (true) 
                { 
                    synchronized (queue) { 
                        while (queue.size() == maxSize) { 
                            try { 
                                System.out .println("Queue is full, " + "Producer thread waiting for " + "consumer to take something from queue"); 
                                queue.wait(); 
                            } catch (Exception ex) { 
                                ex.printStackTrace(); } 
                            } 
                            Random random = new Random(); 
                            int i = random.nextInt(); 
                            System.out.println("Producing value : " + i); 
                            queue.add(i); 
                            queue.notifyAll(); 
                        } 
                    } 
                } 
            } 
    /** 
    * Consumer Thread will consumer values form shared queue. 
    * It will also use wait() method to wait if queue is 
    * empty. It will also use notify method to send 
    * notification to producer thread after consuming values 
    * from queue. 
    * 
    * @author WINDOWS 8 
    * 
    */
    class Consumer extends Thread { 
        private Queue&lt;Integer&gt; queue; 
        private int maxSize; 
        public Consumer(Queue&lt;Integer&gt; queue, int maxSize, String name){ 
            super(name); 
            this.queue = queue; 
            this.maxSize = maxSize; 
        } 
        @Override public void run() { 
            while (true) { 
                synchronized (queue) { 
                    while (queue.isEmpty()) { 
                        System.out.println("Queue is empty," + "Consumer thread is waiting" + " for producer thread to put something in queue"); 
                        try { 
                            queue.wait(); 
                        } catch (Exception ex) { 
                            ex.printStackTrace(); 
                        } 
                    } 
                    System.out.println("Consuming value : " + queue.remove());
                    queue.notifyAll();
                } 
            } 
        } 
    }

```




