---
layout: post
title: Spring Component Scan
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---




```xml
<beans xmlns="http://www.springframework.org/schema/beans"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:context="http://www.springframework.org/schema/context"
	xsi:schemaLocation="http://www.springframework.org/schema/beans
	http://www.springframework.org/schema/beans/spring-beans-2.5.xsd
	http://www.springframework.org/schema/context
	http://www.springframework.org/schema/context/spring-context-2.5.xsd">

	<context:component-scan base-package="com.leibangzhu.foo.bar" />

</beans>
```


```java
@Configuration
@ComponentScan(basePackages = "org.example")
public class AppConfig  {
    ...
}
```


Spring自动扫描会扫描base-package下的类，如果类有如下的注解，则将它们注册为spring的bean。
* @Component,@Service,@Repository等
* @Controller



