---
layout: post
title: Java Design Pattern - Singleton
date: 2015-06-26 15:30:00
tags:
- Linux
categories: Linux
---

# Simple Singleton

```java
public final class Singleton{
    private static final Singleton instance = new Singleton();
    private Singleton(){}
    public static Singleton getInstance(){
        return instance;
    }
}
```

# Initializing on demand Singleton
```java
public final class Singleton{
    private Singleton(){}
    public static Singleton getInstance(){
        return Holder.INSTANCE;
    }
    
    private static class Holder{
        public static final Singleton INSTANCE = new Singleton();
    }
}
// When the class Singleton is loaded by the JVM, the class goes through initialization. Since the class does not have any 
// static variables to initialize, the initialization completes trivially. The static class definition Holder within it is
// not initialized until the JVM determines that Holder must be executed. The static class Holder is only executed when
// the static method getInstance is invoked on the class Something, and the first time this happens the JVM will load and
// initialize the Holder class. The initialization of the Holder class results in static variable INSTANCE being
// initialized by executing the (private) constructor for the outer class Something. 
```

# Thread safe double check Singleton
```java
public final class Singleton{
    private static volatile Singleton instance;
    private Singleton(){
        if(null != instance){
            throw new Exception("Already initialized.");
        }
    }
    
    public static Singleton getInstance(){
        Singleton result = instance;
        if(null == result){
            synchronized (Singleton.class){
                result = instance;
                if(null == result){
                    instance = result = new Singleton();
                }
            }
        }
        return result;
    }
}
```

# Thread safe lazy load Singleton
```java
public final class Singleton{
    private static Singleton instance;
    private Singleton(){}
    public static synchronized Singleton getInstance(){
        if(null == instance){
            instance = new Singleton();
        }
    }
    return instance;
}
```


