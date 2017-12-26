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





# Reference
* [https://www.robustperception.io/life-of-a-label/](https://www.robustperception.io/life-of-a-label/)
* [https://www.robustperception.io/relabel_configs-vs-metric_relabel_configs/](https://www.robustperception.io/relabel_configs-vs-metric_relabel_configs/)
