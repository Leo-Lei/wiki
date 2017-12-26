---
title: Prometheus Service Discovery
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---


# relabel configs
<relabel_action>:
* replace: 使用正则表达式中匹配的组来替换source_labels的值。正则中匹配的组是${1},${2},${3}...
* keep: 删除source_labels不匹配regex的Target
* drop: 删除source_labels匹配regex的Target

# 服务发现和relabel流程
* 通过服务发现，发现一些Target。各个组件的服务发现机制是不一样的，配置也不一样。比如kubernetes的服务发现配置中需要配置自动发现的kubernetes中的角色，有node，service，pod，endpoint等。
* 如果Target没有`__address__`，删除Target
* 执行relabel_configs，这个过程中会通过keep，drop删除一些Target。通过replace等来替换一些label的值。
* 如果`__address__`没有端口号，根据协议添加默认端口，http是80，https是443.
* 删除所有`__meta_`的label
* 如果`instance`这个label不存在，将`__address`这个label的值给instance
* 创建Target


# Reference
* [https://www.robustperception.io/life-of-a-label/](https://www.robustperception.io/life-of-a-label/)
* [https://www.robustperception.io/relabel_configs-vs-metric_relabel_configs/](https://www.robustperception.io/relabel_configs-vs-metric_relabel_configs/)
