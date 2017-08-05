---
layout: post
title: Spring-bean
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---

# 使用xml定义bean
这是最传统的定义bean的方式。
# 使用静态factory method来定义bean
```xml
<bean id="clientService"
    class="examples.ClientService"
    factory-method="createInstance"/>
```
```java
public class ClientService {
    private static ClientService clientService = new ClientService();
    private ClientService() {}

    public static ClientService createInstance() {
        return clientService;
    }
}
```
