---
layout: post
title: Spring Web MVC
date: 2016-09-23 17:55:00
tags:
- Java
categories: Java
description: spring web mvc
---

# Start to use Spring Web MVC

```xml
<web-app>

    <servlet>
        <servlet-name>example</servlet-name>
        <servlet-class>org.springframework.web.servlet.DispatcherServlet</servlet-class>
        <load-on-startup>1</load-on-startup>
    </servlet>

    <servlet-mapping>
        <servlet-name>example</servlet-name>
        <url-pattern>/</url-pattern>
    </servlet-mapping>

</web-app>
```
Spring MVC framework will find the file `<servlet-name>-servlet.xml` file to initialize the **DispatchServlet**.



