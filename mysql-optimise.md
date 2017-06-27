---
layout: post
title: MySql Optimise
date: 2017-06-25 12:00:00
tags:
- Gradle
categories: 
- Java
- Gradle
description: MySql
---




```sql
SELECT @@session.tx_read_only
SET autocommit=1
SET autocommit=0
```



# 资源
[i-get-many-select-session-tx-read-only-would-one-do-the-same](https://stackoverflow.com/questions/32394729/i-get-many-select-session-tx-read-only-would-one-do-the-same/44043246#44043246)
