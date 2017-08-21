---
layout: post
title: lucene
date: 2016-12-12 14:10:00
tags:
- Java
categories: Java
description: Lucene
---

# Lucene分词

|                type            |                                                                           |
| ------------------------------ | ------------------------------------------------------------------------- |
| `Field.Index.ANALYZED`         | 建立索引，并进行分词                                                         |
| `Field.Index.NOT_ANALYZED`     | 建立索引，当不进行分词。即不使用analyzer分析，将整体作为一个token，常用于精确匹配    |

