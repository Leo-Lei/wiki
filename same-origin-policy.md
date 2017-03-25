---
layout: post
title: Same Origin Policy
date: 2016-07-15 11:15:00
tags:
- Java
categories: Java
description: web-authentication
---

# 1. Overview               
The same origin is:
* Same protocal
* Same domain
* Same port

Take the following URL for example, the protocal is *http*, the domain is *store.company.com*, the port is *80*.
http://store.company.com/dir/page.html

|                  URL                                 |  Result |        Reason             |
| ---------------------------------------------------- | ------- | ------------------------- |
| http://store.company.com/dir2/other.html             |    Y    |                           |
| http://store.company.com/dir/inner/another.html      |    Y    |                           |
| https://store.company.com/secure.html                |    N    |  Different protocal       |
| http://store.company.com:81/dir/etc.html             |    N    |  Different port           |
| http://news.company.com/dir/other.html               |    N    |  Different host           |



