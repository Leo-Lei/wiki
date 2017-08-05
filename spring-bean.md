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
`class`属性指定包含静态factory method的类，`factory-method`指定静态方法名。
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
# 使用实例factory method来定义bean

```xml
<!-- the factory bean, which contains a method called createInstance() -->
<bean id="serviceLocator" class="examples.DefaultServiceLocator">
    <!-- inject any dependencies required by this locator bean -->
</bean>

<!-- the bean to be created via the factory bean -->
<bean id="clientService"
    factory-bean="serviceLocator"
    factory-method="createClientServiceInstance"/>
```
```java
public class DefaultServiceLocator {

    private static ClientService clientService = new ClientServiceImpl();
    private DefaultServiceLocator() {}

    public ClientService createClientServiceInstance() {
        return clientService;
    }
}
```
> 一个factory bean可以有多个factory method。

# 为bean指定constructor-arg
```xml
<beans>
    <bean id="exampleBean" class="examples.ExampleBean">
        <constructor-arg name="years" value="7500000"/>
        <constructor-arg name="ultimateAnswer" value="42"/>
        <constructor-arg name="foo" ref="foo"/>
    </bean>

    <bean id="foo" class="x.y.Foo"/>
</beans>
```
