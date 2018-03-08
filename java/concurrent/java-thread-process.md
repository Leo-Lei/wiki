---
layout: post
title: Java Thread
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---

# 线程状态
* New: 新建状态。当线程创建完成时为新建状态，即new Thread(...)，但还没有调用start方法时。
* Runnable：就绪状态。当调用线程的start方法后，线程进入就绪状态，等待CPU资源。
* Running：运行状态。就绪状态的线程获取到CPU执行权后进入运行状态，开始执行run方法。
* Blocked：阻塞状态。线程没有执行完，由于某种原因，如I/O操作，线程同步等。让出CPU执行权，自身进入阻塞状态。
* Dead：死亡状态。线程执行完成，或者执行过程中出现异常，线程进入死亡状态。

# main线程结束，其他线程一样可以正常运行。
当main线程结束后，如果进程中还有其他的非守护线程，Java进程就不会退出。其他线程可以继续运行。
```java
public class ThreadTest {

    public void run(){
        System.out.println("Start-----");
        Thread.sleep(1000);     // 省略了try-catch
        System.out.println("End-------");
    }

    public static void main(String[] args) {
        final ThreadTest test = new ThreadTest();
        new Thread(() -> test.run()).start();
        System.out.println("Main thread exit.......");
    }
}
```
打印出:
```text
Main thread exit.......
Start-----
End-------
```

# 当main线程和所有非守护线程结束时，进程退出
Java进程退出的条件是：虚拟机中只剩下守护线程。


# Main线程是一个非守护线程，不能设置成守护线程
main线程是jvm启动时创建的。不能设置为守护线程。调用Thread.setDaemon()会抛出异常。
