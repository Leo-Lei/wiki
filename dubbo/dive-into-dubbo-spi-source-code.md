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
In the previous step, we get an empty instance of Extension. But in real project, a extension may have some dependencies. Just like the Spring Bean dependencies. But the dependencies in Dubbo extension is more complex than Spring. A dependency in Dubbo extension instance may be another Dubbo extension, or a simple Java class, or a spring bean, or other case, for example, a object managered by another IoC container. In next secion, we will cover this. Now, you just need to know, Dubbo can inject a dependency into a Dubbo extension. The dependency can be many types, a simple java object, a Dubbo extension, a Spring bean, or some other case. 

4. Extension Instance AOP         
Like Sping AOP, Dubbo also has the AOP feature. Dubbo use it to implement some common functionality, such as logging and monitoring. We will cover this later.      

After the above 4 steps, Dubbo has created a extension instance, inject dependencies to it, and wrap it accordingly. Now, the extension is ready to use.

# Dubbo SPI autowired
The code about autowired is in the `injectExtension`method:
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
In order to implement autowired, we first need to know which dependencies does a extension have. Dubbo search the standard Java setter method. If extension has a method whose name start with set, and has only one parameter, this method indicate that this extension has a dependency, the class type of dependency is the class of method parameter. This is similar with Spring setter injection. 

As we have mentioned above, a dependency in a Dubbo extension may be a simple java class, another Dubbo extension, a Spring bean, or other cases. So how does Dubbo process so many type of dependencies? In `在injectExtension` method, it use `Object object = objectFactory.getExtension(pt, property)` to implement this. objectFactory is type of ExtensionFactory, it is initialized when ExtensionLoader is created.
```java
private ExtensionLoader(Class<?> type) {
        this.type = type;
        objectFactory = (type == ExtensionFactory.class ? null : ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension());
    }
```
objectFacory itself is also an extension, created by `ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension())`.    
![Dubbo-ExtensionFactory](https://raw.githubusercontent.com/vangoleo/wiki/master/dubbo/dubbo-extensionfactory.png)
ExtensionFactory has 3 implementations：
1. SpiExtensionFactory：load Extension from Dubbo SPI
2. SpringExtensionFactory：load Extension from Spring Container
3. AdaptiveExtensionFactory: AdaptiveExtensionLoader
Please pay more attention to AdaptiveExtensionFactory，the source code is as below:
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
AdaptiveExtensionFactory has @Adaptive annotation。As mentioned previously，Dubbo will create an adaptive instance for each extension。If extention class has @Adaptive annotation, Dubbo will use it as the adaptive class. Otherwise, Dubbo will create an adaptive class automatically. So `ExtensionLoader.getExtensionLoader(ExtensionFactory.class).getAdaptiveExtension())` will return an AdaptiveExtensionLoader instance.        
AdaptiveExtentionFactory will iterate all the implementations of ExtensionFactory and try to find the required extension。Dubbo has 2 built-in ExtensionFactory which will load extension from Dubbo SPI and Spring container。ExtensionFactory itself is also an extention point，so we can create our customized ExtensionFactory to enhance Dubbo loading extention in our way。For example，the project use Google's Guice as the IoC container。Then we can implements a GuiceExtensionFactory, enhance Dubbo loading extension from Guice Container。

# Dubbo SPI AoP
Dubbo also support the AoP functionality, which is similat to Spring's AoP.        
In Dubbo, there is an particular class called Wrapper class. A wrapper class use the Decorator design pattern to add some additional functionality to the original object.

### what is Wrapper class
What is the Wrapper class in Dubbo? A wrapper class is a Java class which has a copy constructor function, which is knwon as the classical decorator design pattern. Below is a sample:
```java
class A{
    private A a;
    public A(A a){
        this.a = a;
    }
}
```
Class A has a constructor function `public A(A a)`，the parameter is A itself。Class A can be a Dubbo Wrapper class。Dubbo has such Wrapper  classes as ProtocolFilterWrapper, ProtocolListenerWrapper. You can go to read the source code to have better understanding.
### How to configure Wrapper class
Wrapper class is also a extension point. The same to other extensions，Wrapper class is configured in `META-INF` folder。For example, ProtocolFilterWrapper and ProtocolListenerWrapper are configured in `dubbo-rpc/dubbo-rpc-api/src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.Protocol`:
```text
filter=com.alibaba.dubbo.rpc.protocol.ProtocolFilterWrapper
listener=com.alibaba.dubbo.rpc.protocol.ProtocolListenerWrapper
mock=com.alibaba.dubbo.rpc.support.MockProtocol
```
Below is a code segment of Dubbo loading configuration files:
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
Use `clazz.getConstructor(type)` to get copy constructor function。 Please note that, the copy constructor parameter type is the extension interface，not the concrete extension class.        
Take the Protocol for example。Configuration file `dubbo-rpc/dubbo-rpc-api/src/main/resources/META-INF/dubbo/internal/com.alibaba.dubbo.rpc.Protocol` defined `filter=com.alibaba.dubbo.rpc.protocol.ProtocolFilterWrapper`.        
ProtocolFilterWrapper source code as below：
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
ProtocolFilterWrapper has a constructor function `public ProtocolFilterWrapper(Protocol protocol)`，parameter is Protocol，so it is a Dubbo Wrapper class。 

# Extention auto adaptive
Dubbo need judge the method parameters at runtime to create corresponding extension implementation. so dubbo has the auto adaptive extension concept. A extension adaptive instance is a proxy of the extension point. Each extension in Dubbo has an adaptive extension instance. If we don't supply it explicitly, Dubbo will create one for us. By default, dubbo use Javaassist to generate it。        
The code of getAdaptiveExtension
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
continue to go to createAdaptiveExtension method:
```java
private T createAdaptiveExtension() {        
    return injectExtension((T) getAdaptiveExtensionClass().newInstance());
}
```
continue to go to getAdaptiveExtensionClass method:
```java
private Class<?> getAdaptiveExtensionClass() {
        getExtensionClasses();
        if (cachedAdaptiveClass != null) {
            return cachedAdaptiveClass;
        }
        return cachedAdaptiveClass = createAdaptiveExtensionClass();
    }
```
continue to go to createAdaptiveExtensionClass method. The createAdaptiveExtensionClass method first generate the source code of adaptive extension class，then compile the source code to Java bytecode, and load it to JVM.         
```java
private Class<?> createAdaptiveExtensionClass() {
        String code = createAdaptiveExtensionClassCode();
        ClassLoader classLoader = findClassLoader();
        com.alibaba.dubbo.common.compiler.Compiler compiler = ExtensionLoader.getExtensionLoader(com.alibaba.dubbo.common.compiler.Compiler.class).getAdaptiveExtension();
        return compiler.compile(code, classLoader);
    }
```
Source code of Compiler，default implementation is javassist。
```java
@SPI("javassist")
public interface Compiler {
    Class<?> compile(String code, ClassLoader classLoader);
}
```
createAdaptiveExtensionClassCode() method use a StringBuilder to generate source code。 Below is an example of Protocol adaptive extension class generated by Dubbo.
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
First resolve url parameters，get the corresponding value from @Adaptive's value parameter，then get the concrete extension implemetation.
We found that in the generated class Protocol$Adpative，the getDefaultPort and destroy throw exception directly. Why? Let's look at the source code of Protocol:
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
We can see that, Protocol interface has 4 methods，but only export and refer has @Adaptive annotation。When Dubbo generate adaptive extension automatically，only the method with @Adaptive has the method implementation。So，in Protocol$Adpative class，only export and refer has real method implementation，other methods directly throw exception。    
