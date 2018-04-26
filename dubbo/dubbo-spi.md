---
layout: post
title: Dubbo的SPI扩展机制
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---


# Dubbo的扩展机制
在Dubbo的官网上，Dubbo是这样描述自己的:
> Apache Dubbo (incubating) |ˈdʌbəʊ| is a high-performance, java based RPC framework open-sourced by Alibaba.

提到了Dubbo是一个高性能的，基于Java的开源RPC框架。作为一款RPC框架，高性能是首先要保证的，而作为一个良好的框架，除了性能之外，还有诸如高可用，可扩展，可测试等特性。今天我们就来谈谈可扩展性。在某些时候，软件设计者对扩展性的追求甚至超过了性能。        
在谈到软件设计时，可扩展性一直被谈起，那到底什么才是可扩展性，什么样的框架才算有良好的可扩展性呢？以下两点是我对可扩展性的理解:
1. 作为框架的维护者，在添加一个新功能时，只需要添加一些新代码，而不用大量的修改现有的代码，即符合开闭原则。    
2. 作为框架的使用者，在添加一个新功能时，不需要去修改框架的源码，在自己的工程中添加代码即可。    

上面两点，从白盒和黑盒方面，说明了什么是可扩展。                    
Dubbo很好的做到了上面两点。这要得益于Dubbo的微内核+插件的机制。接下来的章节中我们会慢慢揭开Dubbo扩展机制的神秘面纱。    

# 可扩展的几种解决方案
通常可扩展的实现由下面几种:
* Factory模式
* Ioc容器
* OSGI容器

Dubbo作为一个框架，不希望强依赖其他的IoC容器，比如Spring，Guice。OSGI也是一个很重的实现，不适合Dubbo。最终Dubbo的实现参考了Java原生的SPI机制，但对其进行了一些扩展，以满足Dubbo的需求。

# Java SPI机制
既然Dubbo的扩展机制是基于Java原生的SPI机制，那么我们就先来了解下Java SPI吧。了解了Java的SPI，也就是对Dubbo的扩展机制有一个基本的了解。如果对Java SPI比较了解的同学，可以跳过。

Java SPI(Service Provider Interface)是JDK内置的一种动态加载扩展点的实现。在ClassPath的`META-INF/services`目录下放置一个与接口同名的文本文件，文件的内容为接口的实现类，多个实现类用换行符分隔。JDK中使用`java.util.ServiceLoader`来加载具体的实现。         
让我们通过一个简单的例子，来看看Java SPI是如何工作的。

1. 定义一个接口IRepository用于实现数据储存       
```java
public interface IRepository {
    void save(String data);
}
```
2. 提供IRepository的Mysql实现       
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
3. 添加配置文件
在`META-INF/services`目录添加一个文件，文件名和接口全名称相同，所以文件是`META-INF/services/com.demo.IRepository`。文件内容为:
```text
com.demo.MongoRepository
com.demo.MysqlRepository
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
在上面的例子中，我们定义了一个扩展点和它的两个实现。在ClassPath中添加了扩展的配置文件，最后使用ServiceLoader来加载所有的扩展点。

# dubbo的SPI机制

Java SPI的使用很简单。也做到了基本的加载扩展点的功能。但Java SPI有以下的不足:    
* 需要遍历所有的实现，并实例化，然后我们在循环中才能找到我们需要的实现。
* 配置文件中只是简单的列出了所有的扩展实现，而没有给他们命名。导致在程序中很难去准确的引用它们。
* 扩展如果依赖其他的扩展，做不到自动注入和装配
* 不提供类似于Spring的AOP功能
* 扩展很难和其他的框架集成，比如扩展里面依赖了一个Spring bean，原生的Java SPI不支持

所以Java SPI应付一些简单的场景是可以的，但对于Dubbo，它的功能还是比较弱的。所以，Dubbo对原生SPI机制进行了一些扩展。接下来，我们就更深入地了解下Dubbo的SPI机制。 

# Dubbo扩展点机制基本概念
在深入学习Dubbo的扩展机制之前，我们先明确Dubbo SPI中的一些基本概念。在接下来的内容中，我们会多次用到这些术语。
### 一些术语
1. 扩展点(Extension Point)    
是一个Java的接口。
2. 扩展(Extension)
扩展点的实现类。
3. 扩展实例(Extension Instance)
扩展点实现类的实例。
4. 扩展自适应实例(Extension Adaptive Instance)    
第一次接触这个概念时，可能不太好理解(我第一次也是这样的...)。如果称它为扩展代理类，可能更好理解些。扩展的自适应实例其实就是一个Extension的代理，它实现了扩展点接口。在调用扩展点的接口方法时，会根据实际的参数来决定要使用哪个扩展。比如一个IRepository的扩展点，有一个save方法。有两个实现MysqlRepository和MongoRepository。IRepository的自适应实例在调用接口方法的时候，会根据save方法中的参数，来决定要调用哪个IRepository的实现。如果方法参数中有repository=mysql，那么就调用MysqlRepository的save方法。如果repository=mongo，就调用MongoRepository的save方法。和面向对象的延迟绑定很类似。为什么Dubbo会引入扩展自适应实例的概念呢？              
* Dubbo中的配置有两种，一种是固定的系统级别的配置，在Dubbo启动之后就不会再改了。还有一种是运行时的配置，可能对于每一次的RPC，这些配置都不同。比如在xml文件中配置了超时时间是10秒钟，这个配置在Dubbo启动之后，就不会改变了。但针对某一次的RPC调用，可以设置它的超时时间是30秒钟，以覆盖系统级别的配置。对于Dubbo而言，每一次的RPC调用的参数都是未知的。只有在运行时，根据这些参数才能做出正确的决定。
* 很多时候，我们的类都是一个单例的，比如Spring的bean，在Spring bean都实例化时，如果它依赖某个扩展点，但是在bean实例化时，是不知道究竟该使用哪个具体的扩展实现的。这时候就需要一个代理模式了，它实现了扩展点接口，方法内部可以根据运行时参数，动态的选择合适的扩展实现。而这个代理就是自适应实例。        
自适应扩展实例在Dubbo中的使用非常广泛，Dubbo中，每一个扩展都会有一个自适应类，如果我们没有提供，Dubbo会使用字节码工具为我们自动生成一个。所以我们基本感觉不到自适应类的存在。后面会有例子说明自适应类是怎么工作的。            
5. @SPI    
@SPI注解作用于扩展点的接口上，表明该接口是一个扩展点。可以被Dubbo的ExtentionLoader加载。如果没有此ExtensionLoader调用会异常。
6. @Adaptive
@Adaptive注解用在扩展接口的方法上。表示该方法是一个自适应方法。Dubbo在为扩展点生成自适应实例时，如果方法有@Adaptive注解，会为该方法生成对应的代码。方法内部会根据方法的参数，来决定使用哪个扩展。
7. ExtentionLoader    
类似于Java SPI的ServiceLoader，负责扩展的加载和生命周期维护。
8. 扩展别名
和Java SPI不同，Dubbo中的扩展都有一个别名，用于在应用中引用它们。比如
```text
random=com.alibaba.dubbo.rpc.cluster.loadbalance.RandomLoadBalance
roundrobin=com.alibaba.dubbo.rpc.cluster.loadbalance.RoundRobinLoadBalance
```
其中的random，roundrobin就是对应扩展的别名。这样我们在配置文件中使用random或roundrobin就可以了。

### 一些路径
和Java SPI从`/META-INF/services`目录加载扩展配置类似，Dubbo也会从以下路径去加载扩展配置文件:
* `META-INF/dubbo/internal`
* `META-INF/dubbo`
* `META-INF/services`

# Dubbo的LoadBalance扩展点解读
在了解了Dubbo的一些基本概念后，让我们一起来看一个Dubbo中实际的扩展点，对这些概念有一个更直观的认识。

我们选择的是Dubbo中的LoadBalance扩展点。Dubbo中的一个服务，通常有多个Provider，consumer调用服务时，需要在多个Provider中选择一个。这就是一个LoadBalance。我们一起来看看在Dubbo中，LoadBalance是如何成为一个扩展点的。        
### LoadBalance接口
```java
@SPI(RandomLoadBalance.NAME)
public interface LoadBalance {

    @Adaptive("loadbalance")
    <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException;
}
```
LoadBalance接口只有一个select方法。select方法从多个invoker中选择其中一个。上面代码中和Dubbo SPI相关的元素有:    
* @SPI(RandomLoadBalance.NAME)        
@SPI作用于LoadBalance接口，表示接口LoadBalance是一个扩展点。如果没有@SPI注解，试图去加载扩展时，会抛出异常。@SPI注解有一个参数，该参数表示该扩展点的默认实现的别名。如果没有显示的指定扩展，就使用默认实现。`RandomLoadBalance.NAME`是一个常量，值是"random"，是一个随机负载均衡的实现。    
random的定义在配置文件`META-INF/dubbo/internal/com.alibaba.dubbo.rpc.cluster.LoadBalance`中:
```text
random=com.alibaba.dubbo.rpc.cluster.loadbalance.RandomLoadBalance
roundrobin=com.alibaba.dubbo.rpc.cluster.loadbalance.RoundRobinLoadBalance
leastactive=com.alibaba.dubbo.rpc.cluster.loadbalance.LeastActiveLoadBalance
consistenthash=com.alibaba.dubbo.rpc.cluster.loadbalance.ConsistentHashLoadBalance
```
可以看到文件中定义了4个LoadBalance的扩展实现。由于负载均衡的实现不是本次的内容，这里就不过多说明。只用知道Dubbo提供了4种负载均衡的实现，我们可以通过xml文件，properties文件，JVM参数显式的指定一个实现。如果没有，默认使用随机。                
![dubbo-loadbalance](https://raw.githubusercontent.com/vangoleo/wiki/master/dubbo/dubbo_loadbalance.png)
* @Adaptive("loadbalance")
@Adaptive注解修饰select方法，表明方法select方法是一个可自适应的方法。Dubbo会自动生成该方法对应的代码。当调用select方法时，会根据具体的方法参数来决定调用哪个扩展实现的select方法。@Adaptive注解的参数`loadbalance`表示方法参数中的loadbalance的值作为实际要调用的扩展实例。        
但奇怪的是，我们发现select的方法中并没有loadbalance参数，那怎么获取loadbalance的值呢？select方法中还有一个URL类型的参数，Dubbo就是从URL中获取loadbalance的值的。这里涉及到Dubbo的URL总线模式，简单说，URL中包含了RPC调用中的所有参数。URL类中有一个`Map<String, String> parameters`字段，parameters中就包含了loadbalance。        
### 获取LoadBalance扩展
Dubbo中获取LoadBalance的代码如下:
```java
LoadBalance lb = ExtensionLoader.getExtensionLoader(LoadBalance.class).getExtension(loadbalanceName);
```
使用ExtensionLoader.getExtensionLoader(LoadBalance.class)方法获取一个ExtensionLoader的实例，然后调用getExtension，传入一个扩展的别名来获取对应的扩展实例。

# 自定义一个LoadBalance扩展
Dubbo的4种负载均衡的实现，大多数情况下能满足要求。有时候，因为业务的需要，我们可能需要实现自己的负载均衡策略。    
下面，我们通过一个简单的例子，来自己实现一个LoadBalance，来感受下Dubbo的扩展机制。
1. 实现LoadBalance接口
首先，编写一个自己实现的LoadBalance，因为这里主要是演示Dubbo的扩展机制，而不是LoadBalance的实现，所有这里LoadBalance的实现很简单，会选择第一个invoker，并在控制台输出一条日志。
```java
package com.leibangzhu.test.dubbo.consumer;
public class DemoLoadBalance implements LoadBalance {
    @Override
    public <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException {
        System.out.println("Select the first invoker...");
        return invokers.get(0);
    }
}
```
2. 添加资源文件
添加文件:`src/main/resource/META-INF/dubbo/com.alibaba.dubbo.rpc.cluster.LoadBalance`。文件内容如下:
```text
demo=com.leibangzhu.test.dubbo.consumer.MyLoadBalance
```
3. 配置使用自定义LoadBalance
通过上面的两步，已经添加了一个名字为demo的LoadBalance实现，并在Dubbo中进行来注册。接下来，需要显式的告诉Dubbo使用这个demo的负载均衡实现。如果是通过spring的方式使用Dubbo，可以在xml文件中进行设置。
```xml
<dubbo:reference id="helloService" interface="com.leibangzhu.test.dubbo.api.IHelloService" loadbalance="demo" />
```
在consumer端的<dubbo:reference>中配置<loadbalance="demo">
4. 启动Dubbo进行测试    
启动Dubbo，调用一次IHelloService，可以看到控制台会输出一条`Select the first invoker...`日志。说明Dubbo的确是使用了我们自定义的LoadBalance。      
整个过程会发现：
* 没有改动Dubbo的源码
* 新添加的LoadBalane实现类DemoLoadBalance就是一个简单的Java类，除了实现LoadBalane接口，没有引入其他的元素。对代码的侵入性几乎为零
* 将DemoLoadBalane注册到Dubbo中，只需要添加配置文件`src/main/resources/com.alibaba.dubbo.rpc.cluster.LoadBalance`即可，使用简单。而且不会对现有代码造成影响。符合开闭原则。
    
# Dubbo Extension Loader
是不是觉得Dubbo的扩展机制很不错呀，接下来，我们就打开Dubbo的源码，仔细观摩一番。        
DubboExtentionLoader是一个核心的类，加载扩展点的实现都在这个类中。我们就以这个类开始吧。    
extensionLoader的方法比较多，我先列出ExtensionLoader使用的方式吧：
```java
LoadBalance lb = ExtensionLoader.getExtensionLoader(LoadBalance.class).getExtension(loadbalanceName);
```
首先看getExtensionLoader方法，这是一个静态工厂方法，入参是一个可扩展的接口，返回一个该接口的ExtensionLoader实体类。
```java
public static <T> ExtensionLoader<T> getExtensionLoader(Class<T> type)
```
再来看看getExtension方法
```java
public T getExtension(String name)
```

```java
public static <T> ExtensionLoader<T> getExtensionLoader(Class<T> type) {
        if (type == null)
            throw new IllegalArgumentException("Extension type == null");
        // 扩展点必须是接口
        if (!type.isInterface()) {
            throw new IllegalArgumentException("Extension type(" + type + ") is not interface!");
        }
        // 必须要有@SPI注解
        if (!withExtensionAnnotation(type)) {
            throw new IllegalArgumentException("Extension type(" + type +
                    ") is not extension, because WITHOUT @" + SPI.class.getSimpleName() + " Annotation!");
        }
        // 从缓存中根据接口获取对应的ExtensionLoader
        // 每个扩展只会被加载一次
        ExtensionLoader<T> loader = (ExtensionLoader<T>) EXTENSION_LOADERS.get(type);
        if (loader == null) {
            // 初始化扩展
            EXTENSION_LOADERS.putIfAbsent(type, new ExtensionLoader<T>(type));
            loader = (ExtensionLoader<T>) EXTENSION_LOADERS.get(type);
        }
        return loader;
    }
    
private ExtensionLoader(Class<?> type) {
        this.type = type;
        objectFactory = (type == ExtensionFactory.class ? null : ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension());
    }
```

再来看看getExtension方法
```java
public T getExtension(String name) {
        // null判断
        if (name == null || name.length() == 0)
            throw new IllegalArgumentException("Extension name == null");
        // 如果name="true",返回默认的Extention。这个逻辑先不用关心，不影响主流程
        if ("true".equals(name)) {
            return getDefaultExtension();
        }
        Holder<Object> holder = cachedInstances.get(name);
        if (holder == null) {
            cachedInstances.putIfAbsent(name, new Holder<Object>());
            holder = cachedInstances.get(name);
        }
        Object instance = holder.get();
        // 从缓存中获取，如果不存在就创建
        if (instance == null) {
            synchronized (holder) {
                instance = holder.get();
                if (instance == null) {
                    instance = createExtension(name);
                    holder.set(instance);
                }
            }
        }
        return (T) instance;
    }
```
getExtention方法中做了一些判断和缓存，主要的逻辑在createExtension方法中。我们继续看createExtention方法。
createExtension方法会做以下事情:
1. 先根据name来得到扩展类。从ClassPath下`META-INF`文件夹下读取扩展点配置文件    
2. 使用反射创建一个扩展类的实例    
3. 对扩展类的实例进行依赖注入，即常说的IoC    
4. 如果有wrapper，添加wrapper。即常说的AoP    
```java
private T createExtension(String name) {
        // 根据扩展点名称得到扩展类，比如对于LoadBalance，根据random得到RandomLoadBalance类
        Class<?> clazz = getExtensionClasses().get(name);
        if (clazz == null) {
            throw findException(name);
        }
        try {
            T instance = (T) EXTENSION_INSTANCES.get(clazz);
            if (instance == null) {
                // 使用反射调用nesInstance来创建扩展类的一个示例
                EXTENSION_INSTANCES.putIfAbsent(clazz, (T) clazz.newInstance());
                instance = (T) EXTENSION_INSTANCES.get(clazz);
            }
            // 对扩展类示例进行依赖注入
            injectExtension(instance);
            // 如果有wrapper，添加wrapper
            Set<Class<?>> wrapperClasses = cachedWrapperClasses;
            if (wrapperClasses != null && !wrapperClasses.isEmpty()) {
                for (Class<?> wrapperClass : wrapperClasses) {
                    instance = injectExtension((T) wrapperClass.getConstructor(type).newInstance(instance));
                }
            }
            return instance;
        } catch (Throwable t) {
            throw new IllegalStateException("Extension instance(name: " + name + ", class: " +
                    type + ")  could not be instantiated: " + t.getMessage(), t);
        }
    }
```
getExtensionClasses会根据扩展名得到对应的扩展类。也是先从缓存中获取，如果没有，就从CLASSPATH中加载文件。    
Dubbo会从以下的CLASSPATH路径去加载扩展点文件：
* `META-INF/dubbo/internal`
* `META-INF/dubbo`
* `META-INF/services`

```java
private Map<String, Class<?>> getExtensionClasses() {
        Map<String, Class<?>> classes = cachedClasses.get();
        if (classes == null) {
            synchronized (cachedClasses) {
                classes = cachedClasses.get();
                if (classes == null) {
                    classes = loadExtensionClasses();
                    cachedClasses.set(classes);
                }
            }
        }
        return classes;
    }

    // synchronized in getExtensionClasses
    private Map<String, Class<?>> loadExtensionClasses() {
        final SPI defaultAnnotation = type.getAnnotation(SPI.class);
        if (defaultAnnotation != null) {
            String value = defaultAnnotation.value();
            if (value != null && (value = value.trim()).length() > 0) {
                String[] names = NAME_SEPARATOR.split(value);
                if (names.length > 1) {
                    throw new IllegalStateException("more than 1 default extension name on extension " + type.getName());
                }
                if (names.length == 1) cachedDefaultName = names[0];
            }
        }

        Map<String, Class<?>> extensionClasses = new HashMap<String, Class<?>>();
        loadFile(extensionClasses, DUBBO_INTERNAL_DIRECTORY);
        loadFile(extensionClasses, DUBBO_DIRECTORY);
        loadFile(extensionClasses, SERVICES_DIRECTORY);
        return extensionClasses;
    }
```

injectExtension方法实现了Ioc即自动装配的功能。WrapperClass实现了AOP功能。这两点，我们会在后面的章节中单独来说明。到此为止，Dubbo加载和创建扩展的主流程就结束了。总结下来就是加载ClassPath的META-INF目录下的一些文件，再根据name来创建对应的扩展类实例，并对其进行依赖注入和自动包装。        

# Dubbo SPI高级用法之IoC
看看injectExtension方法。我已经把一些无关的代码去掉了。比如日志，不影响主流程的异常处理等。希望大家可以更专注在核心代码上。
```java
private T injectExtension(T instance) {
    for (Method method : instance.getClass().getMethods()) {
        if (method.getName().startsWith("set")
                && method.getParameterTypes().length == 1
                && Modifier.isPublic(method.getModifiers())) {
            Class<?> pt = method.getParameterTypes()[0];
          
            String property = method.getName().length() > 3 ? method.getName().substring(3, 4).toLowerCase() + method.getName().substring(4) : "";
            Object object = objectFactory.getExtension(pt, property);
            if (object != null) {
                method.invoke(instance, object);
            }
        }
    }
    return instance;
}
```
injectExtension方法遍历扩展类的所有set方法，通过set方法进行注入。比如扩展类是A，A有一个字段是B b和一个setB的方法。injectExtension方法会创建一个B的实例，并把它注入到A中，完成依赖注入的过程。这里面有个问题是如果获取B的实例。因为B的实例有可能是一个简单的POJO，也可能是另一个Dubbo的SPI扩展，也可能是一个Spring的Bean，或其他更复杂的情况。那Dubbo是怎么根据类B来获取正确的B的实例呢？    
具体的实现是这句代码`Object object = objectFactory.getExtension(pt, property);`。objectFactory是ExtensionFactory类型的，在创建ExtensionLoader时被初始化了。    
```java
private ExtensionLoader(Class<?> type) {
        this.type = type;
        objectFactory = (type == ExtensionFactory.class ? null : ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension());
    }
```

![Dubbo-ExtensionFactory](https://raw.githubusercontent.com/vangoleo/wiki/master/dubbo/dubbo-extensionfactory.png)
ExtensionLoader有三个实现：
1. SpiExtensionLoader：Dubbo自己的Spi去加载Extension
2. SpringExtensionLoader：从Spring容器中去加载Extension
3. AdaptiveExtensionLoader: 自适应的AdaptiveExtensionLoader
这里要注意AdaptiveExtensionLoader。    
```java
@Adaptive
public class AdaptiveExtensionFactory implements ExtensionFactory {

    private final List<ExtensionFactory> factories;

    public AdaptiveExtensionFactory() {
        ExtensionLoader<ExtensionFactory> loader = ExtensionLoader.getExtensionLoader(ExtensionFactory.class);
        List<ExtensionFactory> list = new ArrayList<ExtensionFactory>();
        for (String name : loader.getSupportedExtensions()) {
            list.add(loader.getExtension(name));
        }
        factories = Collections.unmodifiableList(list);
    }

    public <T> T getExtension(Class<T> type, String name) {
        for (ExtensionFactory factory : factories) {
            T extension = factory.getExtension(type, name);
            if (extension != null) {
                return extension;
            }
        }
        return null;
    }
}
```
AdaptiveExtensionLoader添加了@Adaptive注解，所以会使用它作为ExtensionLoader。AdaptiveExtentionLoader里面，会遍历所有的ExtensionLoader实现，去加载Extension。也就是它会遍历SpiExtensionFactory和SpringExtensionFactory，去加载Extension。SpiExtensionLoader或调用ExtensionLoader来加载Extension。SpringExtensionFactory会从Spring ApplicationContext中去获取bean，作为Extension。如果我们自己也实现了一个ExtensionLoader，Dubbo也会从它去加载extension。通过这种机制，让Dubbo的依赖注入支持了Dubbo的SPI，Spring的Bean，和自定义的实现。
# Dubbo SPI高级用法之AoP
在用Spring的时候，我们经常会用到AOP功能。在目标类的方法前后插入其他逻辑。比如通常使用Spring AOP来实现日志和鉴权等逻辑。    
那么在Dubbo的SPI体系中，是否也有类似的功能呢？答案是有的。在Dubbo中，有一种特殊的类，被称为Wrapper类。通过装饰者模式，使用包装类包装原始的扩展点实例。在原始扩展点实现前后插入其他逻辑，实现AOP功能。    

### 什么是Wrapper类
那什么样类的才是Dubbo SPI中的Wrapper类呢？Dubbo SPI中的Wrapper类就是一个有复制构造函数的类，也是典型的装饰者模式。下面就是一个Wrapper类，有一个复制构造函数`public A(A a)`。
```java
class A{
    private A a;
    public A(A a){
        this.a = a;
    }
}
```
Dubbo中这样的Wrapper类有ProtocolFilterWrapper, ProtocolListenerWrapper等。
### 怎么配置Wrapper类
在Dubbo中Wrapper类也是一个Extension，和其他的Extension一样，也是在Classpath的META-INF文件夹中配置的。比如前面举例的ProtocolFilterWrapper和ProtocolListenerWrapper就是在路径`dubbo-rpc/dubbo-rpc-api/src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.Protocol`中配置的。
```text
filter=com.alibaba.dubbo.rpc.protocol.ProtocolFilterWrapper
listener=com.alibaba.dubbo.rpc.protocol.ProtocolListenerWrapper
mock=com.alibaba.dubbo.rpc.support.MockProtocol
```
ExtensionLoader在创建Protocol的ExtensionLoader实例时，会加载这些文件
```java
loadFile(extensionClasses, DUBBO_INTERNAL_DIRECTORY);  
loadFile(extensionClasses, DUBBO_DIRECTORY);  
loadFile(extensionClasses, SERVICES_DIRECTORY);  
```
在loadFile中有一段如下的代码:
```java
try {  
  clazz.getConstructor(type);    
  Set<Class<?>> wrappers = cachedWrapperClasses;
  if (wrappers == null) {
    cachedWrapperClasses = new ConcurrentHashSet<Class<?>>();
    wrappers = cachedWrapperClasses;
  }
  wrappers.add(clazz);
} catch (NoSuchMethodException e) {}
```
这段代码的意思是，如果类有复制构造函数，就把该类存起来，供以后使用。有复制构造函数的类就是Wrapper类。
获取参数类型是type的构造函数。注意其中type是扩展点接口，而不是具体的扩展类。
以Protocol为例。Protocol是一个@SPI修饰的接口，是Dubbo里的一个扩展点。ExtensionLoader.getExtensionLoader(Protocol.class)时，ExtensionLoader会遍历ClassPath加载文件，当读取到`dubbo-rpc/dubbo-rpc-api/src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.Protocol`时，该文件中定义了`filter=com.alibaba.dubbo.rpc.protocol.ProtocolFilterWrapper`。    
ProtocolFilterWrapper定义如下:
```java
public class ProtocolFilterWrapper implements Protocol {

    private final Protocol protocol;

    // 有一个参数是Protocol的复制构造函数
    public ProtocolFilterWrapper(Protocol protocol) {
        if (protocol == null) {
            throw new IllegalArgumentException("protocol == null");
        }
        this.protocol = protocol;
    }
```
ProtocolFilterWrapper有一个构造函数`public ProtocolFilterWrapper(Protocol protocol)`，参数是扩展点Protocol，所以它是一个Wrapper类。ExtensionLoader会把它缓存起来，供以后创建Extension实例的时候，使用这些包装类依次包装原始扩展点。

```java
private T createExtension(String name) {
        Class<?> clazz = getExtensionClasses().get(name);
        if (clazz == null) {
            throw findException(name);
        }
        try {
            T instance = (T) EXTENSION_INSTANCES.get(clazz);
            if (instance == null) {
                EXTENSION_INSTANCES.putIfAbsent(clazz, (T) clazz.newInstance());
                instance = (T) EXTENSION_INSTANCES.get(clazz);
            }
            injectExtension(instance);
            Set<Class<?>> wrapperClasses = cachedWrapperClasses;
            if (wrapperClasses != null && !wrapperClasses.isEmpty()) {
                for (Class<?> wrapperClass : wrapperClasses) {
                    instance = injectExtension((T) wrapperClass.getConstructor(type).newInstance(instance));
                }
            }
            return instance;
        } catch (Throwable t) {
            throw new IllegalStateException("Extension instance(name: " + name + ", class: " +
                    type + ")  could not be instantiated: " + t.getMessage(), t);
        }
    }
```

# 扩展点自适应
在Dubbo中，一个扩展点会有多个实现。Dubbo究竟选择哪个实现呢？这取决于系统的配置。
1. 在Dubbo配置文件中进行配置，属于系统级别配置。应用启动时会读取该配置
2. 每一次的rpc调用中，可以设置参数来覆盖系统级别的配置
大多数时候，我们都使用系统级别的配置。一旦应用起来后，这些配置就不会改了。比如序列化方式，超时时间等。但有些时候，我们需要动态地修改一些配置，比如对某些RPC调用使用不同的负载均衡策略等。这就要求Dubbo的SPI机制可以根据参数，自动的选择对应的实现。将扩展点实例的选择从应用启动的时候延迟到实际调用的时候。    
Dubbo使用适配器类来实现扩展点自适应功能。在需要实现自适应的接口方法上使用@Adaptive注解。框架会通过字节码工具自动创建一个适配类，默认是Javaassist。
创建自适应扩展类的代码:
```java
public T getAdaptiveExtension() {
    Object instance = cachedAdaptiveInstance.get();
    if (instance == null) {
            synchronized (cachedAdaptiveInstance) {
                instance = cachedAdaptiveInstance.get();
                if (instance == null) {
                      instance = createAdaptiveExtension();
                      cachedAdaptiveInstance.set(instance); 
                }
            }        
    }

    return (T) instance;
}
```
createAdaptiveExtension方法
```java
private T createAdaptiveExtension() {        
    return injectExtension((T) getAdaptiveExtensionClass().newInstance());
}
```
getAdaptiveExtensionClass方法
```java
private Class<?> getAdaptiveExtensionClass() {
        getExtensionClasses();
        if (cachedAdaptiveClass != null) {
            return cachedAdaptiveClass;
        }
        return cachedAdaptiveClass = createAdaptiveExtensionClass();
    }
```
绕了一大圈，终于来到了具体的实现了。看这个createAdaptiveExtensionClass方法，它首先会生成自适应类的Java源码，然后再将源码编译成Java的字节码，加载到JVM中。
```java
private Class<?> createAdaptiveExtensionClass() {
        String code = createAdaptiveExtensionClassCode();
        ClassLoader classLoader = findClassLoader();
        com.alibaba.dubbo.common.compiler.Compiler compiler = ExtensionLoader.getExtensionLoader(com.alibaba.dubbo.common.compiler.Compiler.class).getAdaptiveExtension();
        return compiler.compile(code, classLoader);
    }
```
Compiler的代码，默认实现是javassist。
```java
@SPI("javassist")
public interface Compiler {

    Class<?> compile(String code, ClassLoader classLoader);
}

```

createAdaptiveExtensionClassCode()方法中使用一个StringBuilder来构建自适应类的Java源码。方法实现比较长，这里就不贴代码了。下面是使用createAdaptiveExtensionClassCode方法为Protocol创建的自适应类的Java代码:
```java
package com.alibaba.dubbo.rpc;

import com.alibaba.dubbo.common.extension.ExtensionLoader;

public class Protocol$Adpative implements com.alibaba.dubbo.rpc.Protocol {
    public void destroy() {
        throw new UnsupportedOperationException("method public abstract void com.alibaba.dubbo.rpc.Protocol.destroy() of interface com.alibaba.dubbo.rpc.Protocol is not adaptive method!");
    }

    public int getDefaultPort() {
        throw new UnsupportedOperationException("method public abstract int com.alibaba.dubbo.rpc.Protocol.getDefaultPort() of interface com.alibaba.dubbo.rpc.Protocol is not adaptive method!");
    }

    public com.alibaba.dubbo.rpc.Exporter export(com.alibaba.dubbo.rpc.Invoker arg0) throws com.alibaba.dubbo.rpc.RpcException {
        if (arg0 == null) throw new IllegalArgumentException("com.alibaba.dubbo.rpc.Invoker argument == null");
        if (arg0.getUrl() == null)
            throw new IllegalArgumentException("com.alibaba.dubbo.rpc.Invoker argument getUrl() == null");
        com.alibaba.dubbo.common.URL url = arg0.getUrl();
        String extName = (url.getProtocol() == null ? "dubbo" : url.getProtocol());
        if (extName == null)
            throw new IllegalStateException("Fail to get extension(com.alibaba.dubbo.rpc.Protocol) name from url(" + url.toString() + ") use keys([protocol])");
        com.alibaba.dubbo.rpc.Protocol extension = (com.alibaba.dubbo.rpc.Protocol) ExtensionLoader.getExtensionLoader(com.alibaba.dubbo.rpc.Protocol.class).getExtension(extName);
        return extension.export(arg0);
    }

    public com.alibaba.dubbo.rpc.Invoker refer(java.lang.Class arg0, com.alibaba.dubbo.common.URL arg1) throws com.alibaba.dubbo.rpc.RpcException {
        if (arg1 == null) throw new IllegalArgumentException("url == null");
        com.alibaba.dubbo.common.URL url = arg1;
        String extName = (url.getProtocol() == null ? "dubbo" : url.getProtocol());
        if (extName == null)
            throw new IllegalStateException("Fail to get extension(com.alibaba.dubbo.rpc.Protocol) name from url(" + url.toString() + ") use keys([protocol])");
        com.alibaba.dubbo.rpc.Protocol extension = (com.alibaba.dubbo.rpc.Protocol) ExtensionLoader.getExtensionLoader(com.alibaba.dubbo.rpc.Protocol.class).getExtension(extName);
        return extension.refer(arg0, arg1);
    }
}
```

先看看Protocol的源码
```java
@SPI("dubbo")
public interface Protocol {

    int getDefaultPort();

    @Adaptive
    <T> Exporter<T> export(Invoker<T> invoker) throws RpcException;

    @Adaptive
    <T> Invoker<T> refer(Class<T> type, URL url) throws RpcException;

    void destroy();
```
我们可以看到Prototol接口中有4个方法，但只有export和refer两个方法使用了@Adaptive注解。所以，Dubbo生成的自适应类Protocol$Adpative中，也只有这两个方法有实现，其余方法都是直接跑出异常。    
大致的逻辑和开始说的一样，通过url解析出参数，解析的逻辑由@Adaptive的value参数控制，然后再根据得到的扩展点名获取扩展点实现，然后进行调用。具体拼接逻辑大家可以看createAdaptiveExtensionClassCode的实现
