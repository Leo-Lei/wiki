---
layout: post
title: API vs SPI
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---

**API** stand for Application Programming Interface. **SPI** stands for Service Provider Interface.

Put differently, the API tells you what a specific class/method does for you and the SPI tells you what you must do to conform. So:

* the **API** is the description of classes/interfaces/methods/... that you **call and use** to achieve a goal
* the **SPI** is the description of classes/interfaces/methods/... that you **extend and implement** to achieve a goal

Usually API and SPI are separate. For example in JDBC the Driver class is part of the SPI: if you simply want to use JDBC, you don't need to use it directly, but everyone who implements a JDBC driver must implement that class.      
The JDBC only define a specification for accessding Data Base via Java language. The JDBC specification is a collection of interface, with no concrete implementation. The JDBC driver provider will implement the interface. For example, SQL Server, MySql, Oracle or others.        
Sometimes they overlap, however. The [Connection](http://java.sun.com/javase/6/docs/api/java/sql/Connection.html) interface is both SPI and API: You use it routinely when you use a JDBC driver and it needs to be implemented by the developer of the JDBC driver.

