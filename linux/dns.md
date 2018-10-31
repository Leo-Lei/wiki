---
layout: post
title: DNS
date: 2018-10-30 11:50:00
tags:
- Java
categories: Linux
---

# DNS使用的端口
DNS使用的是`53`端口。    
通常DNS查询时，是以udp这个叫快速的数据传输协议来查询的，但是万一没有查询到完整信息时，就会再以tcp协议来重新查询。所以启动DNS的daemon时，会同时启动tcp即udp的53端口。

# DNS层级架构

![DNS层级示意图](http://cn.linux.vbird.org/linux_server/0350dns_files/dns_dot.gif)
在整个DNS系统的最上方一定是`.`(小数点)，这个DNS服务器称为root服务器。它下面管理的域名有com，org，edu，gov，net，cn，jp，tw，uk等。这些域名是顶级域名。

# 通过DNS查询IP的流程
DNS以类似于树状目录的形态来进行主机名的管理。所以，每一步DNS服务器都仅管理自己的下一层主机名的解析。至于下层的下层，则授权给下层的DNS来管理。
下面就来举例说明：
首先，在浏览器中输入http://www.ksu.edu.tw时，计算机会依据相关设定(在Linux中是利用/etc/resolv.conf这个文件)所提供的DNS去进行查询。
假设`/etc/resolve.conf`中配置了:
```bash
nameserver 168.95.1.1
```

1. 收到用户的查询请求，先查看本地有没有记录，若无，向root查询。
2. 向root查询
    root只记录了.tw的信息，此时root会告知我不知道这部主机的IP，不过，你应该向.tw去询问。
3. 向第二层的.tw服务器查询
    向.tw查询，.tw仅管理了.edu.tw,.com.tw,gov.tw这几部主机。此时.tw会告知，我不知道这部主机的IP，不过，你应该向.edu.tw查询。
4. 向第三层的.edu.tw查询
    .edu.tw会告知，应该向.ksu.edu.tw查询。这里只能告知.ksu.edu.tw的IP。
5. 向第四层的.ksu.edu.tw查询
    .ksu.edu.tw说，没错，这部主机是我管理的。我告诉你，它的IP是XXXXX
6. 查到正确的IP后，168.95.1.1的DNS机器不会再下次有人查询www.ksu.edu.tw的时候，再绕一大圈。所以，168.95.1.1这个DNS会缓存一份查询结果在内存中，以方便响应下一次的相同请求。缓存是有时效性的，通常可能是24小时。
