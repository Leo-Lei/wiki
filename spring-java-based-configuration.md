---
layout: post
title: Spring Java-based Configuration
date: 2016-07-18 16:45:00
tags:
- Java
categories: Java
description: web-authentication
---

# 1. Overview

```java
@Configuration
public class AppConfig {

    @Bean
    public MyService myService() {
        return new MyServiceImpl();
    }

}
```

```xml
<beans>
    <bean id="myService" class="com.acme.services.MyServiceImpl"/>
</beans>
```

