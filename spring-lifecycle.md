---
layout: post
title: Spring Lifecycle
date: 2017-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---



|                            Step                          |                         Desc                          |
| -------------------------------------------------------- | ----------------------------------------------------- |
| BeanFactoryPostProcessor.postProcessBeanFactory()        | 初始化BeanFactoryPostProcessor                         |
| BeanPostProcessor.postProcessBeforeInitialization()      |                                                       |
| @PostConstruct                                           |                                                       |
| InitializingBean.afterPropertiesSet()                    |                                                       |
| init-method方法                                           |                                                       |
| BeanPostProcessor.postProcessAfterInitialization()       |                                                       |
| 发布ContextRefreshedEvent事件                             |                                                       |
