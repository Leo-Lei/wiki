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



# 基本概念:
* Adapter：
业务自身的基础设施模块（比如Prometheus、ACL控制等）需要接入Istio，需要在Mixer组件中开发相应的adapter插件，这个adapter插件负责承接Mixer流转过来的运行时信息，然后以各自的协议与相应的基础设施模块进行交互。

* Handler：
每个adapter需要配置一些参数才能使用，比如配置后端的链接URL、证书、缓存等参数（配置参数用protobuf格式描述）。每个处于配置完成状态的adapter被称为一个handler，可以对同一个adapter配置多次，对应成多个handler实例（除了配置不同以外，每个handler都有自己的name和namespace），在不同的场景下使用。无论是CHECK还是REPORT请求，Mixer会调用一个或多个handler实例，具体看Rule Configurer的配置。

* Template：
不同的adapter的处理需要不同格式的输入数据。这些格式信息由Template资源来描述。Adapter在设计阶段可以指定依赖多个Template资源。总结Template的作用：
a). 定义adapter组件需要处理的数据格式。
b). 定义接口规范，这个接口能够识别并处理template定义的数据格式，依赖这个Template的adapter组件需要实现这个接口规范。

* Instance：
如前所述，传输给adapter处理的数据格式需要由Template资源来描述，具体过程：在运行时，Mixer将Envoy上报的属性数据（attributes）按照Template的格式以及Rule Configurer配置（指定哪些填充字段）封装成instance对象，然后将instance对象传输给依赖这种Template的adapter实例（handler对象），instance可以理解为template的实例（Rule Configurer在配置时通过kind指定为template的name来关联）。

* Rule：
将哪些instances数据传给哪个handler实例，是通过创建rule来指定的。rule规则需要指定某个handler以及需要发送给这个handler的一系列的instances。rule还需要指定匹配规则，上传的attributes需要满足这个匹配规则才会执行将instances传给handler处理的操作。


# Template
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

# Handler
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

# Instance
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
