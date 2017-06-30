---
layout: post
title: Java Class
date: 2017-06-30 12:20:00
tags:
- Java
categories: Java
description: The tutoria will describe the useage of Linux.
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
1. 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
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

2. 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
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

3. 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
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
4. 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
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
5. 应用程序有且仅有两个类A和Main如下，程序的输出是什么样的？
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


















