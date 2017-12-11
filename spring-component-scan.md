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

```<context:component-scan />```隐式的启用了```<context:annotation-config>```。当使用了```<context:annotation-config>```的时候，就没必要使用```<context:component-scan />```了。      
```java
@Configuration
@ComponentScan(basePackages = "org.example")
public class AppConfig  {
    ...
}
```

# 自定义扫描
默认情况下，Spring只扫描如下的两种类，并把它们注册到Spring容器中:        
1. @Component, @Repository, @Service, @Controller
2. 自定义的注解，并且被@Component标注了

### includeFilters, excludeFilters
```java
@Configuration
@ComponentScan(basePackages = "org.example",
        includeFilters = @Filter(type = FilterType.REGEX, pattern = ".*Stub.*Repository"),
        excludeFilters = @Filter(Repository.class))
public class AppConfig {
    ...
}
```

* 如果不希望Spring自动扫描@Component,@Repository,@Service或@Controller等，可以设置```useDefaultFilters=false```






