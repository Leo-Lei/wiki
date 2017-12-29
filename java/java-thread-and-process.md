---
layout: post
title: Java Thread and Process
date: 2015-06-30 18:00:00
tags:
- Java
- Concurrency
categories:
- Java
- Concurrency
---

# 1. Thread Basis


# 2. Create a Thread

# 3. Common API of Thread

# 3.1 Join
1. What does join do
The join method will let the main thread wait the sub-thread to finish. For example, the main thread create and then start a new sub-thread. The sub-thread will do some caculate work, and it may take some time. But the main thread will need the caculate result of the sub-thread. So the main thread need to wait for the sub-thread to finish. At this situation, we can call the join method of sub-thread.
2. How to use `join`
The following code snap show how to use the `join` method.

```java
Thread t = new Thread(); // Create a new sub-thread.
t.start();     // Start the thread.
t.join();      // Invoke the join method on sub-thread.
```
