---
title: Java Dynamic Proxy
date: 2017-06-09 10:22:23
categories:
- Music
tags:
- Music
---

# Java Dynamic Proxy
要代理的接口
```java
public interface IHello {
    public String hello(String str);
}
```
真实的实现对象
```java
public class Hello implements IHello {

    @Override
    public String hello(String str) {
        System.out.println("Say Hello: " + str);
        return "Hello: " + str;
    }
}
```
创建InvocationHandler实例
```java
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;

public class DynamicProxy implements InvocationHandler {
    private Object subject;

    public DynamicProxy(Object subject){
        this.subject = subject;
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        System.out.println("Before say hello...");

        System.out.println("Method: " + method);

        Object result = method.invoke(subject,args);

        System.out.println("After say hello...");

        return result;
    }
}

```

Client端使用动态代理
```java
IHello realSubject = new Hello();

InvocationHandler handler = new DynamicProxy(realSubject);

IHello subject = (IHello) Proxy.newProxyInstance(
        handler.getClass().getClassLoader(),
        new Class[]{IHello.class},
        handler);

System.out.println(subject.getClass().getName());
String result = subject.hello("world");
System.out.println(result);
```

