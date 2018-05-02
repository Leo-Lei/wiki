---
layout: post
title: Dubbo的SPI扩展机制
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---


# Dubbo Extension Mechanism
In Dubbo's offical website, Dubbo describe itself as a Java based, High performance RPC framework. But today, I want to talk about another good feature in Dubbo -- the highly extensible.    

In the world of Software developing, the extensibility is always being talked about。 So what is a high extensible framework? In my option, it should satisfy the following demands:
1. As a framewofk developer, to add some new functionality, just need add some new code, without changing exist code, i.e. open/closed principle.        
2. As a framework user, who use the framework in their project, to add some new functionality, Do not need to change the source code of the project. Just need to add some code in his/her own project.        


Dubbo has already satisfied the above two points very well. In the rest of this article, we will dive into the Dubbo extension mechanism.

# Highly extensibility solution
In general, there are some solutions to achieve high extensibility solution:        
* Factory pattern
* Ioc container, such as Spring, Google guice
* OSGI

As a RPC framework, Dubbo don't want to bring in other IoC container. OSGI is also very heavy technology stack. Dubbo extension is inherited from standard Java SPI(Service Provider Interface) and do some enhancement to make it more powerful. 

# Java SPI机制
Since Dubbo extension is inherited from Java SPI, let's have a look at the Java SPI. If you are familiar with Java SPI, you can skip this section.         

Java SPI is a JDK's build-in mechanism to enable dynamically loading extension at runtime, out of box. In the jar file containing extension class, place a config file `META-INF/dubbo/full_interface_name`, the file contend pattern: `extension_name=the_full_name_of_extension_class`, multiple implementations are separated by new line. The class `java.util.ServiceLoader` is responsible to load extension.    

Let's have a simele sample, to have a quick glance at how Java SPI works.

1. Define an interface IRepository       
```java
public interface IRepository {
    void save(String data);
}
```
2. Provide two implementations of IRepository: MysqlRepository and MongoRepository.       
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
3. Add config file
Put a text file on directory `META-INF/services`, and the file name is the same as the interface, so the file is `META-INF/services/com.demo.IRepository`. The file content is as below: 
```text
com.demo.MongoRepository
com.demo.MysqlRepository
```
4. Use ServiceLoader to load IRepository
```java
ServiceLoader<IRepository> serviceLoader = ServiceLoader.load(IRepository.class);
Iterator<IRepository> it = serviceLoader.iterator();
while (it != null && it.hasNext()){
    IRepository demoService = it.next();
    System.out.println("class:" + demoService.getClass().getName());
    demoService.save("tom");
}
```
In the ablove sample, we have defined an extension point(the IRepository interface) and two implementations, add a config file and use ServiceLoader to load concrete implementation. 

# dubbo SPI machanism
It's easy to use, and can satisify basic extensibility requirement. but it has some disadvantage:    
* Must iterate all the extension implementations, to find the exact one that we need. All the implementations need to be load and initialized, eventhough it will never be used at all.
* In the extension configuration file, extention implementation doesn't have a name, which makes it is difficult to refer to them in our programe.
* If a extension has some other extension as its dependency, Java SPI do not support auto depencency injection.
* Do not support AOP. Not support add some common logic to batch of extensions, rather than explicitly add it to each specific extension. 
* Do not intergrate with other framework well. Such as the extension can not have a dependency of a spring bean.

So, in some simple scenario, Java SPI is competent. But in dubbo, the extension loading and management is very complex. Dubbo enhance native Java SPI to have more power. 

# Basic concept in Dubbo extensibility
Before we dive into Dubbo extensibility, there are some basic concept that we need to know:
### Extension Point
A java interface
### Extension
A class which implementation of Extension point
3. Extension Instance
A instance of Extension class
4. Extension Adaptive Instance
It may be difficult to understand, if you meet it the at first time(both for me...). Maybe a "extension proxy" name is much easier to understand. In fact, the extension adaptive instance is an proxy of the extention point. In code level, it implements the interface of extension point. When invoke the interface method, it will get the parameter value and then decide which specific implementation should be selected to be invoked.    
For example, there is an extention interface `IRepository` with a `save` method. The `IRepository` has two implementations, MysqlRepository and MongoRepository. In addition, there is an adaptive instance of IRepository. When invoke the save method of a adaptive instance, it will check the method parameter, if repository=mysql, it will invoke the underlying MysqlRepository.save() method. If repository=mongo, it will invoke MongoRepository.save() method. This act very similar with the lazy-bindin principle.

Why dubbo bring in the adaptive instance?        

* There are two kinds of configuration file in Dubbo, the first is some predefined system level configuration, which will be loaded from config files and will not changed after dubbo started. The second type is the runtime configuration, which is different for each RPC invocation. For example, we set the timeout=10 second in Dubbo xml configuration file. After Dubbo started, it will not changed. For some RPC invocation, if we don't specify the timeout explicitly, the timeout=10 second will be used. When we specify the timeout=30 second for some RPC, the 30 will override the default 10. So in Dubbo, the configuration for each RPC invocation may be different. It can not be fixed previously, until at the runtime. Dubbo need to make the correct decision according to the runtime configuration. The decision contains select the correct extension implementation.

At most time, the service-type class in our project is singleton, for instance, a Spring bean. The extension in Dubbo are also singleton. Image that an extension A has a dependency B, which is also an extension. The A is singleton, but for its dependency B, it can be any implementation of B. So what should the B be? In this case, we need the B to be a proxy of B. which can redirect the request to the correct underlying implementation. The proxy is the adaptive instance.        

The adaptive instance is widely used in Dubbo. In fact, every extention point has a relate adaptive instance. If we don't supply it, Dubbo will auto generate it for us by bytecode genaration tool. In the following section, you will touch how the adaptive instance work.

5. @SPI
The @SPI annotation can be added to an interface, which indicate that the interface is an extention point.
6. @Adaptive
The @Adaptive annotation can be added to a class or a method. A class with a @Adaptive is the adaptive instace of the interface. A method with a @Adaptive will be implemented while Dubbo auto generated the adaptive instance for us.        
7. ExtentionLoader    
Similar with ServiceLoader in Java SPI, and is responsible to load extension and manage the lifecycle of extensions.        
8. extension name(alias)
Different from Java SPI, each extension has a name in Dubbo. For example:    
```text
random=com.alibaba.dubbo.rpc.cluster.loadbalance.RandomLoadBalance
roundrobin=com.alibaba.dubbo.rpc.cluster.loadbalance.RoundRobinLoadBalance
```
The random and the roundrobin is a alias of its implementation. Then we can refer to them by using `random` or `roundrobin`.    

### config file path
Similar to Java SPI load extensionss from `/META-INF/services`, Dubbo will load config files from below directories:    
* `META-INF/dubbo/internal`
* `META-INF/dubbo`
* `META-INF/services`

# Dubbo的LoadBalance扩展点解读
After have some awareness on above basic Dubbo concepts, let's have a look at a read extension example in Dubbo, aimed to have better understanding of Dubbo extensibility mechanism.

I take the LoadBalance for example. A service in Dubbo may has multiple provider, which become a cluster. A consumer need to select one of them to execute the RPC invocation. This is the LoadBalance. Let's go ahead, to investigate how the LoadBalance becomes a extendion id Dubbo.

### LoadBalance interface
```java
@SPI(RandomLoadBalance.NAME)
public interface LoadBalance {

    @Adaptive("loadbalance")
    <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException;
}
```
LoadBalance has only one `select` method. The select method will select one invoker from multiple invokers. The corresponding code of dubbo SPI is as below:        
* @SPI(RandomLoadBalance.NAME)        
The @SPI is added to LoadBalance means the LoadBalance is an extension point. @SPI annotation has a parameter, which indicate the default implementation of the extension point. If we don't specify the implementation explicitly, the default implementation will be selected.
`RandomLoadBalance.NAME` is a constant value "random", which indicate a random strategy implementation of LoadBalance.
The `random` is defined in file `META-INF/dubbo/internal/com.alibaba.dubbo.rpc.cluster.LoadBalance`:
```text
random=com.alibaba.dubbo.rpc.cluster.loadbalance.RandomLoadBalance
roundrobin=com.alibaba.dubbo.rpc.cluster.loadbalance.RoundRobinLoadBalance
leastactive=com.alibaba.dubbo.rpc.cluster.loadbalance.LeastActiveLoadBalance
consistenthash=com.alibaba.dubbo.rpc.cluster.loadbalance.ConsistentHashLoadBalance
```
We can see that there are 4 implementation of LoadBalance. As LoadBalance is not the topic of this articlt, so we will not spend much time on explaining these 4 implementations. At this point, you only need to know that there are 4 LoadBalance implementation, and we can specify which one will be select by XML config file, properties file or JVM arguments. If not specify, the default one will be select.
![dubbo-loadbalance](https://raw.githubusercontent.com/vangoleo/wiki/master/dubbo/dubbo_loadbalance.png)
* @Adaptive("loadbalance")
The @Adaptive annotation is added to select method, which indicate that the select method is an adaptive method. dubbo will create a proxy for this method automatically. While invoking the select method, it will select a proper LoadBalance according to the method parameters.
@Adaptive annotation has one parameter `loadbalance` means the method will judge the value of loadbalance parameter from the method parameters. You may already found that there is no parameter named loadbalance in parameters. How can select method get the value of loadbalance? The answer is the URL parameter. The full name of URL is `com.alibaba.dubbo.common.URL`. Here comes another concept in Dubbo -- the URL Bus pattern. In some simple words, the URL contains all the cofigurations of a RPC invocation. Inner the URL, there is a `Map<String, String> parameters` field. The loadbalance is obtained in it.         
### 获取LoadBalance扩展
In Dubbo, the code of getting LoadBalance is as below:
```java
LoadBalance lb = ExtensionLoader.getExtensionLoader(LoadBalance.class).getExtension(loadbalanceName);
```
Invoke the `ExtensionLoader.getExtensionLoader(LoadBalance.class)` method to get an ExtensionLoader instance, then invoke th getExtension to get a extension instance.

# 自定义一个LoadBalance扩展            
In this section, we will supply a cusomize implementation of LoadBalance, and intergrate it into Dubbo. By this, we can have a better understanding of Dubbo SPI.        
1. LoadBalance implementation
First, we need to write a implementation of LoadBalance. Since our topic is the Dubbo SPI, not the LoadBalance itself. So in this demo, the implementation is very simple, it simply select the first invoker, and print a log to console. 
```java
package com.leibangzhu.test.dubbo.consumer;
public class DemoLoadBalance implements LoadBalance {
    @Override
    public <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException {
        System.out.println("DemoLoadBalance: Select the first invoker...");
        return invokers.get(0);
    }
}
```
2. add config file
add file:`META-INF/dubbo/com.alibaba.dubbo.rpc.cluster.LoadBalance`。the content of file is as below:
```text
demo=com.leibangzhu.test.dubbo.consumer.MyLoadBalance
```
3. configure to use customize LoadBalance
In above steps, we have created a customize LoadBalance and assigned a name `demo` to it. Then, we need to ask Dubbo to use ourown LoadBalance. If we are using Dubbo by xml file, the configuration is as below:
```xml
<dubbo:reference id="helloService" interface="com.leibangzhu.test.dubbo.api.IHelloService" loadbalance="demo" />
```
Add `loadbalance="demo"` to <dubbo:reference> in consumer side.
4. Start Dubbo    
Start Dubbo，invoke IHelloService, then you can find a log of `DemoLoadBalance: Select the first invoker...`。Dubbo already use ourown LoadBalance.     

At this time, the entire demo is completed. We will find that:
* not change dubbo source code
* the DemoLoadBalance is a simple Java class, except for implements the LoadBalance interface. There is no extra concept being involed in DemoLoadBalance.
* To add a new implementation of LoadBalance, just need to create a new class an a new config file. No need to change existing code. This follows the open-closed principle.
    
 
