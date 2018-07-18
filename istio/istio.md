---
layout: post
title: Istio
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---


|                              Command                                     |                   Desc                    |
| ------------------------------------------------------------------------ | ----------------------------------------- |
| `istioctl get routerules`                                                |                                           | 
| `istioctl delete routerule recommendation-default -n tutorial`           |                                           |
| `istioctl get routerule recommendation-default -o yaml -n tutorial`      |                                           |


```text
.
├── addons                                   
├── bin
├── broker
├── docker
├── galley
├── install
├── istioctl
├── mixer
├── pilot
├── pkg
├── prow
├── release
├── samples
├── security
├── tests
├── tools
├── vendor
├── codecov.requirement
├── codecov.skip
├── CONTRIBUTING.md
├── downloadIstio.sh
├── Gopkg.lock
├── Gopkg.toml
├── istio.deps
├── istio.VERSION
├── istio.yaml
├── LICENSE
├── lintconfig_base.json
├── Makefile
├── OWNERS
└── README.md
```

包，文件夹和文件的功能如下:

| Package/Directory/File          |                                     Desc                                           |
| ------------------------------- | ---------------------------------------------------------------------------------- |
| addons                          | 一些插件，比如展示metrics的grafana和绘制服务调用图的servicegraph                       |
| bin                            	| 存放初始化依赖、编译、插件证书检查、代码生成的脚本                                      |
| broker                          |	Istio对Open Service Broker的一种实现，该API使得外部服务能自动访问Istio服务。broker目前还处于研发阶段。 |
| galley                          |	提供了Istio的配置管理功能，目前还处于研发阶段。                                         |
| install                         |	生成各环境（ansible、consul、ereka、kubernetes等）安装istio时需要yaml配置清单。         |
| istioctl	                      | istio终端控制工具（类似kubectl之于kubernetes），用户通过istioctl来修改istio运行时配置，执行服务治理策略。 |
| mixer                           |	“混音器”，参与到tracffic处理流程。通过对envoy上报的attributes进行处理，结合内部的adapters实现日志记录、监控指标采集展示、配额管理、ACL检查等功能。  |
| pilot                           |	“领航员”，pliot对Envoy的生命周期进行管理，同时提供了智能路由（如A/B测试、金丝雀部署）、流量管理（超时、重试、熔断）功能。    |
| pkg                             |	顶级公共包，包含istio版本处理、tracing、日志记录、缓存管理等。                            |
| release                         |	包含Istio在各平台上进行编译的脚本。                                                     |
| samples                         |	Istio提供的微服务样例，比如bookinfo。                                                   |
| security                        |     	Istio用户身份验证、服务间认证。                 |
| tests	                          | 测试用例、脚本等。    |
| vendor                          |	dep生成的第三方依赖。     |
| Gopkg.*	                        | dep需要version config和version lock文件。     |
| Makefile                        |	Istio Makefile，编译docker镜像时会引用tools/istio-docker.mk这个Makefile。    |

