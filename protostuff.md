---
layout: post
title: Protostuff
date: 2017-06-30 16:30:00
tags:
- docker
categories: Java
description: docker
---



```xml
<dependency>
    <groupId>io.protostuff</groupId>
    <artifactId>protostuff-core</artifactId>
    <version>1.4.4</version>
</dependency>
<dependency>
    <groupId>io.protostuff</groupId>
    <artifactId>protostuff-runtime</artifactId>
    <version>1.4.4</version>
</dependency>
```


```java
private static RuntimeSchema<Person> schema = RuntimeSchema.createFrom(Person.class);
/**
*序列化
*/
Person crab = new Person();  
crab.setName("kaka");  
//参数三缓冲器
byte[] bytes = ProtostuffIOUtil.toByteArray(crab,schema,LinkedBuffer.allocate(LinkedBuffer.DEFAULT_BUFFER_SIZE));
/**
 *反序列化
 */
// 空对象
Person newCrab = schema.newMessage();
ProtostuffIOUtil.mergeFrom(bytes,newCrab,schema);
System.out.println("Hi, My name is " + newCrab.getName());
```
