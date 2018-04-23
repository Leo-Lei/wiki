---
layout: post
title: Dubbo的SPI扩展机制
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---


Dubbo的扩展机制




# 可扩展的几种解决方案

### 工厂
### Java SPI
### Spring等第三方容器 
    
    
# Java SPI机制
Java SPI(Service Provider Interface)是JDK提供的一种加载扩展点的实现。主要是给框架的开发人员来使用。比如，框架提供来一个接口，并允许其他的开发人员提供接口的实现。当服务的提供者提供了一种接口的实现之后，需要在classpath下的META-INF/services/目录里创建一个以服务接口命名的文件，这个文件里的内容就是这个接口的具体的实现类。通过JDK中查找服务的实现的工具类`java.util.ServiceLoader`可以加载具体的实现类。
下面是一个使用Java SPI的例子:
1. 定义一个接口IRepository用于实现数据储存       
```java
public interface IRepository {
    void save(String data);
}
```
实际场景中，IRepository接口一般是在框架中定义的。框架本身可以不提供接口的实现。
2. 提供IRepository的实现       
IRepository有两个实现。MysqlRepository和MongoRepository。
```java
public class MysqlRepository implements IRepository {
    public void save(String data) {
        System.out.println("Save " + data + " to Mysql");
    }
}
```
```java
public class MongoRepository implements IRepository {
    public void save(String data) {
        System.out.println("Save " + data + " to Mongo");
    }
}
```
实际场景中，MysqlRepository和MongoRepository是框架的使用者去实现的。这两个类和IRepository接口是属于不同的Jar包中的。
3. 在`/src/main/resources/META-INF.services`中配置服务的实现
添加文件`src/main/resources/META-INF/services/com.foo.IRepository`
```text
#English implementation
com.bar.MongoRepository

#Chinese implementation
com.bar.MysqlRepository
```
4. 通过ServiceLoader加载IRepository实现
```java
ServiceLoader<IRepository> serviceLoader = ServiceLoader.load(IRepository.class);
Iterator<IRepository> it = serviceLoader.iterator();
while (it != null && it.hasNext()){
    IRepository demoService = it.next();
    System.out.println("class:" + demoService.getClass().getName());
    demoService.save("tom");
}
```

    对SPI的概念有个了解。可以把API和SPI做个对比。
    SPI机制的缺陷
# dubbo的SPI机制

# Dubbo Extension Loader
    ExtentionLoader源码解读
# Dubbo的LoadBalance扩展点解读
    Dubbo中的LoadBalance也是一个SPI，结合源码，分析LoadBalance是如何被加载的
# 自定义一个LoadBalance扩展
    演示如何自己实现一个LoadBbalance，在不改变dubbo源码的情况下，让Dubbo使用我们自定义的LoadBalance实现
# Dubbo SPI高级用法之IoC
   AdaptiveInstance
# Dubbo SPI高级用法之AoP
   wrapper
# Dubbo SPI核心剥离：https://github.com/alibaba/cooma

