---
layout: post
title: Dubbo集群的负载均衡
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---

# 背景
Dubbo是一个分布式服务框架，能避免单点故障和支持服务的横向扩容。一个服务通常会部署多个实例。如何从多个服务提供者组成的集群中挑选出一个进行调用，就涉及到一个负载均衡的策略。


# 几个概念
在说到负载均衡时，还有下面的几个概念也经常被提及:
1. 集群容错
2. 服务路由

很多时候，我们可能会混淆这些概念。他们都是描述了怎么从多个provider中选择一个来进行调用。那他们到底有什么区别呢?下面我来举一个简单的例子，把这几个概念阐述清楚吧。
有一个用户服务，在北京部署了10个，在上海部署了20个。一个杭州的服务消费方发起了一次调用，然后发生了以下的事情:
1. 服务治理系统中预先设置了一个路由，就是如果杭州发起的调用，会路由到比较近的上海的服务提供方。这时候，会选择上海的20个服务提供方
2. 系统中配置了随机的负载均衡策略。在这20个上海的服务器中随机选择了一个来调用，假设随机到了第7个Provider。
3. 调用的时候，失败了，这时候怎么办呢？由于配置了Failover集群容错模式，会重试其他服务器。重试了第13个Provider，调用成功。        

上面的第1，2，3步骤就分别对应了路由，负载均衡和集群容错。

# Dubbo内置负载均衡策略
Dubbo内置了4种负载均衡策略:
1. RandomLoadBalance:随机负载均衡。随机的选择一个。是Dubbo的**默认**负载均衡策略。
2. RoundRobinLoadBalance:轮询负载均衡。轮询选择一个。
3. LeastActiveLoadBalance:最少活跃调用数负载均衡。最少活跃调用数，相同活跃数的随机，活跃数指调用前后计数差。使慢的提供者收到更少请求，因为越慢的提供者的调用前后计数差会越大。
4. ConsistentHashLoadBalance:一致性哈希负载均衡。相同参数的请求总是落在同一台机器上。

# 源码
### LoadBalance:
```java
@SPI(RandomLoadBalance.NAME)
public interface LoadBalance {
    @Adaptive("loadbalance")
    <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException;
}
```
这是SPI的接口，select方法的参数如下:
* invokers: 所有的服务提供者列表。
* url: 一些配置信息，比如接口名，是否check，序列化方式。
* invocation: RPC调用的信息，包括方法名，方法参数类型，方法参数。

### RandomLoadBalance
```java
public class RandomLoadBalance extends AbstractLoadBalance {

    private final Random random = new Random();

    protected <T> Invoker<T> doSelect(List<Invoker<T>> invokers, URL url, Invocation invocation) {
        int length = invokers.size(); // Number of invokers
        int totalWeight = 0; // The sum of weights
        
        // 判断是不是所有的invoker的权重都是一样的
        boolean sameWeight = true; 
        for (int i = 0; i < length; i++) {
            int weight = getWeight(invokers.get(i), invocation);
            totalWeight += weight; // Sum
            if (sameWeight && i > 0 && weight != getWeight(invokers.get(i - 1), invocation)) {
                sameWeight = false;
            }
        }
        if (totalWeight > 0 && !sameWeight) {
            // 如果不是所有的invoker权重都相同，那么基于权重来随机选择。权重越大的，被选中的概率越大
            int offset = random.nextInt(totalWeight);
            for (int i = 0; i < length; i++) {
                offset -= getWeight(invokers.get(i), invocation);
                if (offset < 0) {
                    return invokers.get(i);
                }
            }
        }
        // 如果所有invoker权重相同
        return invokers.get(random.nextInt(length));
    }
}
```
Dubbo的随机负载均衡有权重的概念。Dubbo中某一个服务提供者，可以对其设置权重。比如机器性能好的，可以设置大一点的权重，性能差的，可以设置小一点的权重。权重会对负载均衡产生影响。可以在Dubbo Admin中对provider进行权重的设置。        
随机策略会先判断所有的invoker的权重是不是一样的，如果都是一样的，那么处理就比较简单了。使用random.nexInt(length)就可以随机生成一个invoker的序号。如果没有在Dubbo Admin中对服务提供者设置权重，那么所有的invoker的权重就是一样的，默认是100。
如果权重不一样，那就需要结合权重来设置随机概率了。算法大概如下：
假如有4个invoker

|  invoker  |   weight   |
| --------- | ---------- |
|     A     |     10     |
|     B     |     20     |
|     C     |     20     |
|     D     |     30     |

A，B，C和D总的权重是10 + 20 + 20 + 30 = 80。将80个数按照如下区域划分:

```text
+-----------------------------------------------------------------------------------+
|          |                    |                    |                              |
+-----------------------------------------------------------------------------------+
           10                   30                   50                             80            

|          |
|-----v----|---------v----------| ---------v-------- | --------------v--------------|
      A              B                     C                         D
```
有上面的一个分部图，一共有4块区域，长度分别是A，B，C和D的权重。使用random.nextInt(10 + 20 + 20 + 30)，从80个数中随机选择一个。然后再判断该数分布在哪个区域。比如，如果随机到27，27是分布在B区域的，那么就选择inboker B。


# 负载均衡扩展
Dubbo的4种负载均衡的实现，大多数情况下能满足要求。有时候，因为业务的需要，我们可能需要实现自己的负载均衡策略。

1. 实现LoadBalance接口
```java
package com.leibangzhu.test.dubbo.consumer;
public class MyLoadBalance implements LoadBalance {
    @Override
    public <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException {
        System.out.println("Select the first invoker...");
        return invokers.get(0);
    }
}
```
2. 添加资源文件
添加文件:`src/main/resource/META-INF/dubbo/com.alibaba.dubbo.rpc.cluster.LoadBalance`。这是一个简单的文本文件。文件内容如下:
```text
my=my=com.leibangzhu.test.dubbo.consumer.MyLoadBalance
```
3. 配置使用自定义LoadBalance
```xml
<dubbo:reference id="helloService" interface="com.leibangzhu.test.dubbo.api.IHelloService" loadbalance="my" />
```
在consumer端的<dubbo:reference>中配置<loadbalance="my">

经过上面的3个步骤，我们编写了一个自定义的LoadBalance，并告诉Dubbo使用它了。启动Dubbo，我们就能看到Dubbo已经使用了自定义的MyLoadBalance。
