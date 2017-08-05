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
`src/main/resources/applicationContext.xml`
```xml
<?xml version="1.0" encoding="UTF-8"?>
<beans xmlns="http://www.springframework.org/schema/beans"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xmlns:context="http://www.springframework.org/schema/context"
       xsi:schemaLocation="http://www.springframework.org/schema/beans http://www.springframework.org/schema/beans/spring-beans.xsd http://www.springframework.org/schema/context http://www.springframework.org/schema/context/spring-context.xsd">

    <context:component-scan base-package="com.leolei.rocketmq"/>
    <context:property-placeholder location="classpath:mq.properties" ignore-unresolvable="true"/>
</beans>
```

`src/main/resources/mq.properties`:
```properties
mq.producer.groupName=ProducerGroupName
mq.consumer.groupName=ConsumerGroupName100
mq.nameServer=localhost:9876
mq.producer.instanceName=Producer
mq.consumer.instanceName=Consumer
```

`MQConfiguration.java`:
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



|      annotation     |                            desc                                                 |
| ------------------- | ------------------------------------------------------------------------------- |
| `@Primary`          | 在@Autowired的时候，如果一个type有多个实现，可以通过@Primary，告诉Spring优先选择该组件。  |
| `@Qualifier`        | 当@Autowired时，如果有多个bean，通过@Qualifier指定bean的id来注入                      |


