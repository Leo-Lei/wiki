---
layout: post
title: Spring
date: 2016-11-18 20:40:00
tags:
- Java
categories: Java
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
| `@Primary`          | 在@Autowired的时候，如果一个type有多个实现，可以通过@Primary，告诉Spring优先选择该组件。  |
| `@Qualifier`        | 当@Autowired时，如果有多个bean，通过@Qualifier指定bean的id来注入                      |


# @EnableConfigurationProperties, @Value
```java
@Configuration
@ConditionalOnClass(Exporter.class)
@EnableConfigurationProperties({DubboApplication.class, DubboRegistry.class})
public class DubboAutoConfiguration {

    @Autowired
    private DubboApplication dubboApplication;

    @Autowired
    private DubboRegistry dubboRegistry;

    @Bean
    public static AnnotationBean annotationBean(@Value("${dubbo.annotation.package}") String packageName){
        AnnotationBean annotationBean = new AnnotationBean();
        annotationBean.setPackage(packageName);
        System.out.println("[DubboAutoConfiguration] : " + packageName);
        return annotationBean;
    }

    @Bean
    public ApplicationConfig applicationConfig(){
        ApplicationConfig applicationConfig = new ApplicationConfig();
        applicationConfig.setName(dubboApplication.getName());
        applicationConfig.setLogger(dubboApplication.getLogger());
        return applicationConfig;
    }
    
    @Bean
    public RegistryConfig registryConfig(@Value("${data.dir:}") String dataDir) {

        RegistryConfig registryConfig = new RegistryConfig();
        registryConfig.setProtocol(dubboRegistry.getProtocol());
        registryConfig.setAddress(dubboRegistry.getAddress());
        registryConfig.setRegister(dubboRegistry.isRegister());
        registryConfig.setSubscribe(dubboRegistry.isSubscribe());
        registryConfig.setClient("curator");

        if (StringUtil.isNotBlank(dataDir)){
            registryConfig.setFile(dataDir+"/" + ".dubbo");
        }
        System.out.println(registryConfig.getFile());

        System.out.println("[DubboAutoConfiguration] : " + dubboRegistry);
        return registryConfig;
    }
}

```

```java
@ConfigurationProperties(prefix = "dubbo.registry")
public class DubboRegistry {
    public String getProtocol() {
        return protocol;
    }

    public void setProtocol(String protocol) {
        this.protocol = protocol;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public boolean isRegister() {
        return register;
    }

    public void setRegister(boolean register) {
        this.register = register;
    }

    public boolean isSubscribe() {
        return subscribe;
    }

    public void setSubscribe(boolean subscribe) {
        this.subscribe = subscribe;
    }

    private String protocol = "zookeeper";
    private String address = "127.0.0.1:2181";
    private boolean register = true;

    private boolean subscribe = true;
}

```

```properties
#接口协议
dubbo.registr.protocol=zookeeper

#注册中心地址
dubbo.registry.address=127.0.0.1:2181

#是否向注册中心注册服务
dubbo.registry.register=true

#是否向注册中心订阅服务
dubbo.registry.subscribe=true
```



# @ConditionalOnProperty
@ConditionalOnProperty注解允许根据一个Spring Environment属性来决定是否包含配置。可以使用prefix和name属性指定要检查的配置属性。
默认情况下，只要满足以下条件，就会匹配:
1. 属性必须存在
2. 属性不为false
你也可以使用havingValue和matchIfMissing属性创建更高级的检测。
* havingValue: 

|   Property Value  |  havingValue=""  | havingValue="true" | havingValue="false"  |  havingValue="foo" |
| ----------------- | ---------------- | ------------------ | -------------------- | ------------------ |
| "true"            | yes              | yes                | no                   | no                 |
| "false"           | no               | no                 | yes                  | no                 |
| "foo"             | yes              | no                 | no                   | yes                |

```java
@ConditionalOnProperty(prefix = "spring.application.admin", value = "enabled", havingValue = "true", matchIfMissing = false)
```
