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
# HELP prometheus_target_interval_length_seconds Actual intervals between scrapes.
# TYPE prometheus_target_interval_length_seconds summary
prometheus_target_interval_length_seconds{interval="15s",quantile="0.01"} 14.997978115
prometheus_target_interval_length_seconds{interval="15s",quantile="0.05"} 14.997978115
prometheus_target_interval_length_seconds{interval="15s",quantile="0.5"} 14.999992755
prometheus_target_interval_length_seconds{interval="15s",quantile="0.9"} 15.000945051
prometheus_target_interval_length_seconds{interval="15s",quantile="0.99"} 15.002223278
prometheus_target_interval_length_seconds_sum{interval="15s"} 7950.0376901510035
prometheus_target_interval_length_seconds_count{interval="15s"} 530
```
