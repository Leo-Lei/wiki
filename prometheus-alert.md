---
title: Prometheus Alert
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---


# 报警配置
```yml
# How frequently to evaluate rules.
[ evaluation_interval: <duration> | default = 1m ]
```


# 定义报警规则
```text
ALERT <alert name>
  IF <expression>
  [ FOR <duration> ]
  [ LABELS <label set> ]
  [ ANNOTATIONS <label set> ]
```
通过下面的方式添加变量
```text
# To insert a firing element's label values:
{{ $labels.<labelname> }}
# To insert the numeric expression value of the firing element:
{{ $value }}
```

```text
# Alert for any instance that is unreachable for >5 minutes.
ALERT InstanceDown
  IF up == 0
  FOR 5m
  LABELS { severity = "page" }
  ANNOTATIONS {
    summary = "Instance {{ $labels.instance }} down",
    description = "{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 5 minutes.",
  }

# Alert for any instance that have a median request latency >1s.
ALERT APIHighRequestLatency
  IF api_http_request_latencies_second{quantile="0.5"} > 1
  FOR 1m
  ANNOTATIONS {
    summary = "High request latency on {{ $labels.instance }}",
    description = "{{ $labels.instance }} has a median request latency above 1s (current value: {{ $value }}s)",
  }

```


