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
# xml文件，使用setter-based依赖注入
```xml
<bean id="exampleBean" class="examples.ExampleBean">
    <!-- setter injection using the nested ref element -->
    <property name="beanOne">
        <ref bean="anotherExampleBean"/>
    </property>

    <!-- setter injection using the neater ref attribute -->
    <property name="beanTwo" ref="yetAnotherBean"/>
    <property name="integerProperty" value="1"/>
</bean>

<bean id="anotherExampleBean" class="examples.AnotherBean"/>
<bean id="yetAnotherBean" class="examples.YetAnotherBean"/>
```
```java
public class ExampleBean {

    private AnotherBean beanOne;
    private YetAnotherBean beanTwo;
    private int i;

    public void setBeanOne(AnotherBean beanOne) {
        this.beanOne = beanOne;
    }

    public void setBeanTwo(YetAnotherBean beanTwo) {
        this.beanTwo = beanTwo;
    }

    public void setIntegerProperty(int i) {
        this.i = i;
    }
}
```
# lazy init
```xml
<bean id="lazy" class="com.foo.ExpensiveToCreateBean" lazy-init="true"/>
<bean name="not.lazy" class="com.foo.AnotherBean"/>
```
当spring的ApplicationContext启动后，如果没有lazy-init的单例的bean，会立即被创建。如果有lazy-init，会在第一次被请求的时候被创建。大部分的bean都不是lazy-init的，可以尽快的发现spring bean的配置上的错误，比如某个bean没有被注入。
> 如果一个lazy-init的bean是另一个单例的非lazy-init的bean的依赖，spring的ApplicationContext启动时，也会立即创建lazy-init的bean。因为spring必须要配置单例的bean的依赖。


