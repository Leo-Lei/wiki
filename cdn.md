---
layout: post
title: CDN
date: 2017-05-16 14:30:00
tags:
- docker
categories: Java
description: docker
---


# CDN
Content Distribution Network.

# CDN加速静态内容
CDN加速静态内容比较好理解，将静态资源缓存在多个节点中。用户请求时，从最近的一个节点获取静态资源。


# CDN加速动态内容
1. 优化访问线路。对于动态内容，用户最终的请求，都会到指定的主机上。部分地理位置较远，访问网站路径较长的，可能会出现故障。比如，主机在A，用户在C，A和C之间相距很远。用户直接访问网站，路径可能会很曲折，导致网站访问速度便慢。使用CDN加速，在A和C之间增加一个节点C，节点C长期存在于互联网上，不管时B访问C还是C访问B，速度都不慢，加入C节点后，可以让用户C访问A速度加快。




