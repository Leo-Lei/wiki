---
layout: post
title: Java Class
date: 2017-06-30 12:20:00
tags:
- Java
categories: Java
---


# 类的生命周期
1. 加载
2. 验证
3. 准备
4. 解析
5. 初始化
6. 使用
7. 卸载

# 类的初始化
关于类的初始化，比如静态变量的初始化，网上也有很多的文章，但是感觉好像都是点到为止，看完了，感觉和没看是一样的效果。
再比如，单例模式，使用静态内部类的static holder方案，可以实现线程安全和延迟初始化。但如果我继续问一些人，为什么可以达到这样的效果，绝大多数人又说不清楚了。
其实就2个问题：    
1. 什么时候初始化    
2. 初始化什么东西    

先提出如下几个问题吧：   
### 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    static {
        System.out.println("class A static block ...");
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
输出：
```bash
Main Class main method ...
```
是的，没有输出"class A static block ..."

### 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    static {
        System.out.println("class A static block ...");
    }
}

public class Main {
    private A a;
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```

<!-- more -->

### 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    static {
        System.out.println("class A static block ...");
    }
}

public class Main {
    private static A a;
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
### 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    static {
        System.out.println("class A static block ...");
    }
}

public class Main {
    private static A a = new A();
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
### 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    private static A instance = new A();
    public A(){
        System.out.println("A constructor ...");
    }
}

public class Main {
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
### 应用程序有且仅有三个类A,B和Main如下，程序的输出是什么样的？
```java
public class A {
    private static A instance = new A();
    public A(){
        System.out.println("A constructor ...");
    }
}
public class B {
    private static A instance = new A();
}
public class Main {
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
### 应用程序有且仅有2个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    private static A instance = new A();
    public A(){
        System.out.println("A constructor ...");
    }
}

public class Main {
    private static A;
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
### 应用程序有且仅有2个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    private static A instance = new A();
    public A(){
        System.out.println("A constructor ...");
    }
}

public class Main {
    private static A = new A();
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
### 应用程序有且仅有2个类A和Main如下，程序的输出是什么样的？
```java
public class A {
    private static A instance = new A();
    public A(){
        System.out.println("A constructor ...");
    }
}

public class Main {
    private static A;
    public static void main(String[] args) {
        System.out.println("Main Class main method ...");
    }
}
```
# 什么时候初始化类?
* 虚拟机启动时，包含有main方法的类
* 创建某个类的新实例时。注包括new()，反射，克隆以及反序列化等。
* 调用了某个类的静态方法时
* 当使用了某个类的静态字段时。注意，用final修饰的静态字段除外，因为在编译阶段，它就被替换成一个常量表达式了。
* 调用java API中的某些反射方法时，比如类Class中的方法或java.lang.reflect包中的方法
* 初始化某个类的子类时（某个类初始化时，要求它的父类要已经初始化了）

# 类初始化时，初始化什么东西
* 初始化类的静态字段和静态快，按照类中声明的顺序来依次执行

