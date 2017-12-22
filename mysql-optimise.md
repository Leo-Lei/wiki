---
layout: post
title: MySql Optimise
date: 2017-06-25 12:00:00
tags:
- Gradle
categories: 
- Java
- Gradle
---




```sql
SELECT @@session.tx_read_only
SET autocommit=1
SET autocommit=0
```



# 资源
* [i-get-many-select-session-tx-read-only-would-one-do-the-same](https://stackoverflow.com/questions/32394729/i-get-many-select-session-tx-read-only-would-one-do-the-same/44043246#44043246)
* [http://www.iteye.com/topic/494179](http://www.iteye.com/topic/494179)
* [https://stackoverflow.com/questions/26053654/what-is-uselocalsessionstate-used-for](https://stackoverflow.com/questions/26053654/what-is-uselocalsessionstate-used-for)
* [http://www.tuicool.com/articles/emaEniv](http://www.tuicool.com/articles/emaEniv)
