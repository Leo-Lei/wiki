---
title: Prometheus Kubernetes
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---

# 




Kubernetes SD

其次，上次提到，我们可以用到 Service Discovery 这个功能，其中就包含 Kubernetes SD。

它包含四种角色：

node
service
pod
endpoints
由于篇幅所限，这里只是简单介绍下其中的 node 还有 pod 角色：
```yaml
- job_name: 'kubernetes-nodes'
  scheme: https
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
  
  kubernetes_sd_configs:
  - role: node
  
  relabel_configs:
    # 即从 __meta_kubernetes_node_label_<labelname> 这个配置中取出 labelname 以及 value
  - action: labelmap
    regex: __meta_kubernetes_node_label_(.+)
    
    # 配置 address 为 k8s api 的地址，相关的 ca 证书以及 token 在上面配置
  - target_label: __address__
    replacement: kubernetes.default.svc:443
    
    # 取出所有的 node，然后设置 /api/v1/nodes/<node_name>/proxy/metrics 为 metrics path
  - source_labels: 
    - __meta_kubernetes_node_name
    regex: (.+)
    target_label: __metrics_path__
    replacement: /api/v1/nodes/${1}/proxy/metrics
```
接下来的这个 pod 角色挺重要：

```yaml
- job_name: 'kubernetes-pods'
  kubernetes_sd_configs:
  - role: pod
  relabel_configs:
  - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
    action: keep
    regex: true
  - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
    action: replace
    target_label: __metrics_path__
    regex: (.+)
  - source_labels: [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
    action: replace
    regex: ([^:]+)(?::\d+)?;(\d+)
    replacement: $1:$2
    target_label: __address__
  - action: labelmap
    regex: __meta_kubernetes_pod_label_(.+)
  - source_labels: [__meta_kubernetes_namespace]
    action: replace
    target_label: kubernetes_namespace
  - source_labels: [__meta_kubernetes_pod_name]
    action: replace
    target_label: kubernetes_pod_name
```

在定义了这个角色之后，你只要在你部署的应用 Pod 描述中，加入以下 annotations 就能让 Prometheus 自动发现此 Pod 并采集监控数据了：

```yaml
annotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "<your app port>"
```
其它详细配置请看 这里。

