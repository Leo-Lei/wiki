---
layout: post
title: SQL Server Basis
date: 2015-09-24 11:30:00
tags:
- SQL
categories: SQL
---

# 1. Check the version and edition of SQL Server Database Engine

```sql
Select @@version
```

```plain
Microsoft SQL Server 2012 (SP1) - 11.0.3460.0 (X64)   Jul 22 2014 15:22:00   
Copyright (c) Microsoft Corporation  Standard Edition (64-bit) on Windows NT 6.1 <X64> (Build 7601: Service Pack 1) (Hypervisor) 
```



# 2. Change SQL Server authentication mode

**Using Transact-SQL:**

```sql
ALTER LOGIN sa ENABLE ;
GO
ALTER LOGIN sa WITH PASSWORD = '<enterStrongPasswordHere>' ;
GO
```


