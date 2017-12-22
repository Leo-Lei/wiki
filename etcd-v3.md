---
layout: post
title: Etcd
date: 2015-06-30 15:50:00
tags:
- Atom
categories: Text Editor
---

# 1. Etcd V3

```bash
export ETCDCTL_API=3
```

API V3和V2的一些区别:
* 事务：ETCD V3提供了多键条件事务（multi-key conditional transactions），应用各种需要使用事务代替原来的Compare-And-Swap操作。    
* 平键空间（Flat key space）：ETCD V3不再使用目录结构，只保留键。例如：”/a/b/c/“是一个键，而不是目录。V3中提供了前缀查询，来获取符合前缀条件的所有键值，这变向实现了V2中查询一个目录下所有子目录和节点的功能。在V2中,`put abc 100`会创建一个`/abc`,值为100，但V3中只会创建一个`abc`的key。    
* 简洁的响应：像DELETE这类操作成功后将不再返回操作前的值。如果希望获得删除前的值，可以使用事务，来实现一个原子操作，先获取键值，然后再删除。    
* 租约：租约代替了V2中的TTL实现，TTL绑定到一个租约上，键再附加到这个租约上。当TTL过期时，租约将被销毁，同时附加到这个租约上的键也被删除。     


# 访问远程etcd服务器上的数据
`etcdctl --endpoints=http://192.168.100.200:2379 get / --prefix`

|              Command                 |                                            |
| ------------------------------------ | ------------------------------------------ |
| `put abc 100`                        |                                            |
| `put /foo/bar 100`                   |                                            |
| `del abc`                            |                                            |
| `get abc`                            |                                            |
| `get /foo/bar`                       |                                            |
| `get /foo` --prefix                  |                                            |
| `watch abc`                          |                                            |
| `watch /foo` --prefix                |                                            |
