---
layout: post
title: hangout
date: 2016-12-12 14:00:00
tags:
- Java
categories: Java
---

# hangout是什么？
hangout是Logstash的java版本。    
在使用`kafka`+`logstash`+`elasticsearch`的时候，kafka有消息堆积的情况，应该是logstash的吞吐量不够。    
logstash是ruby写的，后来尝试了下github上一个开源的hangout，文档上说性能是logstash的5倍。将我们的日志系统中的logstash替换成hangout之后，果然处理能力大了很多，消息也没有堆积了。    
下面是hangout的官网地址：    
[https://github.com/childe/hangout](https://github.com/childe/hangout)
