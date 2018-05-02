---
layout: post
title: Dubbo SPI源码解析
date: 2018-04-09 16:30:00
tags:
- Dubbo
categories: Java
---


# Dubbo Extension Loader
In this section, we will dive into the source code of Dubbo SPI.
### ExtensionLoader
extensionLoader is the most important class, which is responsible to load extension configuration and create extension. Let's start from this class. ExtensionLoader has many method, below are some common methods: 
* `public static <T> ExtensionLoader<T> getExtensionLoader(Class<T> type)`
* `public T getExtension(String name)`
* `public T getAdaptiveExtension()`
Some common usage:
* `LoadBalance lb = ExtensionLoader.getExtensionLoader(LoadBalance.class).getExtension(loadbalanceName)`
* `RouterFactory routerFactory = ExtensionLoader.getExtensionLoader(RouterFactory.class).getAdaptiveExtension()`

1. getExtensionLoader method
This is an static factory method, the method parameter is the interface of extension point, and return an ExtensionLoader instance.
```java
public static <T> ExtensionLoader<T> getExtensionLoader(Class<T> type) {
        // the extension point must be an interface
        if (!type.isInterface()) {
            throw new IllegalArgumentException("Extension type(" + type + ") is not interface!");
        }
        // must have @SPI annotation
        if (!withExtensionAnnotation(type)) {
            throw new IllegalArgumentException("Extension type without @SPI Annotation!");
        }
        // get ExtensionLoader from cache
        // each Extension will only be initialized once
        ExtensionLoader<T> loader = (ExtensionLoader<T>) EXTENSION_LOADERS.get(type);
        if (loader == null) {
            // init
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
2. getExtension method
```java
public T getExtension(String name) {
        Holder<Object> holder = cachedInstances.get(name);
        if (holder == null) {
            cachedInstances.putIfAbsent(name, new Holder<Object>());
            holder = cachedInstances.get(name);
        }
        Object instance = holder.get();
        // try to get from cache
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
The main process logic is in the createExtension method, let's go ahead to createExtension method.
```java
private T createExtension(String name) {
        // get extension by extension point
        Class<?> clazz = getExtensionClasses().get(name);
        
        T instance = (T) EXTENSION_INSTANCES.get(clazz);
        if (instance == null) {
              // use reflection to create a new instance of Extension
            EXTENSION_INSTANCES.putIfAbsent(clazz, (T) clazz.newInstance());
            instance = (T) EXTENSION_INSTANCES.get(clazz);
        }
        // IoC
        injectExtension(instance);
        // AoP
        Set<Class<?>> wrapperClasses = cachedWrapperClasses;
        if (wrapperClasses != null && !wrapperClasses.isEmpty()) {
            for (Class<?> wrapperClass : wrapperClasses) {
                instance = injectExtension((T) wrapperClass.getConstructor(type).newInstance(instance));
            }
        }
        return instance;
}
```
createExtension will do the following things:    
1. Get concrete extension by name. Load extension configuration from config files in `META-INF` path.
2. Use reflection to create a new instance of extension.
3. Auto inject the dependency of Extension Instance. i.e. IoC.
4. Auto wrap the Extension instance, i.e. AoP.

Let's dive into these 4 process:
1. get extension by name
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
The process is quite simple, first try to get from cache. If not exist in cache, then load from below config files:
* `META-INF/dubbo/internal`
* `META-INF/dubbo`
* `META-INF/services`

2. Use reflection to create a nes instance of Extension.
Use the `clazz.newInstance()` to create instance. After this process, the instance of Extension is created, but the properties of it is all empty value.

3. Auto dependency injection
In the previous step, we get an empty instance of Extension. 
在实际的场景中，类之间都是有依赖的。扩展实例中也会引用一些依赖，比如简单的Java类，另一个Dubbo的扩展或一个Spring Bean等。依赖的情况很复杂，Dubbo的处理也相对复杂些。我们稍后会有专门的章节对其进行说明，现在，我们只需要知道，Dubbo可以正确的注入扩展点中的普通依赖，Dubbo扩展依赖或Spring依赖等。

4. 扩展实例自动包装         
自动包装就是要实现类似于Spring的AOP功能。Dubbo利用它在内部实现一些通用的功能，比如日志，监控等。关于扩展实例自动包装的内容，也会在后面单独讲解。

经过上面的4步，Dubbo就创建并初始化了一个扩展实例。这个实例的依赖被注入了，也根据需要被包装了。到此为止，这个扩展实例就可以被使用了。        

# Dubbo SPI高级用法之自动装配
自动装配的相关代码在injectExtension方法中:
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
要实现对扩展实例的依赖的自动装配，首先需要知道有哪些依赖，这些依赖的类型是什么。Dubbo的方案是查找Java标准的setter方法。即方法名以set开始，只有一个参数。如果扩展类中有这样的set方法，Dubbo会对其进行依赖注入，类似于Spring的set方法注入。    
但是Dubbo中的依赖注入比Spring要复杂，因为Spring注入的都是Spring bean，都是由Spring容器来管理的。而Dubbo的依赖注入中，需要注入的可能是另一个Dubbo的扩展，也可能是一个Spring Bean，或是Google guice的组件，或其他任何一个框架中的组件。Dubbo需要能够从任何一个场景中加载扩展。在injectExtension方法中，是用`Object object = objectFactory.getExtension(pt, property)`来实现的。objectFactory是ExtensionFactory类型的，在创建ExtensionLoader时被初始化: 
```java
private ExtensionLoader(Class<?> type) {
        this.type = type;
        objectFactory = (type == ExtensionFactory.class ? null : ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension());
    }
```
objectFacory本身也是一个扩展，通过`ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension())`来获取。
![Dubbo-ExtensionFactory](https://raw.githubusercontent.com/vangoleo/wiki/master/dubbo/dubbo-extensionfactory.png)
ExtensionLoader有三个实现：
1. SpiExtensionLoader：Dubbo自己的Spi去加载Extension
2. SpringExtensionLoader：从Spring容器中去加载Extension
3. AdaptiveExtensionLoader: 自适应的AdaptiveExtensionLoader
这里要注意AdaptiveExtensionLoader，源码如下:
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
AdaptiveExtensionLoader类有@Adaptive注解。前面提到了，Dubbo会为每一个扩展创建一个自适应实例。如果扩展类上有@Adaptive，会使用该类作为自适应类。如果没有，Dubbo会为我们创建一个。所以`ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension())`会返回一个AdaptiveExtensionLoader实例，作为自适应扩展实例。
AdaptiveExtentionLoader会遍历所有的ExtensionFactory实现，尝试着去加载扩展。如果找到了，返回。如果没有，在下一个ExtensionFactory中继续找。Dubbo内置了两个ExtensionFactory，分别从Dubbo自身的扩展机制和Spring容器中去寻找。由于ExtensionFactory本身也是一个扩展点，我们可以实现自己的ExtensionFactory，让Dubbo的自动装配支持我们自定义的组件。比如，我们在项目中使用了Google的guice这个IoC容器。我们可以实现自己的GuiceExtensionFactory，让Dubbo支持从guice容器中加载扩展。

# Dubbo SPI高级用法之AoP
在用Spring的时候，我们经常会用到AOP功能。在目标类的方法前后插入其他逻辑。比如通常使用Spring AOP来实现日志，监控和鉴权等功能。    
Dubbo的扩展机制，是否也支持类似的功能呢？答案是yes。在Dubbo中，有一种特殊的类，被称为Wrapper类。通过装饰者模式，使用包装类包装原始的扩展点实例。在原始扩展点实现前后插入其他逻辑，实现AOP功能。    

### 什么是Wrapper类
那什么样类的才是Dubbo扩展机制中的Wrapper类呢？Wrapper类是一个有复制构造函数的类，也是典型的装饰者模式。下面就是一个Wrapper类:
```java
class A{
    private A a;
    public A(A a){
        this.a = a;
    }
}
```
类A有一个构造函数`public A(A a)`，构造函数的参数是A本身。这样的类就可以成为Dubbo扩展机制中的一个Wrapper类。Dubbo中这样的Wrapper类有ProtocolFilterWrapper, ProtocolListenerWrapper等, 大家可以查看源码加深理解。
### 怎么配置Wrapper类
在Dubbo中Wrapper类也是一个扩展点，和其他的扩展点一样，也是在`META-INF`文件夹中配置的。比如前面举例的ProtocolFilterWrapper和ProtocolListenerWrapper就是在路径`dubbo-rpc/dubbo-rpc-api/src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.Protocol`中配置的:
```text
filter=com.alibaba.dubbo.rpc.protocol.ProtocolFilterWrapper
listener=com.alibaba.dubbo.rpc.protocol.ProtocolListenerWrapper
mock=com.alibaba.dubbo.rpc.support.MockProtocol
```
在Dubbo加载扩展配置文件时，有一段如下的代码:
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
这段代码的意思是，如果扩展类有复制构造函数，就把该类存起来，供以后使用。有复制构造函数的类就是Wrapper类。通过`clazz.getConstructor(type)`来获取参数是扩展点接口的构造函数。注意构造函数的参数类型是扩展点接口，而不是扩展类。    
以Protocol为例。配置文件`dubbo-rpc/dubbo-rpc-api/src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.Protocol`中定义了`filter=com.alibaba.dubbo.rpc.protocol.ProtocolFilterWrapper`。            
ProtocolFilterWrapper代码如下：
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
ProtocolFilterWrapper有一个构造函数`public ProtocolFilterWrapper(Protocol protocol)`，参数是扩展点Protocol，所以它是一个Dubbo扩展机制中的Wrapper类。ExtensionLoader会把它缓存起来，供以后创建Extension实例的时候，使用这些包装类依次包装原始扩展点。

# 扩展点自适应
前面讲到过，Dubbo需要在运行时根据方法参数来决定该使用哪个扩展，所以有了扩展点自适应实例。其实是一个扩展点的代理，将扩展的选择从Dubbo启动时，延迟到RPC调用时。Dubbo中每一个扩展点都有一个自适应类，如果没有显式提供，Dubbo会自动为我们创建一个，默认使用Javaassist。        
先来看下创建自适应扩展类的代码:
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
继续看createAdaptiveExtension方法
```java
private T createAdaptiveExtension() {        
    return injectExtension((T) getAdaptiveExtensionClass().newInstance());
}
```
继续看getAdaptiveExtensionClass方法
```java
private Class<?> getAdaptiveExtensionClass() {
        getExtensionClasses();
        if (cachedAdaptiveClass != null) {
            return cachedAdaptiveClass;
        }
        return cachedAdaptiveClass = createAdaptiveExtensionClass();
    }
```
继续看createAdaptiveExtensionClass方法，绕了一大圈，终于来到了具体的实现了。看这个createAdaptiveExtensionClass方法，它首先会生成自适应类的Java源码，然后再将源码编译成Java的字节码，加载到JVM中。
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
createAdaptiveExtensionClassCode()方法中使用一个StringBuilder来构建自适应类的Java源码。方法实现比较长，这里就不贴代码了。这种生成字节码的方式也挺有意思的，先生成Java源代码，然后编译，加载到jvm中。通过这种方式，可以更好的控制生成的Java类。而且这样也不用care各个字节码生成框架的api等。因为xxx.java文件是Java通用的，也是我们最熟悉的。只是代码的可读性不强，需要一点一点构建xx.java的内容。            
下面是使用createAdaptiveExtensionClassCode方法为Protocol创建的自适应类的Java代码范例:
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
大致的逻辑和开始说的一样，通过url解析出参数，解析的逻辑由@Adaptive的value参数控制，然后再根据得到的扩展点名获取扩展点实现，然后进行调用。如果大家想知道具体的构建.java代码的逻辑，可以看`createAdaptiveExtensionClassCode`的完整实现。        
在生成的Protocol$Adpative中，发现getDefaultPort和destroy方法都是直接抛出异常的，这是为什么呢？来看看Protocol的源码
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
可以看到Protocol接口中有4个方法，但只有export和refer两个方法使用了@Adaptive注解。Dubbo自动生成的自适应实例，只有@Adaptive修饰的方法才有具体的实现。所以，Protocol$Adpative类中，也只有export和refer这两个方法有具体的实现，其余方法都是抛出异常。    
