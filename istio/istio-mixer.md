---
layout: post
title: Istio Mixer
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---




Istio中每个请求，每个Envoy会调用两次Mixer:
1. 转发前，调用Mixer，进行前置检查
2. 转发后，调用Mixer，上报日志和监控数据








# Templates
Envoy -> Mixer -> adapter。不同的adapter接收不同类型的input数据来进行处理。一个logging adapter需要一个log数据，一个metric adapter需要一个metric数据。Istio使用Mixer template来描述adapter需要的具体数据。        











# Mixer的模型配置
Mixer的yaml配置可以抽象成三种模型:
1. **Handler**
2. **Instance**
3. **Rule**
这3种模型主要通过yaml中的kind字段进行区分。kind值有如下几种:
* [adapter kind](https://preliminary.istio.io/docs/reference/config/policy-and-telemetry/adapters/) : 表示此配置是Handler
* [template kind](https://preliminary.istio.io/docs/reference/config/policy-and-telemetry/templates/): 表示此配置是Template
* "rule":表示此为规则

以下面的yaml为例:
```yaml
apiVersion: "config.istio.io/v1alpha2"
kind: metric
metadata:
  name: requestsize
  namespace: istio-system
spec:
  value: request.size | 0
  dimensions:
    source_service: source.service | "unknown"
    source_version: source.labels["version"] | "unknown"
    destination_service: destination.service | "unknown"
    destination_version: destination.labels["version"] | "unknown"
    response_code: response.code | 200
  monitored_resource_type: '"UNSPECIFIED"'
```
kind是metric，那metric到底是个什么呢？Handler？Instance还是rule？在下面的链接中可以找到metric，它是一个Template。
[https://preliminary.istio.io/docs/reference/config/policy-and-telemetry/templates/metric/](https://preliminary.istio.io/docs/reference/config/policy-and-telemetry/templates/metric/)

> 我个人不理解为什么Istio要这么设计，为什么不在yaml中体现出这是一个template，不在yaml中出现template。而需要对Istio特别了解，通过查看相关的文档才知道metric是一个template。很费解。。。导致在谈论Istio模型的时候，很混乱。其实核心模型是三个:Handler, Instance, rule。结果导致yaml中有太多太多的kind。 

> 是不是可以这样？yaml中提供kind和subkind，kind=template，subkind=metric?至少这样比较清晰。



