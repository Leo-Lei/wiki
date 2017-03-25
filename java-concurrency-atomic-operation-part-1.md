---
layout: post
title: Java Concurrency (2) Atomic operation part 1
date: 2015-06-30 17:00:00
tags:
- Java
- Concurrency
categories:
- Java
- Concurrency
description: The series tutorial talk about the Java Concurrency programing model
---

# 1. Talk about the i++ operation
By default, the `i++` or `i--` is not a thread-safe operation, as it is not a atomic operation. The `i++` is made up of 3 individual operations:
1. read the value of variable i.
2. add 1 to i.
3. write the i to new value.
I think most programers should know why the above `i++` operation is not thread-safe. I will not explain more about `i++` operation here. In order to let the `i++` thread-safe, we may use the `synchronized` keyword.
But is there another simple, high-performance and thread-safe increacing/decreacing solution? Yes! Java supply some Class to do this. For example, AtomicInteger.

# 2. Begin From AtomicInteger

The common API for AtomicInteger:

| method                                        | Description                                                                               |
| :-------------------------------------------- | :---------------------------------------------------------------------------------------- |
| int addAndGet(int delta)                      | Add delta atomically. A thread-safe version of `i = i + delta`.                           |
| int decrementAndGet()                         | Reduce 1 atomically. A thread-safe version of `i = i - 1`.                                |
| int get()                                     | Get current value.                                                                        |
| int getAndAdd(int delta)                      | A thread-safe version of `t = i; i += delta; return t;`                                   |
| int getAndIncrement()                         | A thread-safe version of `t = i; i ++; return t;`                                         |
| int getAndDecrement()                         | A thread-safe version of `t = i; i --; return t;`                                         |
| boolean compareAndSet(int except, int update) | If the value is except, update the value atomically and return true, while return false.  |
