---
layout: post
title: Envoy API
date: 2017-06-16 14:35:00
tags:
- docker
categories: Java
---

# Clusters
### Cluster
Cluster代表一个上游的集群。一个集群中包含多个提供相同功能的Host。Cluster内部通过LoadBalance来选出一个Host。    
Cluster中包含的字段如下:
```json
{
  "name": "...",
  "alt_stat_name": "...",
  "type": "...",
  "eds_cluster_config": "{...}",
  "connect_timeout": "{...}",
  "per_connection_buffer_limit_bytes": "{...}",
  "lb_policy": "...",
  "hosts": [],
  "load_assignment": "{...}",
  "health_checks": [],
  "max_requests_per_connection": "{...}",
  "circuit_breakers": "{...}",
  "tls_context": "{...}",
  "common_http_protocol_options": "{...}",
  "http_protocol_options": "{...}",
  "http2_protocol_options": "{...}",
  "extension_protocol_options": "{...}",
  "dns_refresh_rate": "{...}",
  "dns_lookup_family": "...",
  "dns_resolvers": [],
  "outlier_detection": "{...}",
  "cleanup_interval": "{...}",
  "upstream_bind_config": "{...}",
  "lb_subset_config": "{...}",
  "ring_hash_lb_config": "{...}",
  "original_dst_lb_config": "{...}",
  "common_lb_config": "{...}",
  "transport_socket": "{...}",
  "metadata": "{...}",
  "protocol_selection": "...",
  "upstream_connection_options": "{...}",
  "close_connections_on_host_health_failure": "...",
  "drain_connections_on_host_removal": "..."
}
```

|      字段                          |                类型               |                                          描述                                           |
| --------------------------------- | --------------------------------- | -------------------------------------------------------------------------------------- |
| `name`                            | `string`                          | Cluster的名字。必须全局唯一。                                                              |
| `alt_stat_name`                   | `string`                          | `alt`是`alternative`的意思。`stat`是`statistics`的意思。                                   | 
| `type`                            | `Cluster.DiscoveryType`           | DiscoveryType，即服务发现的类型。有STATIC, STRICT_DNS, LOGICAL_DNS, EDS, ORIGINAL_DST。    | 
| `eds_cluster_config`              | `Cluster.EdsClusterConfig`        | Endpoint Discovery Service配置                                                          |







### Cluster.EdsClusterConfig
只有服务发现类型是EDS的时候，才需要改配置
```json
{
  "eds_config": "{...}",
  "service_name": "..."
}
```


# core
### core.ConfigSource

```json
{
  "path": "...",
  "api_config_source": "{...}",
  "ads": "{...}"
}
```

|            字段              |               类型              |                描述               |
| --------------------------- | ------------------------------ | --------------------------------- |
| `path`                      | `string`                       | 配置文件在文件系统中的路径。           |
| `api_config_source`         | `core.ApiConfigSource`         | API Configuration source。         |
| `ads`                       | `core.AggregatedConfigSource`  |                                   |


### core.Node





# Bootstrap

### config.bootstrap.v2.Bootstrap
完整的Bootstrap配置如下:
```json
{
  "node": "{...}",
  "static_resources": "{...}",
  "dynamic_resources": "{...}",
  "cluster_manager": "{...}",
  "hds_config": "{...}",
  "flags_path": "...",
  "stats_sinks": [],
  "stats_config": "{...}",
  "stats_flush_interval": "{...}",
  "watchdog": "{...}",
  "tracing": "{...}",
  "rate_limit_service": "{...}",
  "runtime": "{...}",
  "admin": "{...}",
  "overload_manager": "{...}"
}
```

|            字段          |                           类型                    |                                          说明                                 |
| ----------------------- | ------------------------------------------------- | ---------------------------------------------------------------------------- |
| `node`                  | `core.Node`                                       |                                                                              |
| `static_resources`      | `config.bootstrap.v2.Bootstrap.StaticResources`   | 静态资源                                                                      |
| `dynamic_resources`     | `config.bootstrap.v2.Bootstrap.DynamicResources`  | xDS配置源                                                                     |
| `cluster_manager`       | `config.bootstrap.v2.ClusterManager`              | cluster manager的配置                                                         |
| `hds_config`            | `core.ApiConfigSource`                            | Health Discovery Service配置                                                  |
| `flags_path`            | `string`                                          | 指定一个文件路径，用于查找启动时的flag                                             |
| `stats_sinks`           | `config.metrics.v2.StatsSink`                     | stats是Statistics缩写                                                          |
| `stats_config`          | `config.metrics.v2.StatsConfig`                   | stats是Statistics缩写                                                          |
| `stats_flush_interval`  | `Duration`                                        | 发送统计信息的时间间隔，默认5秒钟                                                   |
| `watchdog`              | `config.bootstrap.v2.Watchdog`                    | watchdog配置                                                                   |
| `tracing`               | `config.trace.v2.Tracing`                         | 外部的tracing配置，如果没指定，不执行tracing                                       |
| `rate_limit_service`    | `config.ratelimit.v2.RateLimitServiceConfig`      | 外部的rate limit配置。如果没指定，所有对rate limit service的调用立即返回success       |
| `runtime`               | `config.bootstrap.v2.Runtime`                     | Runtime配置                                                                    |
| `admin`                 | `config.bootstrap.v2.Admin`                       | local adminnistration HTTP server的配置                                        |
| `overload_manager`      | `config.overload.v2alpha.OverloadManager`         | overload manager配置                                                           |

### config.bootstrap.v2.Bootstrap.StaticResources

```json
{
  "listeners": [],
  "clusters": [],
  "secrets": []
}
```

|        字段            |               类型                  |                                  说明                                      |
| --------------------- | ----------------------------------- | --------------------------------------------------------------------------- |
| `listeners`           | `Listener`                          | Listener配置                                                                 |
| `clusters`            | `Cluster`                           | Cluster配置                                                                  |
| `secret`              | `auth.Secret`                       | Secret配置                                                                   |


### config.bootstrap.v2.Bootstrap.DynamicResources

```json
{
  "lds_config": "{...}",
  "cds_config": "{...}",
  "ads_config": "{...}"
}
```

|        字段            |               类型                  |                                  说明                                        |
| --------------------- | ----------------------------------- | --------------------------------------------------------------------------- |
| `lds_config`          | `core.ConfigSource`                 | Listener Discovery Service配置                                               |
| `cds_config`          | `core.ConfigSource`                 | Cluster Discovery Service配置                                                |
| `ads_config`          | `core.ApiConfigSource`              | Aggregated Discovery Service配置                                             |


### config.bootstrap.v2.Admin

```json
{
  "access_log_path": "...",
  "profile_path": "...",
  "address": "{...}"
}
```

|        字段            |         类型          |                                                             说明                                                             |
| --------------------- | -------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `access_log_path`     | `string`             | Administration server写access log的文件路径。如果不想写access log，设置为`/dev/null`。仅当设置了address时，才需要该配置。               |
| `profile_path`        | `string`             | Administration server写CPU profile文件的路径。                                                                                  |
| `address`             | `core.Address`       | Administration server监听的TCP地址。如果不指定，Envoy不会启动administration server。                                               |


# 资源链接
* [http://www.servicemesher.com/envoy/](http://www.servicemesher.com/envoy/)
* [https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto#envoy-api-msg-cluster](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto#envoy-api-msg-cluster)
* [https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/service_discovery#arch-overview-service-discovery-types](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/service_discovery#arch-overview-service-discovery-types)
* [https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview#config-overview-v2-bootstrap](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview#config-overview-v2-bootstrap)
* [https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#envoy-api-msg-config-bootstrap-v2-clustermanager](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#envoy-api-msg-config-bootstrap-v2-clustermanager)
