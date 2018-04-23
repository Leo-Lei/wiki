---
layout: post
title: Dubbo的SPI扩展机制
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---


# Dubbo的扩展机制
Dubbo作为一个通用的RPC框架，被广泛的使用的原因，也是因为它的基于微内核和扩展点的架构设计。使得Dubbo的核心部分比较稳定，通过扩展点来实现丰富的功能。第三方开发者也可以根据需要，自己编写扩展点，在不改变Dubbo源码的情况下集成到Dubbo中。

# 可扩展的几种解决方案

### 工厂
利用工厂模式，可以通过传入不同的参数，来创建出不同的类的实例。但是工厂创建出的这些类是如何被加载的是个问题。如果在工厂中使用new的方式，那么所有可能的类都要预先被定义。这肯定是不行的。也可以使用反射，传入一个类名，根据类名来动态的创建一个类的实例。这可以处理一些简单的场景。如果情况变得复杂了，比如，类之间是有依赖的。实际的项目中通常会有多个扩展点一起配合工作。即一个扩展点中会依赖其他的扩展点。这时候，在工厂里使用反射来获取扩展点就不能满足要求来。

### Spring等第三方容器 
Java中有很多的IoC容器，比如spring，Google的Guice。应该都能够满足我们的要求。但是这些框架往往都很胖，因为它们都是一个公共的基础框架，需要满足各种开发者，在各种场景下的要求，必然会导致本身越来越胖。相信这点，用过Spring的会有体会。在一个框架中强依赖另一个很重的框架，这肯定不是一个好的设计。

### Java SPI    
Java本身其实提供了一个加载扩展点的机制，就是Java SPI。无需引入Spring等第三方框架即可动态的加载扩展点。

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
通过上面的代码可以发现，我们使用SPI查找具体的实现的时候，需要遍历所有的实现，并实例化，然后我们在循环中才能找到我们需要的实现。这是Java原生SPI的缺点，需要把所有的实现都实例化了，即便我们不需要，也都给实例化了。

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

