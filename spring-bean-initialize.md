---
layout: post
title: Spring bean initialization
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---


# @InitializingBean
```java
public interface InitializingBean {
    void afterPropertiesSet() throws Exception;
}
```

# init-method

# @BeanFactoryPostProcessor
```java
public interface BeanFactoryPostProcessor {
    void postProcessBeanFactory(ConfigurableListableBeanFactory var1) throws BeansException;
}
```

# @BeanPostProcessor
```java
public interface BeanPostProcessor {
	Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException;

	Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException;
}

```
