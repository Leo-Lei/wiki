---
layout: post
title: PowerMock
date: 2017-08-16 14:30:00
tags:
- docker
categories: Java
description: PowerMock
---

环境：    
* Mockito 2.8.x 
* PowerMock 1.7.1 
* Junit 4.4及以上


```groovy
dependencies {
    testCompile("junit:junit:4.11")
    testCompile("org.mockito:mockito-core:2.8.47")
    testCompile("org.powermock:powermock-module-junit4:1.7.1")
    testCompile("org.powermock:powermock-api-mockito2:1.7.1")
}
```


# mock方法内部new出来的对象

```java
public class FileFinder {
    public boolean exists(String path){
        File file = new File(path);
        return file.exists();
    }
}
```

```java
@RunWith(PowerMockRunner.class)
public class FileFinderTest {

    @Test
    @PrepareForTest(FileFinder.class)
    public void test() throws Exception {
        File mockFile = mock(File.class);
        when(mockFile.exists()).thenReturn(true);
        FileFinder fileFinder = new FileFinder();
        PowerMockito.whenNew(File.class).withArguments("bbb").thenReturn(mockFile);
        //PowerMockito.when(mockFile.exists()).thenReturn(true);
        boolean b = fileFinder.exists("bbb");
        System.out.println(b);
    }
}
```

<!-- more -->

# mock final方法
```java
public class Foo {
    public boolean isAlive(Bar bar){
        return bar.isAlive();
    }
}

public class Bar {
    public final boolean isAlive(){
        // do something
        return false;
    }
}

@RunWith(PowerMockRunner.class)
public class FooTest {

    @Test
    @PrepareForTest(Bar.class)
    public void test(){
        Bar mockBar = PowerMockito.mock(Bar.class);    // 使用PowerMockito.mock(),而不是原生的mockito的mock方法
        Foo foo = new Foo();
        PowerMockito.when(mockBar.isAlive()).thenReturn(true);
        System.out.println(foo.isAlive(mockBar));
    }
}

```

# mock static方法
```java
public class Foo {
    public boolean isAlive(){
        return Bar.isAlive();
    }
}
```

```java
public class Bar {
    public static boolean isAlive(){
        // do something
        return false;
    }
}
```

```java
@RunWith(PowerMockRunner.class)
public class FooTest {

    @Test
    @PrepareForTest(Bar.class)
    public void test(){
        //Bar mockBar = PowerMockito.mock(Bar.class);
        Foo foo = new Foo();
        PowerMockito.mockStatic(Bar.class);
        PowerMockito.when(Bar.isAlive()).thenReturn(true);
        System.out.println(foo.isAlive());
    }
}
```

# mock private方法
mock private方法的时候，我们会创建一个真实的对象，然后希望mock它的某些private方法。        
```java
public class Foo {
    public boolean isAlive0(){
        return isAlive();
    }

    private boolean isAlive(){
        return false;
    }
}
```

```java
@RunWith(PowerMockRunner.class)
public class FooTest {

    @Test
    @PrepareForTest(Foo.class)
    public void test() throws Exception {
        Foo foo = PowerMockito.spy(new Foo());
        PowerMockito.when(foo,"isAlive").thenReturn(true);
        boolean b = foo.isAlive0();
        System.out.println(b);
        PowerMockito.verifyPrivate(foo,times(1)).invoke("isAlive");
    }
}
```











