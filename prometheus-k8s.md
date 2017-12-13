---
title: Prometheus Kubernetes
date: 2017-06-09 10:22:23
categories:
- Monitoring
tags:
- Prometheus
---

# Prometheus和k8s集成

Prometheus以容器的形式运行在k8s环境中，在Deployment中为Pod添加Prometheus相关的annotation，这样，k8s就可以自动发现这些Pod暴露的metrics，去拉取数据。

# 部署Prometheus
在k8s中部署Prometheus就是要创建一些k8s资源。相关的文件在以下地址可以找到:        
https://github.com/Leo-Lei/file-repo/tree/master/k8s-prometheus        

这样Prometheus就已经运行在k8s环境中了。


# 在Pod中添加Prometheus的annotation
```yaml
annotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "<your app port>"
```

其它详细配置请看 [这里](https://github.com/prometheus/prometheus/blob/master/documentation/examples/prometheus-kubernetes.yml)。


# Resources
* [https://coreos.com/blog/monitoring-kubernetes-with-prometheus.html](https://coreos.com/blog/monitoring-kubernetes-with-prometheus.html)

* [https://xizhibei.github.io/2017/08/19/deploy-prometheus-in-k8s/](https://xizhibei.github.io/2017/08/19/deploy-prometheus-in-k8s/)

