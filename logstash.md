---
layout: post
title: Logstash
date: 2016-11-12 14:00:00
tags:
- Java
categories: Java
---

# Logstash是什么？
Logstash是一个event和log的采集工具。可以从不同的datasource采集数据，处理／过滤，然后，输出到不同的destination。    
![logstash workflow](https://www.elastic.co/guide/en/logstash/current/static/images/basic_logstash_pipeline.png)


# 安装Logstash        
推荐使用apt／yum／homebrew来安装，也可以去官网下载tar文件来安装。我是用Mac上的homebrew来安装的。    
`brew install logstash`    
安装完之后，logstash的可执行文件在`/usr/local/bin/logstash`中，已经在PATH里面了。可以直接使用`logstash`命令了。    
**测试logstash是否安装正确**        
在终端输入`logstash -e 'input { stdin { } } output { stdout {} }'`，会启动logstash，监控终端的输入，然后将它输出到终端。等待logstash打印出"Pipeline main started",在终端输入hello，logstash会在终端输出`2016-11-12T06:08:15.159Z leiweideMacBook-Pro.local hello world`.      
说明logstash已安装正确。用`CTRL` + `D`退出logstash。    
# 运行Logstash    
|            command              |                              |                                                          |
| ------------------------------- | ---------------------------- | -------------------------------------------------------- |
| `logstash -e <config>`          | 在命令行中指定配置              | `logstash -e 'input { stdin { } } output { stdout {} }'` |
| `logstash -f <config_file>`     | 指定配置文件                   | `logstash -f /etc/logstash/logstash.conf`                |
| `logstash -l <log_file>`        | 指定日志文件                   | `logstash -l /etc/logstash/logstash.log`                 |
| `logstash -f <log_file> &`      | `&`让logstash在后台运行        | `logstash -l /etc/logstash/logstash.log &`               |
