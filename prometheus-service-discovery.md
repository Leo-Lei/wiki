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
* keep: 删除source_labels不匹配regex的Target
* drop: 删除source_labels匹配regex的Target

