---
layout: post
title: Protostuff
date: 2017-06-30 16:30:00
tags:
- docker
categories: Java
description: docker
---


# Protostuff简介
Protostuff主页:[http://www.protostuff.io/](http://www.protostuff.io/)

# 添加Maven依赖

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
<dependency>
    <groupId>org.objenesis</groupId>
    <artifactId>objenesis</artifactId>
    <version>2.1</version>
</dependency>
```

# 序列化和反序列化工具
```java
@SuppressWarnings("unchecked")
public static <T> byte[] serialize(T obj) {
    Class<T> cls = (Class<T>) obj.getClass();            //获得对象的Class
    LinkedBuffer buffer = LinkedBuffer.allocate(LinkedBuffer.DEFAULT_BUFFER_SIZE);  //使用LinkedBuffer分配一块默认大小的buff空间
    try {
        Schema<T> schema = getSchema(cls);               //通过对象的Class构建对应的Schema
        return ProtostuffIOUtil.toByteArray(obj, schema, buffer);  //使用给定的Schema将对象序列化为一个byte数组
    } catch (Exception e) {
        throw new IllegalStateException(e.getMessage(), e);
    } finally {
        buffer.clear();
    }
}

public static <T> T deserialize(byte[] data, Class<T> cls) {
    try {
        T message = objenesis.newInstance(cls);        //使用objenesis实例化一个类的对象
        Schema<T> schema = getSchema(cls);             //通过对象的类构建对应的schema
        ProtostuffIOUtil.mergeFrom(data, message, schema);   //使用给定的schema将byte数组和对象合并，并返回
        return message;
    } catch (Exception e) {
        throw new IllegalStateException(e.getMessage(), e);
    }
}

private static Map<Class<?>, Schema<?>> cachedSchema = new ConcurrentHashMap<>();

// 构建Schema比较耗时，可以把Schema缓存起来
private static <T> Schema<T> getSchema(Class<T> cls) {
    Schema<T> schema = (Schema<T>) cachedSchema.get(cls);
    if (schema == null) {
        schema = RuntimeSchema.createFrom(cls);
        if (schema != null) {
            cachedSchema.put(cls, schema);
        }
    }
    return schema;
}
```
