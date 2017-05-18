---
layout: post
title: Redis
date: 2016-07-15 13:05:00
tags:
- Java
categories: Java
description: web-authentication
---


# Redis数据类型

| type         |               Desc              |
| ------------ | ------------------------------- |
| string       | 字符串                           |
| List         | 按插入顺序排序的字符串集合           |
| Set          | 不重复，无序的字符串集合             |
| Sorted set   | 不重复，有序的字符串集合             |
| Hash         | key-value的map                   |

# Redis的key
key可以是"hello"字符串，也可以是一个JPEG文件的内容。一般都用string做为key。    
关于key的几条建议：
* `object-type:id:field`是一个比较不错的模式。比如"user:1000:mypassword",或"comment:1000:reply.to"

# Redis string
```bash
> set mykey somevalue
OK
> get mykey
"somevalue"
```
使用`incr`来递增一个数字str,是原子递增的。
```bash
> set counter 100
OK
> incr counter
(integer) 101
> incr counter
(integer) 102
> incrby counter 50
(integer) 152
```
类似的命令有`incr`,`decr`,`incrby`,`decrby`。
```bash
> set mykey hello
OK
> exists mykey
(integer) 1
> del mykey
(integer) 1
> exists mykey
(integer) 0
```


# yum安装redis-cli
```bash
yum install epel-release
```

```bash
yum install redis
```
