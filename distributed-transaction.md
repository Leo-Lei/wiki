---
layout: post
title: Distributed Transaction
date: 2017-04-11 13:05:00
tags:
- Java
categories: Java
---

# 分布式事务的例子
从支付宝向余额宝转账。
可能会涉及到两张表：
* 支付宝账户表：A (id,userId,amount)
* 余额宝账户表：B (id,userId,amount)

# 本地事务
使用一个数据库事务来实现
```sql
Begin transaction
         update A set amount=amount-10000 where userId=1;
         update B set amount=amount+10000 where userId=1;
End transaction
commit;
```
# 分布式事务-两阶段提交
分为协调器和若干事务执行者两种角色。协调器可以和事务执行者在同一台机器上。

缺点：
性能太差，不适合高并发的系统。
1) 两阶段提交涉及多次节点间的网络通信，通信时间太长
2) 事务时间变长了，锁定的资源的时间也变长了，资源等待时间变长

# 补偿




# 使用消息队列来避免分布式事务















