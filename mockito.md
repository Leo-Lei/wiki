---
layout: post
title: Mockito
date: 2017-08-16 14:30:00
tags:
- docker
categories: Java
description: mockito
---

Mockito可以mock:    
* public方法


Mockito不能mock:        
* static方法
* final方法
* private方法


# 创建Mock对象
```java
//基于接口创建mock对象
IHelloService mockService = mock(IHelloService.class);
//基于类创建mock对象
User mockUser = mock(User.class);
```

<!-- more -->

# 对mock对象进行stub
```java
LinkedList mockList = mock(LinkedList.class);
  
when(mockList.get(0)).thenReturn("first");
when(mockList.get(3)).thenReturn("third");
when(mockList.get(4)).thenThrow(new RuntimeException());
  
System.out.println(mockList.get(0));    // 输出first
System.out.println(mockList.get(1));    // mockito框架报错，因为我们没有对get(1)进行stub
System.out.println(mockList.get(4));    // 抛出我们设定的RuntimeException
```

# 验证mock对象的方法被执行了
```java
LinkedList mockList = mock(LinkedList.class);
when(mockList.get(3)).thenReturn("third");
// 验证mockList的get方法被执行了，且参数是3
verify(mockList).get(3);
```

# 在stub和verify时支持参数匹配
```java
LinkedList mockList = mock(LinkedList.class);
// 调用mockList的get方法，传入任何int，都返回“element”。
when(mockList.get(anyInt())).thenReturn("element");
mockList.get(0);
mockList.get(1);
// 验证mockList方法被执行了2次，参数是任意的int类型
verify(mockList,times(2)).get(anyInt());
```

```java
@Test
public void test(){
    List mock = mock(List.class);
    when(mock.addAll(argThat(new ListOfTwoElements))).thenReturn(true);
    mock.addAll(Arrays.asList("one", "two"));
    verify(mock).addAll(argThat(new ListOfTwoElements()));
}
 
class ListOfTwoElements implements ArgumentMatcher<List> {
     public boolean matches(List list) {
         return list.size() == 2;
     }
 }

```

```java
@Test
public void test() {
    List mockList = mock(List.class);
    // 当get()的index参数小于5时，返回hello
    // 由于get(int index)方法的参数是原始类型int，不是Integer。所以，这里要使用intThat，而不能使用argThat，不然会有NullPointException。
    when(mockList.get(intThat(new IntLessThan5ArgsMatcher()))).thenReturn("hello");
    // 当get()的index参数大于5时，返回world
    when(mockList.get(intThat(new IntGreaterThan5ArgsMatcher()))).thenReturn("world");
 
    System.out.println(mockList.get(3));       // 输出hello
    System.out.println(mockList.get(10));      // 输出world
    // 验证mockList的get方法执行了，且参数index小于5
    verify(mockList).get(intThat(new IntLessThan5ArgsMatcher()));
    // 验证mockList的get方法执行了，且参数index大于5
    verify(mockList).get(intThat(new IntGreaterThan5ArgsMatcher()));
}
 
public static class IntLessThan5ArgsMatcher implements ArgumentMatcher<Integer> {
    @Override
    public boolean matches(Integer argument) {
        return argument < 5;
    }
}
 
public static class IntGreaterThan5ArgsMatcher implements ArgumentMatcher<Integer> {
 
    @Override
    public boolean matches(Integer argument) {
        return argument > 5;
    }
}
```

# 验证一个方法被执行了几次
```java
List mockList = mock(List.class);
  
mockList.add("once");
  
mockList.add("twice");
mockList.add("twice");
  
mockList.add("three times");
mockList.add("three times");
mockList.add("three times");
  
// 验证方法add("once")被执行了一次
verify(mockList).add("once");
// 和上面的效果一摸一样，默认是times(1)
verify(mockList,times(1)).add("once");
  
verify(mockList,times(2)).add("twice");
verify(mockList,times(3)).add("three times");
  
verify(mockList,never()).add("never happened");
  
verify(mockList,atLeastOnce()).add("three times");
verify(mockList,atLeast(2)).add("three times");
verify(mockList,atMost(5)).add("three times");
```


# stub void方法
使用doNothing        
```java
@Test
public void test(){
    List mockList = mock(List.class);
    doNothing().when(mockList).clear();
    mockList.clear();
}
```

自定义一个Answer
```java
@Test
public void test001(){
    List mockList = mock(List.class);
 
    doAnswer(new Answer() {
        @Override
        public Object answer(InvocationOnMock invocation) throws Throwable {
            Object[] args = invocation.getArguments();
            System.out.println("add..." + args[0] + " " + args[1]);
            return null;
        }
    }).when(mockList).add(anyInt(),any());
 
    mockList.add(1,"element");
}
```


# SPY
mockito支持基于一个真实对象，创建一个SPY。SPY可以对真实对象的某些方法进行stub，其余方法保持不变。

```java
@Test
public void test(){
    List list = new LinkedList();
    List spy = spy(list);
 
    when(spy.size()).thenReturn(100);
 
    spy.add("one");
    spy.add("two");
 
    System.out.println(spy.get(0));
 
    System.out.println(spy.size());
 
    verify(spy).add("one");
    verify(spy).add("two");
}
```


# 同一个方法执行多次时，返回不同的值
```java
IHelloService helloService = mock(IHelloService.class);

when(helloService.sayHello())
        .thenReturn("11")
        .thenReturn("22")
        .thenReturn("33");

System.out.println(helloService.sayHello());
System.out.println(helloService.sayHello());
System.out.println(helloService.sayHello());
```
输出:        
```text
11
22
33
```




