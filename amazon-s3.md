---
layout: post
title: AWS S3
date: 2017-07-02 11:10:00
tags:
- docker
categories: Java
description: AWS Route 53
---


# S3上托管静态站点，并使用自定义域名
假设有域名abc.com，但是该域名不在AWS的Route 53上管理，比如是在aliyun上管理的。
现在想在AWS的S3上托管一个静态站点，并使用console.abc.com来访问。可以做如下的操作：
1. 在S3上创建存储桶，存储桶的名字必须是和域名一样，是console.abc.com。
2. 将静态站点的文件都上传至S3的console.abc.com存储桶中
3. 设置console.abc.com存储桶可以公共读
4. 编辑console.abc.com存储桶的属性，启动静态站点托管功能，并设置站点首页。
5. 在Route 53上建立一个域abc.com。AWS会自动为该域创建一些DNS服务器。
6. 在Route 53上建立一个A记录，alias选择Yes，alias type选择S3上对应的存储桶。
7. 进入aliyun控制台，为abc.com添加一个NS记录，记录值指向AWS的Route 53的DNS服务器地址。
8. 等待几分钟，让改动生效，然后在浏览器中输入console.abc.com就可以访问托管在Route53上的静态站点了

# Resources
[http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html](http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html)
[http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/MigratingDNS.html](http://docs.aws.amazon.com/zh_cn/Route53/latest/DeveloperGuide/MigratingDNS.html)        
[website-hosting-custom-domain-walkthrough.html](http://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/website-hosting-custom-domain-walkthrough.html#root-domain-walkthrough-update-ns-record)
