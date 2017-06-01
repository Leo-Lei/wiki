---
layout: post
title: Amazon DynamoDB
date: 2017-05-26 11:10:00
tags:
- docker
categories: Java
description: Amazon DynamoDB
---

# 全局二维索引
分区键和排序键可以与基表的分区键和排序键不同。全局二维索引被视为“全局”，因为对索引执行的查询可以跨基表中所有分区的所有数据。

# 本地二级索引
分区键和基表相同，但排序键不同的索引。被视为本地，因为对索引执行的查询被限定为具有相同分区键的基表分区。



# 链接
[SecondaryIndexes.html](http://docs.aws.amazon.com/zh_cn/amazondynamodb/latest/developerguide/SecondaryIndexes.html)
[https://segmentfault.com/a/1190000008232397](https://segmentfault.com/a/1190000008232397)
