---
layout: post
title: elastic search
date: 2016-09-30 14:20:00
tags:
- Java
categories: Java
description: elastic search
---

# Elasticsearch是什么？    
Elasticsearch是java开发的，基于Lucene的开源搜索引擎。 Elasticsearch的特点：    

1. 基于java的开源全文搜索引擎。基于Lucene。    
2. 通过简单的RESTful API来隐藏Lucene的复杂性，让搜索变得简单。        
3. 可以扩展到上百台机器，处理PB级别的结构化和非结构化数据。    
4. 分布式的实时分析搜索引擎。    

# 安装Elasticsearch
**前提条件**    

1. Elastic search是java开发的，所以需要安装jdk。 

**tar安装**    
从官网[elasticsearch.org/download](http://www.elasticsearch.org/download) 下载.是以zip包的形式发布的，文件格式为`elasticsearch-$VERSION.zip`。
**apt/yum/homebrew安装**    

* Mac／Homebrew安装： `brew install elasticsearch`     

# 运行/停止Elasticsearch
**前台运行:**    
运行`./bin/elasticsearch`可以在前台运行Elasticsearch。        
**后台运行:**    
添加`-d`参数在后台以守护进程模式运行。    
```bash
curl -L -O http://download.elasticsearch.org/PATH/TO/VERSION.zip <1>
unzip elasticsearch-$VERSION.zip
cd  elasticsearch-$VERSION
./bin/elasticsearch
```
**检查elasticsearch是否已运行**    
运行`curl 'http://localhost:9200/?pretty'`,会返回如下的信息：
```json
{
   "status": 200,
   "name": "Shrunken Bones",
   "version": {
      "number": "1.4.0",
      "lucene_version": "4.10"
   },
   "tagline": "You Know, for Search"
}
```    
说明Elasticsearch集群已经启动并运行正常。    
### 集群和节点        
节点(node)是一个运行着的Elasticsearch实例。集群是一组具有相同`cluster.name`的节点集合，它们协同工作，共享数据并提供fail over和扩展功能。一个节点也可以组成一个集群。可以在`config/elasticsearch.yml`文件中修改cluster.name，修改完后需要重启。
 
### 关闭elasticsearch
但elasticserch在前台运行时，用`Ctrl` + `C`来关闭elasticsearch。

# 索引(index)/文档(document)      
```text
Relation DB    -> Databases -> Tables -> Rows     -> Columns
Elasticsearch  -> Index     -> Types  -> Documents -> Fields
```


# Resources
http://es.xiaoleilu.com/010_Intro/10_Installing_ES.html


