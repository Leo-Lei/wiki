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
3. 在`/src/main/resources/META-INF/services`中配置服务的实现
添加了MysqlRepository和MongoRepository后，需要告诉Java SPI，我们为IRepository添加了两个实现。需要:
* 在`src/main/resources/META-INF/services`目录添加一个文件
* 文件名就是接口的全名称，即`com.foo.IRepository`
* 文件的内容需要列出所有的接口的实现  

添加文件`src/main/resources/META-INF/services/com.foo.IRepository`
```text
#Mongo implementation
com.bar.MongoRepository

#Mysql implementation
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

Dubbo的扩展点加载机制类似于Java的SPI，在前面的描述中，我们知道了Java的SPI只能通过遍历来进行实现的查找和实例化，有可能会一次性把所有的实现都实例化，这样会造成有些不使用的扩展实现也会被实例化，这就会造成一定的资源浪费。Dubbo对这一点进行了优化。除此之外，Dubbo还进行了其他方面的优化。有关Dubbo的改进，参照文档上的说明:

1. JDK标准的SPI会一次性实例化扩展点所有实现，如果有扩展实现初始化很耗时，但如果没用上也加载，会很浪费资源。    
2. 如果扩展点加载失败，连扩展点的名称都拿不到了。比如：JDK标准的ScriptEngine，通过getName();获取脚本类型的名称，但如果RubyScriptEngine因为所依赖的jruby.jar不存在，导致RubyScriptEngine类加载失败，这个失败原因被吃掉了，和ruby对应不起来，当用户执行ruby脚本时，会报不支持ruby，而不是真正失败的原因。 
3. 增加了对扩展点IoC和AOP的支持，一个扩展点可以直接setter注入其它扩展点。

# Dubbo扩展点机制基本概念
1. Extension Point    
扩展点。是一个Java的接口。
2. Extension    
Extension是扩展点的实现类。
3. Extension Instance    
扩展点实现类的实例。
4. @SPI    
@SPI注解作用于扩展点的接口上，表明该接口是一个扩展点。可以被Dubbo的ExtentionLoader加载。如果没有此ExtensionLoader调用会异常。
5. ExtentionLoader    
负责加载对应的扩展。
6. Extension Adaptive Instance    
扩展的自适应实例。扩展的自适应实例是一个Extension的代理，在调用Extension Adaptive Instance的某个方法时，会根据参数真正决定要调用的那个扩展。
7. src/main/resources/META-INF/dubbo/internal
该目录类似于Java SPI的`META-INF/services`目录。有扩展点的配置文件。格式也和Java SPI的有些类似。

# Dubbo的LoadBalance扩展点解读
Dubbo中的LoadBalance也是一个扩展点，我们可以结合源码，分析LoadBalance是如何被加载的。
1. LoadBalance扩展点
Dubbo中负载均衡的扩展点是LoadBalance接口。
```java
@SPI(RandomLoadBalance.NAME)
public interface LoadBalance {

    @Adaptive("loadbalance")
    <T> Invoker<T> select(List<Invoker<T>> invokers, URL url, Invocation invocation) throws RpcException;
}
```
LoadBalance接口只有一个select方法。select方法从多个invoker中选择其中一个。上面的代码中需要关注以下几点:    
1. @SPI(RandomLoadBalance.NAME)        
@SPI加在LoadBalance接口上，表示接口LoadBalance是一个扩展点。如果没有@SPI注解修饰，ExtensionLoader会抛出异常。RandomLoadBalance.NAME是一个常量，值是"random"。如果没有显式指定LoadBalance的实现，默认使用random代表的扩展实现。
random所代表的扩展是哪一个类呢？答案就在文件`src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.cluster.LoadBalance`中:
```text
random=com.alibaba.dubbo.rpc.cluster.loadbalance.RandomLoadBalance
roundrobin=com.alibaba.dubbo.rpc.cluster.loadbalance.RoundRobinLoadBalance
leastactive=com.alibaba.dubbo.rpc.cluster.loadbalance.LeastActiveLoadBalance
consistenthash=com.alibaba.dubbo.rpc.cluster.loadbalance.ConsistentHashLoadBalance
```
可以看到文件中定义了4个LoadBalance的扩展实现。格式是`name=class全名称`。和Java SPI不同，dubbo SPI允许为每一个扩展实现取一个名字，然后就可以在Dubbo中通过name来引用对应的扩展实现。这也是Dubbo SPI优于Java SPI的地方。        
![dubbo-loadbalance](https://raw.githubusercontent.com/vangoleo/wiki/master/dubbo/dubbo_loadbalance.png)
2. @Adaptive("loadbalance")
@Adaptive注解修饰select方法，表明方法select方法是一个可自适应的方法。可以使用ExtenLoader获取一个LoadBalance的自适应实例，本质是一个代理。当调用实例的select方法时，会根据具体的方法参数来决定调用哪个扩展实现的select方法。@Adaptive注解的参数`loadbalance`表示方法参数中的loadbalance的值作为实际要调用的扩展实例。类似于从http的request中获取参数值。好比Dubbo的consumer端发送来一个请求http://domain.com/some/path?foo=100&loadbalance=random。     Provider端，获取参数中loadbalance的值为random。根据random来选择RandomLoadBalance。               
select的方法中好像没有loadbalance参数，那怎么获取loadbalance参数的值呢？我们看到select方法中有一个URL参数，可能很多人也猜到了，loadbalance就是以参数的形式存在于URL中的。URL是Java的一个类`com.alibaba.dubbo.common`。Dubbo使用了URL总线的模式，就是Dubbo的系统参数和每一次调用的参数都以URL的形式在各个层中传递。关于Dubbo的URL总线模式在后续章节会进行讨论。        
下面是URL类的一部分。可以看到里面有一个parameters的Map。这里面有Dubbo请求的相关参数。比如，序列化方式，负载均衡策略等信息。
```java
public final class URL implements Serializable {

    private final String protocol;

    private final String username;

    private final String password;

    // by default, host to registry
    private final String host;

    // by default, port to registry
    private final int port;

    private final String path;

    private final Map<String, String> parameters;
    
    // ......
    
```

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
                    throw new IllegalStateException("more than 1 default extension name on extension " + type.getName()
                            + ": " + Arrays.toString(names));
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


# Dubbo SPI高级用法之IoC
   AdaptiveInstance
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

