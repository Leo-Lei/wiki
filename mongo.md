---
layout: post
title: Mongo
date: 2017-04-11 13:05:00
tags:
- Java
categories: Java
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
| `db.foo.getIndexes()`                                       | 显示collection的索引                                    |
| `db.createCollection("mycollection")`                       | 创建集合                                                |



# mongo添加二维索引
```bash
db.device.ensureIndex({loc:"2dsphere"})
```

# 停止mongo
```bash
use admin
db.shutdownServer()
```






# Mongo使用用户名和密码

MongoDB 缺省是没有设置鉴权的，业界大部分使用 MongoDB 的项目也没有设置访问权限。这就意味着只要知道 MongoDB 服务器的端口，任何能访问到这台服务器的人都可以查询和操作 MongoDB 数据库的内容。在一些项目当中，这种使用方式会被看成是一种安全漏洞。

本文介绍如何在单台 MongoDB 服务器上设置鉴权。设置完后，MongoDB 客户端必须用正确的用户名和密码登录，才能在指定的数据库中操作。

首先介绍下 MongoDB 的用户和权限。每个数据库都有自己的用户，创建用户的命令是db.createUser()（文档），当你创建一个用户时，该用户就属于你当前所在的数据库。

每个用户包含三个要素：用户名、密码和角色列表。下面是一个例子：


{
user: "dbuser",
pwd : "dbpass",
roles: ["readWrite", "clusterAdmin"]
}

这个例子表示一个名为dbuser的用户，它在当前的数据库中拥有 readWrite 和 clusterAdmin 两个角色。

--------------------------------------------------------------------------------------------------------------------------


MongoDB 内置了很多角色，但要注意，不是每个数据库的内置角色都一样。其中 admin 数据库就包含了一些其他数据库所没有的角色。

熟悉 Oracle 的童鞋们都知道，数据库用户有两种，一种是管理员，用来管理用户，一种是普通用户，用来访问数据。类似的，为 MongoDB 规划用户鉴权时，至少要规划两种角色：用户管理员和数据库用户。如果搭建了分片或主从，可能还会要规划数据库架构管理员的角色，它们专门用来调整数据库的分布式架构。

在创建用户之前，我们首先要修改 MongoDB 的启动方式。缺省方式下 MongoDB 是不进行鉴权检查的。我们只要在运行 MongoDB 的命令后面加上一个 --auth 参数即可，例如：

mongod --dbpath ./db1 --port 20000 --auth

 

如何创建用户管理员


用户管理员是第一个要创建的用户。在没有创建任何用户之前，你可以随意创建用户；但数据库中一旦有了用户，那么未登录的客户端就没有权限做任何操作了，除非使用db.auth(username, password)方法登录。

用户管理员的角色名叫 userAdminAnyDatabase，这个角色只能在 admin 数据库中创建。下面是一个例子：

> use admin
switched to db admin
> db.createUser({user:"root",pwd:"root123",roles:["userAdminAnyDatabase"]})
Successfully added user: { "user" : "root", "roles" : [ "userAdminAnyDatabase" ] }

这个例子创建了一个名为 root 的用户管理员。创建完了这个用户之后，我们应该马上以该用户的身份登录：

> db.auth("root","root123")
1
db.auth() 方法返回 1 表示登录成功。接下来我们为指定的数据库创建访问所需的账号。

--------------------------------------------------------------------------------------------------------------------------

如何创建数据库用户


首先保证你已经以用户管理员的身份登录 admin 数据库。然后用 use 命令切换到目标数据库，同样用 db.createUser() 命令来创建用户，其中角色名为 “readWrite”。

普通的数据库用户角色有两种，read 和 readWrite。顾名思义，前者只能读取数据不能修改，后者可以读取和修改。
下面是一个例子：

> use test
switched to db test
> db.createUser({user:"testuser",pwd:"testpass",roles:["readWrite"]})
Successfully added user: { "user" : "testuser", "roles" : [ "readWrite" ] }
> db.auth("testuser","testpass")
1

这样 MongoDB 的数据安全性就得到保障了，没有登录的客户端将无法执行任何命令。




