---
layout: post
title: Java Synchronized
date: 2015-08-11 14:40:00
tags:
- Java
categories: Java
---

# Volatile的作用
Volatile有2个作用：
1. 保证共享变量的可见性
2. 解决重排序问题

注意： Volatile不能保证原子性。

# Java内存模型
假设有变量count=10, 线程a对其进行读写
### 没有volatile修饰
1. 线程a从主存中读取count=10，将count=10保存到线程a的工作内存中。
2. 线程a在工作内存中进行操作count=count+1，所以，工作内存中count=11。
3. 这个时候，主存中的count还是10。工作内存中的count=11还没有同步到主存中。
4. 经过了一段不确定的时间，线程a的工作内存中的count=11同步到了主存中。
5. 当线程退出的时候，也会把count=11同步到主存中。

### 有volatile修饰
count=10，线程a和b对其进行读写
1. 线程a从主存中读取count=10，保存到线程a的工作内存中
2. 线程b从主存中读取count=10，保存到线程b的工作内存中
3. 线程a将count自增，count=11
4. 由于count是valatile的，当线程a将count=11写入线程a的工作线程中时，count=11立即被同步到了主存中
5. 同时，CPU中的所有线程的工作缓存中的count都被设置为无效了。
6. 当线程b需要读取count的值时，因为线程b的工作线程中的count已无效了。线程b从主存中去读取，count=11。是最新的值

# Volatile实现可见性
可见性问题是指一个线程修改了共享变量值，而另一个线程却看不到，这是由于JVM的内存模型决定的。每个线程在CPU中都有自己的一个高速缓存区--线程工作内存。volatile可以解决这个问题。
先看下面的代码：
```java
boolean stop = false;

// 线程1
while(!stop){
    doSomething();
}

// 线程2
stop = true;
```
这是一段很典型的代码，在线程2中将stop设置为true，来中断线程1的工作。但这段代码，不一定能正常工作。就是因为，线程2写入的stop=true，可能还没有从线程的工作线程同步到主存中。这时候，stop=true对线程1是不可见的。给stop加上了volatile后，就可以了。

# Volatile不能保证原子性
```java
public class Test {
    public volatile int count = 0;
     
    public void increase() {
        count++;
    }
     
    public static void main(String[] args) {
        final Test test = new Test();
        for(int i=0;i<10;i++){
            new Thread(){
                public void run() {
                    for(int j=0;j<1000;j++)
                        test.increase();
                };
            }.start();
        }
         
        while(Thread.activeCount()>1)  //保证前面的线程都执行完
            Thread.yield();
        System.out.println(test.inc);
    }
}
```
上面这段代码，有10个线程，每个线程对i增加1000。inc也是由valatile修饰过了。也许有人认为结果是10000。但事实是每次运行结果都不一样，都小于10000。
自增操作不是原子性的，它包含读取变量原始值，进行加1操作，写入内存。那么自增的三个操作可能进行分割，就有可能导致下面这种情况：
1. 假设某个时刻count=10
2. 线程1开始对变量进行自增操作。线程1从主存中读取count=10。但还没有进行加1操作。可能CPU切换去做别的事情了。
3. 线程2开始对变量进行自增操作。线程2从主存中读取count=10。
4. 线程2开始对count进行加1操作。count=count+1，这样count=11。线程将count=11写入线程的工作内存。
5. 因为有volatile修饰。线程2的工作内存中的count=11立即同步到主存中。
6. 线程1接着进行加1操作。由于线程已经读取了count的值，线程1的工作内存中count=10。线程不再从主存中去读取count值。线程1对count加1，count=11。然后将count=11写入工作，同步到主存中。

上面的代码之所以，不能正常运行，就是因为count++操作不是原子性的。它是被拆分为3个小的操作的。这3个小的操作在执行的时候，可能其中某几个操作失败了。可能这3个操作不是连续的，在执行了其中1个之后，CPU切换去做别的事情了。有其他操作影响了这3个小的操作。                
可以对上面的代码进行改进：
使用Synchronized：
```java
public class Test {
    public  int inc = 0;
    
    public synchronized void increase() {
        inc++;
    }
    
    public static void main(String[] args) {
        final Test test = new Test();
        for(int i=0;i<10;i++){
            new Thread(){
                public void run() {
                    for(int j=0;j<1000;j++)
                        test.increase();
                };
            }.start();
        }
        
        while(Thread.activeCount()>1)  //保证前面的线程都执行完
            Thread.yield();
        System.out.println(test.inc);
    }
}
```
使用Lock：
```java
public class Test {
    public  int inc = 0;
    Lock lock = new ReentrantLock();
    
    public  void increase() {
        lock.lock();
        try {
            inc++;
        } finally{
            lock.unlock();
        }
    }
    
    public static void main(String[] args) {
        final Test test = new Test();
        for(int i=0;i<10;i++){
            new Thread(){
                public void run() {
                    for(int j=0;j<1000;j++)
                        test.increase();
                };
            }.start();
        }
        
        while(Thread.activeCount()>1)  //保证前面的线程都执行完
            Thread.yield();
        System.out.println(test.inc);
    }
}
```
使用AtomicInteger：
```java
public class Test {
    public  AtomicInteger inc = new AtomicInteger();
     
    public  void increase() {
        inc.getAndIncrement();
    }
    
    public static void main(String[] args) {
        final Test test = new Test();
        for(int i=0;i<10;i++){
            new Thread(){
                public void run() {
                    for(int j=0;j<1000;j++)
                        test.increase();
                };
            }.start();
        }
        
        while(Thread.activeCount()>1)  //保证前面的线程都执行完
            Thread.yield();
        System.out.println(test.inc);
    }
}
```

