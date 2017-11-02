---
title: Prometheus Metrics
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---

# Counter

```text
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 14928
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 6.447458e+06
```

# gauge

```text
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 6944
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 244112
```

# histogram

```text
# HELP prometheus_local_storage_series_chunks_persisted The number of chunks persisted per series.
# TYPE prometheus_local_storage_series_chunks_persisted histogram
prometheus_local_storage_series_chunks_persisted_bucket{le="1"} 183
prometheus_local_storage_series_chunks_persisted_bucket{le="2"} 285
prometheus_local_storage_series_chunks_persisted_bucket{le="4"} 444
prometheus_local_storage_series_chunks_persisted_bucket{le="8"} 444
prometheus_local_storage_series_chunks_persisted_bucket{le="16"} 444
prometheus_local_storage_series_chunks_persisted_bucket{le="32"} 444
prometheus_local_storage_series_chunks_persisted_bucket{le="64"} 444
prometheus_local_storage_series_chunks_persisted_bucket{le="128"} 444
prometheus_local_storage_series_chunks_persisted_bucket{le="+Inf"} 444
prometheus_local_storage_series_chunks_persisted_sum 973
prometheus_local_storage_series_chunks_persisted_count 444
```

# summary

```text
# HELP prometheus_sd_azure_refresh_duration_seconds The duration of a Azure-SD refresh in seconds.
# TYPE prometheus_sd_azure_refresh_duration_seconds summary
prometheus_sd_azure_refresh_duration_seconds{quantile="0.5"} NaN
prometheus_sd_azure_refresh_duration_seconds{quantile="0.9"} NaN
prometheus_sd_azure_refresh_duration_seconds{quantile="0.99"} NaN
prometheus_sd_azure_refresh_duration_seconds_sum 0
prometheus_sd_azure_refresh_duration_seconds_count 0
# HELP prometheus_sd_consul_rpc_duration_seconds The duration of a Consul RPC call in seconds.
# TYPE prometheus_sd_consul_rpc_duration_seconds summary
prometheus_sd_consul_rpc_duration_seconds{call="service",endpoint="catalog",quantile="0.5"} NaN
prometheus_sd_consul_rpc_duration_seconds{call="service",endpoint="catalog",quantile="0.9"} NaN
prometheus_sd_consul_rpc_duration_seconds{call="service",endpoint="catalog",quantile="0.99"} NaN
prometheus_sd_consul_rpc_duration_seconds_sum{call="service",endpoint="catalog"} 0
prometheus_sd_consul_rpc_duration_seconds_count{call="service",endpoint="catalog"} 0
prometheus_sd_consul_rpc_duration_seconds{call="services",endpoint="catalog",quantile="0.5"} NaN
prometheus_sd_consul_rpc_duration_seconds{call="services",endpoint="catalog",quantile="0.9"} NaN
prometheus_sd_consul_rpc_duration_seconds{call="services",endpoint="catalog",quantile="0.99"} NaN
prometheus_sd_consul_rpc_duration_seconds_sum{call="services",endpoint="catalog"} 0
prometheus_sd_consul_rpc_duration_seconds_count{call="services",endpoint="catalog"} 0
```

