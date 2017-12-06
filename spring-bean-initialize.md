---
layout: post
title: Spring bean initialization
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---

|      Item                   |                           Desc                               |
| --------------------------- | ------------------------------------------------------------ |
| InitializingBean            |                                                              |
| init-method                 |                                                              |
| BeanFactoryPostProcessor    |                                                              |
| BeanPostProcessor           |                                                              |
| PostConstruct               |                                                              |

# InitializingBean
```java
public interface InitializingBean {
    void afterPropertiesSet() throws Exception;
}
```

# init-method

# BeanFactoryPostProcessor
```java
public interface BeanFactoryPostProcessor {
    void postProcessBeanFactory(ConfigurableListableBeanFactory var1) throws BeansException;
}
```

# BeanPostProcessor
```java
public interface BeanPostProcessor {
	Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException;

	Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException;
}
```

# @PostConstruct
```java

```


# Bean的初始化顺序
1. 设置属性值
2. 调用`BeanNameAware.setBeanName()`
3. 调用`BeanFactoryAware.setBeanFactory()`
4. 调用`BeanPostProcessor.postProcessBeforeInitialization()`方法
5. 调用`InitializingBean.afterPropertiesSet()`方法
6. 调用`Bean的init-method`方法。通常是在配置bean的时候指定了`init-method`，例如:`<bean class="beanClass" init-method="init"></bean>`
7. 调用`BeanPostProcessor.postProcessAfterInitialization()`

