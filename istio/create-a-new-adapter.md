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


进入MIXER_REPO目录，执行命令:
```go
go build ./...
```
确保没有报错。

# 步骤2： 编写adapter的配置

Istio中每个adapter都会有自己的配置，比如，后端server的地址等。本文章中我们要创建的mysampleadapter比较简单，会将从Mixer接收到的数据写入到一个文件中，所以，我们需要配置一个文件地址，这样adapter才知道将数据写入到哪个文件。在创建mysampleadapter的时候，需要把文件地址传给mysampleadapter。

创建一个config目录`$MIXER_REPO/adapter/mysampleadapter/config`。在config目录新建一个文件:`$MIXER_REPO/adapter/mysampleadapter/config/config.proto`，内容如下:

```proto
syntax = "proto3";

package adapter.mysampleadapter.config;

import "google/protobuf/duration.proto";
import "gogoproto/gogo.proto";

option go_package="config";

message Params {
    // Path of the file to save the information about runtime requests.
    string file_path = 1;
}
```

我们创建了一个标准的protobuff文件，来表示mysampleadapter需要的配置信息，有一个file_path字段。表示mysampleadapter将数据写入哪个文件。    
接下来，我们需要根据这个config.proto文件来生成一些go文件。

> Istio中经常会先定义一个proto文件，然后再根据proto文件自动生成一些go文件。大家要习惯。。。

为了能够自动生成go文件，我们还需要在mysampleadapter.go文件中添加一些go generate的注释。如下所示:

```go
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f mixer/adapter/mysampleadapter/config/config.proto
package mysampleadapter

import (
  "context"

  "github.com/gogo/protobuf/types"
  "istio.io/istio/mixer/pkg/adapter"
  "istio.io/istio/mixer/template/metric"
)
..
..
```

添加完了注释后，就可以来生成go文件了。进入$ISTIO目录，执行命令:
```bash
go generate ./...
go build ./...
```
执行完命令后，你能在`$MIXER_REPO/adapter/mysampadapter/config`目录看到自动生成的文件:
* adapter.mysampleradapter.config.pb.html：一个html问价，没什么用
* config.pb.go：生成的go代码，包含一个Params结构体
* config.proto_descriptor：一个纯文本文件，貌似也没什么用





