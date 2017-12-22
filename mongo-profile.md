---
layout: post
title: Mongo Profile
date: 2017-06-11 13:05:00
tags:
- Java
categories: Java
---


# mongo启用profile

mongodb可以通过profile来监控数据，进行优化。

查看当前是否开启profile功能用命令

db.getProfilingLevel()  返回level等级，值为0|1|2，分别代表意思：0代表关闭，1代表记录慢命令，2代表全部

 

db.setProfilingLevel(level);  #level等级，值同上

level为1的时候，慢命令默认值为100ms，更改为db.setProfilingLevel(level,slowms)如db.setProfilingLevel(1,50)这样就更改为50毫秒

通过db.system.profile.find() 查看当前的监控日志。

如：

> db.system.profile.find({millis:{$gt:500}})
{ "ts" : ISODate("2011-07-23T02:50:13.941Z"), "info" : "query order.order reslen:11022 nscanned:672230  \nquery: { status: 1.0 } nreturned:101 bytes:11006 640ms", "millis" : 640 }
{ "ts" : ISODate("2011-07-23T02:51:00.096Z"), "info" : "query order.order reslen:11146 nscanned:672302  \nquery: { status: 1.0, user.uid: { $gt: 1663199.0 } }  nreturned:101 bytes:11130 647ms", "millis" : 647 }
 

 这里值的含义是

 ts：命令执行时间

info：命令的内容

query：代表查询

order.order： 代表查询的库与集合

reslen：返回的结果集大小，byte数

nscanned：扫描记录数量

nquery：后面是查询条件

nreturned：返回记录数及用时
