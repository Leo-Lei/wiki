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
* "rule":表示是规则

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



# Mixer Handler
一个Handler是一个配置好的Adapter实例。Handler从yaml配置文件中取出Adapter需要的配置数据。一个典型的Prometheus Handler配置如下所示：
```yaml
apiVersion: config.istio.io/v1alpha2
kind: prometheus
metadata:
  name: handler
  namespace: istio-system
spec:
  metrics:
  - name: request_count
    instance_name: requestcount.metric.istio-system
    kind: COUNTER
    label_names:
    - destination_service
    - destination_version
    - response_code
```

Handler的完全名称是{metadata.name}.{kind}.{metadata.namespace}。是全局唯一的。    
每个adapter配置的格式都不一样，可以在[这里](https://istio.io/docs/reference/config/policy-and-telemetry/adapters/)查看相关配置。上述Handler中引用了requestduration.metric.istio-system这个Instance。

# Mixer Instance
Instance定义了attributes到adapter输入的映射，一个处理requestduration metric数据的Instance配置如下:
```yaml
apiVersion: config.istio.io/v1alpha2
kind: metric
metadata:
  name: requestduration
  namespace: istio-system
spec:
  value: response.duration | "0ms"
  dimensions:
    destination_service: destination.service | "unknown"
    destination_version: destination.labels["version"] | "unknown"
    response_code: response.code | 200
  monitored_resource_type: '"UNSPECIFIED"'
```
上述Instance的完全限定名是requestduration.metric.istio-system，Handler和Rule可以通过这个名称对此Instance进行引用。

# Rule
Rule定义了一个特定的Instance什么情况下调用一个特定的Handler，一个典型的Rule配置如下所示:
```yaml
apiVersion: config.istio.io/v1alpha2
kind: rule
metadata:
  name: promhttp
  namespace: istio-system
spec:
  match: destination.service == "service1.ns.svc.cluster.local" && request.headers["x-user"] == "user1"
  actions:
  - handler: handler.prometheus
    instances:
    - requestduration.metric.istio-system
```
上述例子中，定义了这样一个Rule:        
当目标是service1，header中的x-user=user1时，把requestduration的metric发送到prometheus的handler来处理。


