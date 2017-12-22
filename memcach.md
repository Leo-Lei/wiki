---
layout: post
title: Memcached
date: 2017-04-10 13:05:00
tags:
- Java
categories: Java
---


# 连接Memcached
使用`telnet`方式来连接Memcached。
```bash
telnet 192.168.10.10 11211
```

```
[root@memcached ~]# telnet 192.168.10.10 11211
Trying 192.168.10.10...
Connected to 192.168.10.10.
Escape character is '^]'.
stats items
STAT items:2:number 1
STAT items:2:age 302
STAT items:2:evicted 0
STAT items:2:evicted_nonzero 0
STAT items:2:evicted_time 0
STAT items:2:outofmemory 0
STAT items:2:tailrepairs 0
STAT items:2:reclaimed 7
STAT items:2:expired_unfetched 0
STAT items:2:evicted_unfetched 0
STAT items:2:crawler_reclaimed 0
STAT items:2:crawler_items_checked 0
STAT items:2:lrutail_reflocked 0
END
quit
Connection closed by foreign host.

```
