---
layout: post
title: Java Wait and Notify
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---

# wait

|            method            |                                     remark                                 |
| ---------------------------- | -------------------------------------------------------------------------- |
| wait()                       | 将当前运行的线程挂起(即让其进入阻塞状态),直到别的线程用notify或notifyAll方法来唤醒    |
| wait(long timeout)           | 与wait()类似。区别就是在指定时间内，如果没有notify或notifyAll唤醒，也会自动唤醒      |
| wait(long timeout,int nanos) | 更精准的控制调度时间，精确到纳秒                                                 |



当在一个对象的实例上调用wait()方法后，当前线程会变成等待状态。一直等到别的线程调用了这个对象实例的notify()方法。比如，线程T1中调用obj.wait()方法，那么线程T1就会进入等待状态。一段时间后，线程T2中调用了obj.notify()方法，这样，T1线程又可以继续执行了。这时，obj对象就成为多个线程间通信的手段。     
关于wait的使用，注意以下几点:        
* 必须在synchronized语句块中使用wait方法
* wait方法内部会释放持有的obj的monitor，即释放obj的锁。


# notify



# 只能在Synchronized块中使用wait和notify
```java
public class WaitNotifyTest {

    public void testWait(){
        System.out.println("Start-----");
        wait(1000);    // 省略了try-catch块
        System.out.println("End-------");
    }

    public static void main(String[] args) {
        final WaitNotifyTest test = new WaitNotifyTest();
        new Thread(() -> test.testWait()).start();
    }
}
```
运行上面的代码，抛出异常:***java.lang.IllegalMonitorStateException***。在JDK中对IllegalMonitorStateException的描述:
> 线程试图等待对象的monitor或试图通知其他正在等待对象monitor的线程，但本身没有对应的monitor的所有权。    

wait是一个本地方法，其底层是通过对象的monitor来实现的。上面会出现这个异常，是因为在调用wait方法时，没有获取到monitor对象的所有权。那如何获取到monitor对象的所有权?Java中只能通过Synchronized关键字来完成。



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
        Queue<Integer> buffer = new LinkedList<>(); 
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
    { private Queue<Integer> queue; 
        private int maxSize; 
        public Producer(Queue<Integer> queue, int maxSize, String name){ 
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
        private Queue<Integer> queue; 
        private int maxSize; 
        public Consumer(Queue<Integer> queue, int maxSize, String name){ 
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



