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
    Class<T> cls = (Class<T>) obj.getClass();
    LinkedBuffer buffer = LinkedBuffer.allocate(LinkedBuffer.DEFAULT_BUFFER_SIZE);
    try {
        Schema<T> schema = getSchema(cls);
        return ProtostuffIOUtil.toByteArray(obj, schema, buffer);
    } catch (Exception e) {
        throw new IllegalStateException(e.getMessage(), e);
    } finally {
        buffer.clear();
    }
}

public static <T> T deserialize(byte[] data, Class<T> cls) {
    try {
        T message = objenesis.newInstance(cls);
        Schema<T> schema = getSchema(cls);
        ProtostuffIOUtil.mergeFrom(data, message, schema);
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


# Demo
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






