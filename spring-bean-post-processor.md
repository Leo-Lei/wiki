---
layout: post
title: Spring bean post processor
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
---

# BeanPostProcessor
BeanPostProcessor在Spring容器实例化好bean后，但还没执行bean的一些初始化方法之前，执行@BeanPostProcessor的某些方法。        
```java
public interface BeanPostProcessor {
    Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException;
    Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException;
}
```
其中：        
1. postProcessBeforeInitialization: Spring实例化bean后，执行初始化方法之前被调用        
2. postProcessAfterInitialization: Spring实例化bean，并执行完初始化方法之前被调用        

bean的初始化方法包括：
1. @PreConstruct
2. InitializingBean
3. 自定义init方法
        
* Spring容器每实例化完一个bean后，都会调用BeanPostProcessor。        

# Demo
```java
public interface IBar {
}
```

```java
@Component
public class Foo {
    private String id = "foo";

    @Autowired
    private IBar bar;

    public Foo(){
        System.out.println("foo construct...");

    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }
}
```


```java
@Component
public class Bar implements IBar {

    private String name = "bar";

    public Bar(){
        System.out.println("Bar construct...");
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
```

```java
@Component
public class SomeBeanPostProcessor implements BeanPostProcessor {

    @Override
    public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException {
        System.out.println("SomeBeanPostProcessor.postProcessBeforeInitialization ..." + beanName);
        return bean;
    }

    @Override
    public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException {
        System.out.println("SomeBeanPostProcessor.postProcessAfterInitialization ..." + beanName);
        return bean;
    }
}
```

# 执行顺序

1. spring初始化BeanPostProcessor
2. Bar.name = "bar"
3. Bar()构造函数
4. BeanPostProcessor.postProcessBeforeInitialization(Bar bar,"bar")
5. BeanPostProcessor.postProcessAfterInitialization(Bar bar,"bar")
6. Foo.id = "foo"
7. Foo()构造函数,此时，Foo.bar=null
8. BeanPostProcessor.postProcessBeforeInitialization(Foo,"foo"). Foo.bar已初始化好了。将Bar注入了IBar bar。
9. BeanPostProcessor.postProcessAfterInitialization(Foo,"foo"). 在上一步中Foo.bar已初始化好了。将Bar注入了IBar bar。

