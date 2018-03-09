---
layout: post
title: Spring Lifecycle
date: 2017-11-19 20:40:00
tags:
- Java
categories: Java
---



|                            Step                                           |                 Desc              |
| ------------------------------------------------------------------------- | --------------------------------- |
| BeanFactoryPostProcessor.postProcessBeanFactory()                         |     只执行一次                      |
| Spring实例化Bean1完成                                                       |                                    |
| BeanPostProcessor.postProcessBeforeInitialization(bean1)                   |                                    |
| 调用Bean1的@PostConstruct方法                                                |                                   |
| 如果Bean1实现了InitializingBean，调用InitializingBean.afterPropertiesSet()   |                                    |
| 如果Bean1配置了init-method方法，调用init-method                               |                                    |
| BeanPostProcessor.postProcessAfterInitialization(bean1)                   |                                     |
| Spring实例化Bean2完成                                                       |                                     |
| BeanPostProcessor.postProcessBeforeInitialization(bean1)                   |                                    |
| 调用Bean1的@PostConstruct方法                                                |                                   |
| 如果Bean1实现了InitializingBean，调用InitializingBean.afterPropertiesSet()   |                                    |
| 如果Bean1配置了init-method方法，调用init-method                               |                                    |
| BeanPostProcessor.postProcessAfterInitialization(bean1)                   |                                     |
| Spring实例化所有Bean完成                                                     |                                     |
| 发布ContextRefreshedEvent事件                                               |                                     |

