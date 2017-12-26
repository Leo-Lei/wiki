---
title: Prometheus 2.0 Migration
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---

# Prometheus
从1.8到2.0，Prometheus有一些向后不兼容的改动。在从1.8升级到2.0的时候，需要注意。

# 命令行参数
* `-alertmanager.url`: 在2.0中，该参数已被移除。不支持从命令行通过该参数配置一个静态的alertmanager。必须要通过服务发现机制被发现。
* `-storage.local.*`: 在2.0中，被移除。2.0使用了一个新的存储引擎。








# 官方说明
[https://prometheus.io/docs/prometheus/2.0/migration](https://prometheus.io/docs/prometheus/2.0/migration/)