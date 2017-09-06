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
        Bar mockBar = PowerMockito.mock(Bar.class);
        Foo foo = new Foo();
        PowerMockito.when(mockBar.isAlive()).thenReturn(true);
        System.out.println(foo.isAlive(mockBar));
    }
}

```
