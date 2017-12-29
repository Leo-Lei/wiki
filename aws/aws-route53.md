---
layout: post
title: AWS Route 53
date: 2017-07-02 11:10:00
tags:
- docker
categories: Java
---


# Route 53



# 转移域名到Route 53
迁移顶级域名，比如`.com`,`.org`到Route 53是有一些要求的。每个顶级域名的要求会有些差异，但下面的这些要求是比较有代表性的：
* 域名在当前注册中心要注册了超过60天
* 域名不要过期
* 域名下不能有如下状态的记录：
pengingDelete
pendingTransfer
redemptionPeriod
clientTransferProhibited


## 1. 确认Route 53 支持待转移的顶级域名
## 2. 转移DNS服务到Route 53








# S3 + Route 53实现自定义域托管静态站点

在S3上创建一个存储桶，将其托管为一个静态站点，这样，S3会为存储桶生成一个终端节点，类似于：
`http://www.example.com.s3-website-us-west-2.amazonaws.com/#/`。在浏览器中就可以通过该url来访问静态站点了。






# Resources
[http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html](http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html)
[http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/MigratingDNS.html](http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/MigratingDNS.html)        
[website-hosting-custom-domain-walkthrough.html](http://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/website-hosting-custom-domain-walkthrough.html#root-domain-walkthrough-update-ns-record)
