---
layout: post
title: AWS Route 53
date: 2017-07-02 11:10:00
tags:
- docker
categories: Java
description: AWS Route 53
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








# Resources
[http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html](http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html)
[http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/MigratingDNS.html](http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/MigratingDNS.html)
[website-hosting-custom-domain-walkthrough.html](http://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/website-hosting-custom-domain-walkthrough.html#root-domain-walkthrough-update-ns-record)
