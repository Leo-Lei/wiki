---
layout: post
title: Spring Autowiring
date: 2016-11-19 20:40:00
tags:
- Java
categories: Java
description: spring
---



|      annotation     |                            desc                                                 |
| ------------------- | ------------------------------------------------------------------------------- |
| `@Primary`          | 在@Autowired的时候，如果一个type有多个实现，可以通过@Primary，告诉Spring优先选择该组件。  |
| `@Qualifier`        | 当@Autowired时，如果有多个bean，通过@Qualifier指定bean的id来注入                      |

