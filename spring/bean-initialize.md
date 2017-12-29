---
layout: post
title: Spring bean initialization
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
---

|      Item                   |                           Desc                                                       |
| --------------------------- | ------------------------------------------------------------------------------------ |
| InitializingBean            | Spring提供的接口.bean初始化时会调用接口的afterPropertiesSet方法。                          |
| DisposableBean              | Spring提供的接口.bean析构时会调用接口的destroy方法。                                       |
| init-method                 |                                                                                      |
| BeanFactoryPostProcessor    | Spring加载完所有bean的元数据，但还没有去实例化bean的时候，执行`postProcessBeanFactory`方法    |
| BeanPostProcessor           |                                                                                      |
| @PostConstruct              |                                                                                      |
| @PreDestroy                 | JSR-250标准中提供的注解                                                                 |


# InitializingBean
`org.springframework.beans.factory.InitializingBean`接口，在容器设置好bean的所有必要属性后，执行初始化事情。
```java
public interface InitializingBean {
    void afterPropertiesSet() throws Exception;
}
```

# DisposableBean
`org.springframework.beans.factory.DisposableBean`接口，允许一个bean当容器需要其销毁时获得一次回调。
```java
void destroy() throws Exception;
```

# init-method


# 多种生命周期机制
不同的生命周期机制，执行顺序如下:
1. @PostConstruct,@PreDestroy
2. InitializingBean和DisposableBean
3. 自定义init(),destroy()


# BeanFactoryPostProcessor
```java
public interface BeanFactoryPostProcessor {
    void postProcessBeanFactory(ConfigurableListableBeanFactory var1) throws BeansException;
}
```

# BeanPostProcessor
1. BeanPostProcessor是在Spring创建每一个Bean时，都会执行的。BeanFactoryPostProcessor只执行一次。
2. BeanPostProcessor在比较早的阶段被实例化。在其他所有普通bean被实例化之前。

```java
public interface BeanPostProcessor {
	Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException;

	Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException;
}
```

# @PostConstruct
```java

```

# ApplicationContextAware
bean对ApplicationContext是感知的，可以通过获取ApplicationContext获取BeanFactory，来获取容器里面的Bean等。        
这些接口的使用将您的代码联系到Spring API，并且不遵循反转控制方式。因此，它们被推荐用于要求对容器进行编程访问的基础bean。


# Bean的初始化顺序
* 设置属性值
* 调用`BeanNameAware.setBeanName()`
* 调用`BeanFactoryAware.setBeanFactory()`
* 设置bean的属性，依赖等,到这里，bean的所有字段基本已经实例化好了
* 调用`BeanPostProcessor.postProcessBeforeInitialization()`方法
* 调用@PostConstruct修饰的方法
* 调用`InitializingBean.afterPropertiesSet()`方法
* 调用`Bean的init-method`方法。通常是在配置bean的时候指定了`init-method`，例如:`<bean class="beanClass" init-method="init"></bean>`
* 调用`BeanPostProcessor.postProcessAfterInitialization()`

# Spring容器中所有bean的加载顺序
我们通过xml,java注解或者java code的方式定义了很多的java bean，那么Spring容器在获得了这些bean的元数据之后，是按照什么顺序来初始化他们呢？比如，有些bean先，有些bean后初始化。
1. 初始化BeanFactoryPostProcessor类型的bean
2. 初始化BeanPostProcessor类型的bean
3. 初始化普通bean
