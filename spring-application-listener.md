---
layout: post
title: Spring Application Listener
date: 2017-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---

# ApplicationListener定义
```java
@FunctionalInterface
public interface ApplicationListener<E extends ApplicationEvent> extends EventListener {
    void onApplicationEvent(E var1);
}
```


# ApplicationEvent
Spring提供如下几个内置的事件:        
* ContextRefreshedEvent: ApplicationContext容器初始化完成或刷新完成时，触发该事件。初始化完成指: 所有的bean被成功装载，后处理Bean被检测并激活，所有Singleton Bean被实例化。ApplicationContext容器已就绪可用。        
* ContextStartedEvent: 当使用ConfigurableApplicationContext(ApplicationContext的子接口）接口的start()方法启动ApplicationContext容器时触发该事件。容器管理声明周期的Bean实例将获得一个指定的启动信号，这在经常需要停止后重新启动的场合比较常见


