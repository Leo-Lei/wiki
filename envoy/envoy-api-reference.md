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


### config.bootstrap.v2.ClusterManager

```json
{
  "local_cluster_name": "...",
  "outlier_detection": "{...}",
  "upstream_bind_config": "{...}",
  "load_stats_config": "{...}"
}
```

|             字段             |         类型                                           |                                                             说明                            |
| --------------------------- | ----------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `local_cluster_name`        | `string`                                              | Envoy所属的cluster。如果想实现`zone aware routing`,必须要设置该值。                               |
| `outlier_detection`         | `config.bootstrap.v2.ClusterManager.OutlierDetection` | outlier检测。                                                                                |
| `upstream_bind_config`      | `core.BindConfig`                                     | Optional configuration used to bind newly established upstream connections                  |

### config.bootstrap.v2.ClusterManager.OutlierDetection

```json
{
  "event_log_path": "..."
}
```

|             字段             |         类型          |              说明                 |
| --------------------------- | -------------------- | --------------------------------- |
| `event_log_path`            | `string`             | outlier event log path            |




# Listener

```json
{
  "name": "...",
  "address": "{...}",
  "filter_chains": [],
  "use_original_dst": "{...}",
  "per_connection_buffer_limit_bytes": "{...}",
  "metadata": "{...}",
  "drain_type": "...",
  "listener_filters": [],
  "transparent": "{...}",
  "freebind": "{...}",
  "socket_options": [],
  "tcp_fast_open_queue_length": "{...}"
}
```


|                  字段                 |              类型              |                            说明                              |
| ------------------------------------ | ------------------------------ | ------------------------------------------------------------ |
| `name`                               | `string`                       | Listener的名字。如果不指定，Envoy分配一个UUID。                   |
| `address`                            | `core.Address`                 | Listener监听的地址。通常，Listener监听的地址是唯一的。             |
| `filter_chains`                      | `listener.FilterChain`         | Listener包含的filter链                                        |
| `use_original_dst`                   | `BoolValue`                    | 这个字段被废弃了，请使用`original_dst listener filter`           |
| `per_connection_buffer_limit_bytes`  | `UInt32Value`                  | 读和写的buffers大小，默认1M                                     |
| `metadata`                           | `core.Metadata`                | Listener metadata                                            |
| `drain_type`                         | `Listener.DrainType`           |                                                              |
| `listener_filters`                   | `listener.ListenerFilter`      |                                                              |
| `transparent`                        | `BoolValue`                    |                                                              |
| `freebind`                           | `BoolValue`                    |                                                              |
| `socket_options`                     | `core.SocketOption`            |                                                              |
| `tcp_fast_open_queue_length`         | `UInt32Value`                  |                                                              |


### listener.FilterChain

```json
{
  "filter_chain_match": "{...}",
  "tls_context": "{...}",
  "filters": [],
  "use_proxy_proto": "{...}",
  "transport_socket": "{...}"
}
```

|                  字段                 |              类型              |                            说明                              |
| ------------------------------------ | ------------------------------ | ------------------------------------------------------------ |
| `filter_chain_match`                 | `listener.FilterChainMatch`    | 指定connection匹配filter chain的条件                           |
| `tls_context`                        | `auth.DownstreamTlsContext`    | TLS上下文                                                     |
| `filters`                            | `listener.Filter`              | listener chain包含的filter列表。                               |
| `use_proxy_proto`                    | `BoolValue`                    |                                                              |
| `transport_socket`                   | `core.TransportSocket`         |                                                              |


### listener.Filter

```json
{
  "name": "...",
  "config": "{...}"
}
```

|          字段       |              类型              |                                                     说明                                               |
| ------------------ | ------------------------------ | ----------------------------------------------------------------------------------------------------- |
| `name`             | `string`                       | filter类型的名字。必须是内置filter中的某一种。比如`envoy.tcp_proxy`或`envoy.http_connection_manager`         |
| `config`           | `Struct`                       | filter的配置。取决于filter的类型。                                                                        |

1. `name`字段必须是一个Envoy支持的filter。内置的filter有：
    * envoy.client_ssl_auth
    * envoy.echo
    * envoy.http_connection_manager
    * envoy.mongo_proxy
    * envoy.ratelimit
    * envoy.redis_proxy
    * envoy.tcp_proxy

# HTTP connection manager

### config.filter.network.http_connection_manager.v2.HttpConnectionManager

```json
{
  "codec_type": "...",
  "stat_prefix": "...",
  "rds": "{...}",
  "route_config": "{...}",
  "http_filters": [],
  "add_user_agent": "{...}",
  "tracing": "{...}",
  "http_protocol_options": "{...}",
  "http2_protocol_options": "{...}",
  "server_name": "...",
  "idle_timeout": "{...}",
  "stream_idle_timeout": "{...}",
  "request_timeout": "{...}",
  "drain_timeout": "{...}",
  "delayed_close_timeout": "{...}",
  "access_log": [],
  "use_remote_address": "{...}",
  "xff_num_trusted_hops": "...",
  "internal_address_config": "{...}",
  "skip_xff_append": "...",
  "via": "...",
  "generate_request_id": "{...}",
  "forward_client_cert_details": "...",
  "set_current_client_cert_details": "{...}",
  "proxy_100_continue": "...",
  "represent_ipv4_remote_address_as_ipv4_mapped_ipv6": "...",
  "upgrade_configs": [],
  "bugfix_reverse_encode_order": "{...}"
}
```


|                 字段                 |                                                 类型                                                  |                  说明                |
| ----------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------ |
| `codec_type`                        | `config.filter.network.http_connection_manager.v2.HttpConnectionManager.CodecType`                    | 使用的编码，解码类型                    |
| `stat_prefix`                       | `string`                                                                                              | statistics前缀                       |
| `rds`                               | `config.filter.network.http_connection_manager.v2.Rds`                                                | route table通过RDS API动态加载        |
| `route_config`                      | `RouteConfiguration`                                                                                  | 静态配置route table                   |
| `http_filters`                      | `config.filter.network.http_connection_manager.v2.HttpFilter`                                         | Http filter列表                      |
| `add_user_agent`                    | `BoolValue`                                                              | 是否添加`user-agent`和` x-envoy-downstream-service-cluster`header。 |
| `tracing`                           | `config.filter.network.http_connection_manager.v2.HttpConnectionManager.Tracing`                      | tracing配置                         |   
| `http_protocol_options`             | `core.Http1ProtocolOptions`                                                                           | 额外的HTTP1配置，会传递给HTTP1 codec   | 
| `http2_protocol_options`            | `core.Http2ProtocolOptions`                                                                           | 额外的HTTP2配置，会传递给HTTP1 codec   |
| `server_name`                       | `string`                                                                                              | 默认值是envoy                        |
| `idle_timeout`                      | `Duration`                                                                                            |                                     |
| `stream_idle_timeout`               | `Duration`                                                                                            |                                     |
| `request_timeout`                   | `Duration`                                                                                            |                                     |
| `drain_timeout`                     | `Duration`                                                                                            |                                     |
| `delayed_close_timeout`             | `Duration`                                                                                            |                                     |
| `access_log`                        | `config.filter.accesslog.v2.AccessLog`                                                                | HTTP access logs配置                 |
| `use_remote_address`                | `BoolValue`                                                                                           |                                     |
| `xff_num_trusted_hops`              | `uint32`                                                                                              |                                     |
| `internal_address_config`           | `config.filter.network.http_connection_manager.v2.HttpConnectionManager.InternalAddressConfig`        |                                     |
| `skip_xff_append`                   | `bool`                                                                                                |                                     |
| `via`                               | `string`                                                                                              |                                     |
| `generate_request_id`               | `BoolValue`                                                                                           | 是否生成x-request-id                 |
| `fowward_client_cert_details`       | `config.filter.network.http_connection_manager.v2.HttpConnectionManager.ForwardClientCertDetails`     |                                     |
| `set_current_client_cert_details`   | `config.filter.network.http_connection_manager.v2.HttpConnectionManager.SetCurrentClientCertDetails`  |                                     |
| `proxy_100_continue`                | `bool`                                                                                                |                                     |
| `represent_ipv4_remote_address_as_ipv4_mapped_ipv6` | `bool`                                                                                |                                     |
| `upgrade_configs`                   | `config.filter.network.http_connection_manager.v2.HttpConnectionManager.UpgradeConfig`                |                                     |
| `bugfix_reverse_encode_order`       | `BoolValue`                                                                                           |                                     |

1. `rds`和`route_config`必须有一个被指定






### config.filter.network.http_connection_manager.v2.HttpFilter

```json
{
  "name": "...",
  "config": "{...}"
}
```


|      字段      |       类型     |                         说明                    |
| ------------- | -------------- | ---------------------------------------------- |
| `name`        | `string`       | filter类型的名字。必须是envoy支持的一种filter。     |
| `config`      | `struct`       | filter的配置。取决于具体的filter类型。             |


1. filter类型的名字必须是envoy支持的一种。内置的filter有:
    * envoy.buffer
    * envoy.cors
    * envoy.fault
    * envoy.gzip
    * envoy.http_dynamo_filter
    * envoy.grpc_http1_bridge
    * envoy.grpc_json_transcoder
    * envoy.grpc_web
    * envoy.health_check
    * envoy.header_to_metadata
    * envoy.ip_tagging
    * envoy.lua
    * envoy.rate_limit
    * envoy.router
    * envoy.spuash




###  Router

[https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_filters/router_filter#config-http-filters-router)

### config.filter.http.router.v2.Router

```json
{
  "dynamic_stats": "{...}",
  "start_child_span": "...",
  "upstream_log": [],
  "suppress_envoy_headers": "..."
}
```

|             字段           |                      类型              |                         说明                    |
| ------------------------- | -------------------------------------- | ---------------------------------------------- |
| `dynamic_stats`           | `BoolValue`                            | 是否生成动态statistics数据。                      |
| `start_child_span`        | `bool`                                 |                                                |
| `upstream_log`            | `config.filter.accesslog.v2.AccessLog` |                                                |
| `suppress_envoy_headers`  | `bool`                                 |                                                |



# HTTP route configuration

```json
{
  "name": "...",
  "virtual_hosts": [],
  "internal_only_headers": [],
  "response_headers_to_add": [],
  "response_headers_to_remove": [],
  "request_headers_to_add": [],
  "request_headers_to_remove": [],
  "validate_clusters": "{...}"
}
```


|               字段             |             类型              |                          说明                      |
| ----------------------------- | ----------------------------- | ------------------------------------------------- |
| `name`                        | `string`                      | route的名字。                                      |
| `virtual_hosts`               | `route.VirtualHost`           | virtual host列表。                                 |
| `internal_only_headers`       | `string`                      | 指定内部的header。                                  |
| `response_headers_to_add`     | `core.HeaderValueOption`      | 指定header列表，被添加到response header中。           |
| `response_headers_to_remove`  | `string`                      | 指定HTTP header列表，将被从response header中移除。    |
| `request_headers_to_add`      | `core.HeaderValueOption`      | 指定HTTP header列表，将被添加到request中              |
| `request_headers_to_remove`   | `string`                      | 指定HTTP header列表，将从request中被移除。            |
| `validate_clusters`           | `BoolValue`                   | 是否验证cluster                                    |

# 资源链接
* [http://www.servicemesher.com/envoy/](http://www.servicemesher.com/envoy/)
* [https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto#envoy-api-msg-cluster](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto#envoy-api-msg-cluster)
* [https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/service_discovery#arch-overview-service-discovery-types](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/service_discovery#arch-overview-service-discovery-types)
* [https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview#config-overview-v2-bootstrap](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview#config-overview-v2-bootstrap)
* [https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#envoy-api-msg-config-bootstrap-v2-clustermanager](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#envoy-api-msg-config-bootstrap-v2-clustermanager)
