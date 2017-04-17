---
layout: post
title: Spring
date: 2016-11-18 20:40:00
tags:
- Java
categories: Java
description: spring
---

# Java Code注入Bean

```java
package com.leolei.rocketmq;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class MQConfiguration {

    @Bean(name = "rocketMqProducerA")
    public IMsgProducer rocketMqProducerA(@Value("${mq.nameServer}") String nameServer) {

        return MQ.RocketMQ.producer()
                .groupName("com_leolei_rocketmq_producer_A")
                .nameServer(nameServer)
                .build();
    }

    @Bean(name = "rocketMqConsumerA")
    public IMsgConsumer rocketMqConsumerA(@Value("${mq.nameServer}") String nameServer,
                                          @Value("${mq.consumer.instanceName}") String instanceName){

        return MQ.RocketMQ.consumer()
                .groupName("com_leolei_rocketmq_consumer_A")
                .nameServer(nameServer)
                .instanceName(instanceName)
                .subscribe("TopicTest")
                .registerMsgHandler(msg -> {
                    System.out.println( "Consumer A start to process message: ");
                    return IMsgHandler.MsgHandleStatus.SUCCESS;
                })
                .build();
    }
}

```
