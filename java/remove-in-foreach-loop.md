---
layout: post
title: Calling remove method in foreach loop
date: 2016-06-20 16:30:00
tags:
- Windows
categories: Windows
---

The following code will throw an exception `java.util.ConcurrentModificationException`
```java  
List<String> list = ......
for (String item : list){
    // do something...
    list.remove(item);
}
```

You can use below code, which use an iterator.
```java  
List<String> list = ......
Iterator<String> iterator = list.iterator();
while (iterator.hasNext()){
    // do something...
    intrator.remove();
}
```
