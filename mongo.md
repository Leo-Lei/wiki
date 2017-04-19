---
layout: post
title: Mongo
date: 2017-04-11 13:05:00
tags:
- Java
categories: Java
description: web-authentication
---


# yum安装mongo
```bash
/etc/yum.repos.d/mongodb-org-3.4.repo
```

```bash
[mongodb-org-3.4]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/3.4/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-3.4.asc
```

```bash
yum install mongodb-org            # 安装所有mongo的组件
yum install mongodb-org-shell      # 安装mongo客户端
```

# 连接mongodb
```bash
mongo --host mongo-server:3717 -u root -p
```


# mongo命令
|                command                                      |                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------- |
| `show dbs`                                                  | `show databases`                                      |
| `use <db_name>`                                             | `use <db_name>`                                       |
| `show collections`                                          | `show tables`                                         |
| `db.foo.find()`                                             | `select * from foo`                                   |
| `db.foo.find({id:1})`                                       | `select * from foo where id = 1`                      |
| `db.foo.find({id:1}).limit(10)`                             | `select * from foo where id = 1 limit 10`             |
| `db.foo.find({id:/123456/})`                                | `select * from foo where id like '%123456%'`          |
| `db.foo.count()`                                            | `select count(*) from foo`                            |
| `db.foo.update({},{$set:{aa:"bb"}},{multi:1})`              | 给foo集合所有记录添加一个新字段aa，值为bb                   |
| `db.foo.update({id:/123/},{$set:{aa:"bb"}},{multi:1})`      | 给foo集合id模糊匹配123的记录添加一个新字段aa，值为bb         |
| `db.foo.update({},{$set:{aa:NumberInt(0)}},{multi:1})`      | 添加字段aa，类型为Int，值为0                              |
| `db.foo.update({},{$unset:{aa:''}},{upsert:0,multi:1})`     | 删除字段aa                                              |



