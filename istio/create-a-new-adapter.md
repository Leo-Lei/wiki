---
layout: post
title: Create a New Adapter in Istio
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---



本文章会介绍如何在Istio中新建一个Adapter。该Adapter:
1. 支持Istio自带的metric模板
2. 对于每个请求，Adapter会把它接收到的数据打印到一个文件中


# 准备工作
### 安装go
安装go，配置`GOPATH`。
### 下载Istio源码

```bash
git clone git clone https://github.com/istio/istio
```
将istio源码复制到你的GOPATH中。比如`$GOPATH/src/istio.io/istio`。

### 安装protoc(protobuff编译器)
Istio中的template，adapter等代码都是根据proto文件来生成的，所以需要安装protoc。
> 注: 和protobuf的通常用法不一样，在Istio中，我们会先定义一个`template.proto`文件，是一个标准的protobuf格式的文件。然后Istio根据自己的脚本和protoc来生成对应的go代码。包括Template，Instance，Adapter，Handler等。而不只是用于序列化的代码。

从[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)下载protoc。并将protoc添加到PATH中。

### 设置一些环境变量
为了方便操作，和统一路径的描述，我们定义一些环境变量:
* MIXER_REPO: `export MIXER_REPO=$GOPATH/src/istio.io/istio/mixer`
* ISTIO: `export $ISTIO=$GOPATH/src/istio.io`

# 步骤1： 编写基本的adapter框架代码

创建一个mysampleadapter目录`$MIXER_REPO/adapter/mysampleadapter`。在该目录中创建一个mysampleadapter.go文件`$MIXER_REPO/adapter/mysampleadapter.mysampleadapter.go`。

mysampleadapter.go的代码如下:
```go
package mysampleadapter

import (
  "context"

  "github.com/gogo/protobuf/types"
  "istio.io/istio/mixer/pkg/adapter"
  "istio.io/istio/mixer/template/metric"
)

type (
  builder struct {
  }
  handler struct {
  }
)

// ensure types implement the requisite interfaces
var _ metric.HandlerBuilder = &builder{}
var _ metric.Handler = &handler{}

///////////////// Configuration-time Methods ///////////////

// adapter.HandlerBuilder#Build
func (b *builder) Build(ctx context.Context, env adapter.Env) (adapter.Handler, error) {
  return &handler{}, nil
}

// adapter.HandlerBuilder#SetAdapterConfig
func (b *builder) SetAdapterConfig(cfg adapter.Config) {
}

// adapter.HandlerBuilder#Validate
func (b *builder) Validate() (ce *adapter.ConfigErrors) { return nil }

// metric.HandlerBuilder#SetMetricTypes
func (b *builder) SetMetricTypes(types map[string]*metric.Type) {
}

////////////////// Request-time Methods //////////////////////////
// metric.Handler#HandleMetric
func (h *handler) HandleMetric(ctx context.Context, insts []*metric.Instance) error {
  return nil
}

// adapter.Handler#Close
func (h *handler) Close() error { return nil }

////////////////// Bootstrap //////////////////////////
// GetInfo returns the adapter.Info specific to this adapter.
func GetInfo() adapter.Info {
  return adapter.Info{
     Name:        "mysampleadapter",
     Description: "Logs the metric calls into a file",
     SupportedTemplates: []string{
        metric.TemplateName,
     },
     NewBuilder:    func() adapter.HandlerBuilder { return &builder{} },
     DefaultConfig: &types.Empty{},
  }
}
```
mysampleadapter目前不包含任何功能，只是实现了一个adapter的基本框架。有几点需要注意:
* 定义了两个struct：handler和builder。几乎所有的adapter中都会定义了handler和builder这两个struct。这两个struct会实现template的接口。
* 一个adapter支持哪个template，就要实现template定义的接口。mysampleadapter支持metric模板，就必须要实现metric模板的接口。
* builder实现了metric模板的接口: 
    - Build: `istio.io/istio/imxer/pkg/adapter/handler.go`中的HandlerBuilder#Build
    - SetAdapterConfig: `istio.io/istio/imxer/pkg/adapter/handler.go`中的HandlerBuilder#SetAdapterConfig
    - Validate: `istio.io/istio/imxer/pkg/adapter/handler.go`中的HandlerBuilder#Validate
    - SetMetricTypes: `istio.io/istio/mixer/template/metric/template_handler.gen.go`中的HandlerBuilder#SetMetricTypes
* handler实现了metric模板的接口:
    - HandleMetric: `istio.io/istio/mixer/template/metric/template_handler.gen.go`中的Handler#HandleMetric
    - Close: `istio.io/istio/imxer/pkg/adapter/handler.go`中的Handler#Close
* mysampleadapter定义了一个GetInfo方法，这个方法并没有实现任何接口，但Istio中的每一个Adapter都必须要有这样一个方法，方法名叫GetInfo，返回adapter.Info对象。即返回adapter的相关信息。












