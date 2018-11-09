---
layout: post
title: Envoy Bootstrap配置
date: 2017-06-16 14:35:00
tags:
- docker
categories: Java
---



Envoy v2 API在[data plane API repository](https://github.com/envoyproxy/data-plane-api/tree/master/envoy/api)中，以proto3的形式被定义。



# Bootstrap configuration

如果要使用envoy的V2 API，必须提供一个bootstrap配置文件。通常是一个yaml文件。yaml文件中提供静态的配置和动态配置。启动envoy时，通过`-c`参数指定bootstrap配置文件。

```bash
./envoy -c path/to/config.yaml
```

# Example

下面是一个YAMl的配置，将HTTP请求从127.0.0.1:10000路由到127.0.0.2:1234。

```yaml
admin:
  access_log_path: /tmp/admin_access.log
  address:
    socket_address: { address: 127.0.0.1, port_value: 9901 }

static_resources:
  listeners:
  - name: listener_0
    address:
      socket_address: { address: 127.0.0.1, port_value: 10000 }
    filter_chains:
    - filters:
      - name: envoy.http_connection_manager
        config:
          stat_prefix: ingress_http
          codec_type: AUTO
          route_config:
            name: local_route
            virtual_hosts:
            - name: local_service
              domains: ["*"]
              routes:
              - match: { prefix: "/" }
                route: { cluster: some_service }
          http_filters:
          - name: envoy.router
  clusters:
  - name: some_service
    connect_timeout: 0.25s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: some_service
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 1234
```


# 资源链接
* [https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview#config-overview-v2-bootstrap](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/v2_overview#config-overview-v2-bootstrap)