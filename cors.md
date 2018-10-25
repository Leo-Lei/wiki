---
layout: post
title: Cross Origin Resource Sharing
date: 2016-07-14 15:40:00
tags:
- Java
categories: Java
---

# 1. Overview               
The CORS stands for Cross Origin Resource Sharing. In summary, it allows client send **AJAX(XMLHttpRequest)** resuest to another domain server.


Spring Boot中配置CORS跨站请求

注册一个Filter的Bean:

```java

@Configuration
public class WebConfiguration {

    @Bean
    public FilterRegistrationBean corsFilter(){

        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        CorsConfiguration config = new CorsConfiguration();
        config.setAllowCredentials(true);
        config.addAllowedOrigin("*");
        config.addAllowedHeader("*");
        config.addAllowedMethod("*");
        source.registerCorsConfiguration("/**", config);
        FilterRegistrationBean bean = new FilterRegistrationBean(new CorsFilter(source));
        bean.setOrder(0);
        return bean;
    }
}
```


在`src/main/resources/META-INF/spring.factories` 文件中添加Spring自动配置：
```java
org.springframework.boot.autoconfigure.EnableAutoConfiguration=\
com.alibaba.middleware.edas.service.platform.web.config.HelloWorldConfiguration,\
com.alibaba.middleware.edas.service.platform.web.config.WebConfiguration
```

注意：
* CORS Filter的Order要设置的尽量小，比如0。

#Spring-Boot #CORS

