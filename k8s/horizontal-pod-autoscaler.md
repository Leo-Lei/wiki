---
layout: post
title: Kubernetes Horizontal Pod Autoscaler(HPA)
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---


# 自动化，智能化的Pod扩容

可以通过`kubectl scale`命令来实现Pod的扩容或缩容。
使用HPA来实现根据当前系统的负载变化来自动触发水平扩容或缩容。


```yaml
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata: 
  name: php-apache
  namespace: default
spec:
  maxReplicas: 10
  minReplicas: 1
  scaleTargetRef:
    kind: Deployment
    name: php-apache
  targetCPUUtilizationPercentage: 90
```

上面的HPA控制目标对象为一个名叫php-apache的Deployment里的Pod副本，当这些Pod副本的CPUUtilizationPercentage值超过90%时会触发自动扩容行为。
