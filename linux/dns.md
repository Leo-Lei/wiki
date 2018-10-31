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
