---
title: Prometheus Service Discovery for Kubernetes
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---



# Alertmanager Discovery for Kubernetes
知道发现:
* 有label:`name:alertmanager`
* namespace: default
* 端口不为空
```yml
alerting:
  alertmanagers:
  - kubernetes_sd_configs:
      - role: pod
    tls_config:
      ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
    relabel_configs:
    - source_labels: [__meta_kubernetes_pod_label_name]
      regex: alertmanager
      action: keep
    - source_labels: [__meta_kubernetes_namespace]
      regex: default
      action: keep
    - source_labels: [__meta_kubernetes_pod_container_port_number]
      regex:
      action: drop
```
