---
layout: post
title: Shiro
date: 2017-04-27 11:10:00
tags:
- Security
categories: Java
---


|       Term           | Desc                                                  |
| -------------------- | ----------------------------------------------------- |
| `Realm`              |   DAO                                                 |
| `Subject`            | 用户或需要访问资源的进程                                  |





# Filter Chain

```properties
[urls]

/index.html = anon
/user/create = anon
/user/** = authc
/admin/** = authc, roles[administrator]
/rest/** = authc, rest
/remoting/rpc/** = authc, perms["remote:invoke"]
```

**First Match Wins**
```properties
/account/** = authc
/account/signup = anon
```



# Resources
[Shiro](http://shiro.apache.org)
