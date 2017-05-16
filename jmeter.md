---
layout: post
title: JMeter
date: 2017-05-16 12:20:00
tags:
- Java
categories: Java
description: The tutoria will describe the useage of Linux.
---



# 设置QPS限制
 
本次性能测试的需求中提到测试的目的是“了解博客的首页在负载达到20 QPS时的响应时间”，因此需要控制向博客首页发送请求的负载为20QPS。

一种可行的方法是逐步调整测试计划中的线程计算的数量 以及为取样器（Sampler）添加定时器（Timer），以使HTTP取样器发出的请求的QPS保持在20个左右。但这种方法耗时耗力，需要经过多次尝 试才能达到；另一方法，完全通过设置定时器来控制QPS，一旦取样器的响应时间发生改变（网络环境发生改变），就需要重新调整定时器的等待时间。

Jmeter提供了一个非常有用的定时器，称为Constant Throughput Timer （常数吞吐量定时器），该定时器可以方便地控制给定的取样器发送请求的吞吐量。

（添加--->定时器--->Constant Throughput Timer）选择Constant Throughput Timer

Constant Throughput Timer 的主要属性介绍：

名称 ：定时器的名称

Target throughput（in samples per minute）：目标吞吐量。注意这里是每分钟发送的请求数，因此，对应测试需求中所要求的20 QPS ，这里的值应该是1200 。

Calculate Throughput based on ：有5个选项，分别是：

　　This thread only ：控制每个线程的吞吐量，选择这种模式时，总的吞吐量为设置的 target Throughput 乘以矣线程的数量。

　　All active threads ： 设置的target Throughput 将分配在每个活跃线程上，每个活跃线程在上一次运行结束后等待合理的时间后再次运行。活跃线程指同一时刻同时运行的线程。

　　 All active threads in current thread group ：设置的target Throughput将分配在当前线程 组的每一个活跃线程上，当测试计划中只有一个线程组时，该选项和All active threads选项的效果完全相同。

　　All active threads （shared ）：与All active threads 的选项基本相同，唯一的区别是，每个活跃线程都会在所有活跃线程上一次运行结束后等待合理的时间后再次运行。

　　 All cative threads in current thread group （shared ）：与 All active threads in current thread group 基本相同，唯一的区别是，每个活跃线程都会在所有活跃线程的上 一次运行结束后等待合理的时间后再次运行。

 

设置定时器的Target throughput为1200/分钟（20 QPS），设置Calculate Throughput based on 的值为All active threads 。

　　当 然，Constant Throughput Timer只有在线程组中的线程产生足够多的request 的情况下才有意义，因此，即使设置了 Constant Throughput Timer的值，也可能由于线程组中的线程数量不够，或是定时器设置不合理等原因导致总体的QPS不能达到预期 目标。
  
  比如：如果想测试QPS为100，那么如下设置：
  Target throughput=6000
  线程数：50，不能太小，不然可能达不到100的QPS
  准备时长：1
  循环次数：不填，即永远
  
  
  
  

