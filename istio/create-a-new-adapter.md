---
layout: post
title: Create a New Adapter in Istio
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---



# 背景
在看Istio源码时，一开始的时候有点一头雾水。主要是有以下几个原因:
1. 看Istio源码时，遇到很多自动生成的代码,比如:
    - `mixer/template/metric/template_handler.gen.go`
    - `mixer/template/metric/template_handler_service.pb.go`
    - `mixer/adapter/prometheus/config/config.pb.go`
2. Istio中的很多概念看文档还是比较模糊，比如：
    - Adapter
    - Handler
    - Template
    - Instance
    - Type
    - HandlerBuilder
   之前看Java框架的源码，可以调试进去一步一步看调用流程，非常清晰。但Istio需要k8s环境，这个方法行不通了。

参考了[https://github.com/istio/istio/wiki/Mixer-Adapter-Walkthrough](https://github.com/istio/istio/wiki/Mixer-Adapter-Walkthrough)，一步一步的实现了一个自己的adapter，并在本地启动Mixer Server和Mixer Client来进行测试。

本文章会介绍如何在Istio中新建一个Adapter，名字叫mysampleadapter。该Adapter:
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
* ISTIO: `export $ISTIO=$GOPATH/src/istio.io/istio`
* PATH: `export PATH=$GOPATH/out/darwin_amd64/release:$PATH`。mixs和mixc编译的二进制文件在`$GOPATH/out/darwin_amd64`目录中。

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


# 步骤3：将adapter的配置和adapter的代码进行关联

这一步中我们要修改mysampleadapter.go的代码，来使用mysampleadapter相关的配置(mysampleadapter/config/config.proto文件定义的)。同时更新GetInfo方法来允许Istio管理员传递mysampleadapter相关的配置(即file_path)给mysampleadapter。

添加如下的代码:
```go
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f mixer/adapter/mysampleadapter/config/config.proto
package mysampleadapter

import (
    // "github.com/gogo/protobuf/types"
	"context"
	"os"
	"path/filepath"

	"istio.io/istio/mixer/adapter/mysampleadapter/config"
	"istio.io/istio/mixer/pkg/adapter"
    "istio.io/istio/mixer/template/metric"
)

type (
	builder struct {
		adpCfg *config.Params
	}
	handler struct {
		f *os.File
	}
)

// ensure types implement the requisite interfaces
var _ metric.HandlerBuilder = &builder{}
var _ metric.Handler = &handler{}

///////////////// Configuration-time Methods ///////////////

// adapter.HandlerBuilder#Build
func (b *builder) Build(ctx context.Context, env adapter.Env) (adapter.Handler, error) {
	file, err := os.Create(b.adpCfg.FilePath)
	return &handler{f: file}, err

}

// adapter.HandlerBuilder#SetAdapterConfig
func (b *builder) SetAdapterConfig(cfg adapter.Config) {
	b.adpCfg = cfg.(*config.Params)
}

// adapter.HandlerBuilder#Validate
func (b *builder) Validate() (ce *adapter.ConfigErrors) {
	// Check if the path is valid
	if _, err := filepath.Abs(b.adpCfg.FilePath); err != nil {
		ce = ce.Append("file_path", err)
	}
	return
}

// metric.HandlerBuilder#SetMetricTypes
func (b *builder) SetMetricTypes(types map[string]*metric.Type) {
}

////////////////// Request-time Methods //////////////////////////
// metric.Handler#HandleMetric
func (h *handler) HandleMetric(ctx context.Context, insts []*metric.Instance) error {
	return nil
}

// adapter.Handler#Close
func (h *handler) Close() error {
	return h.f.Close()
}

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
		DefaultConfig: &config.Params{},
	}
}
```

进入`$ISTIO`目录，执行以下命令，确保代码是可以编译通过的
```bash
go build ./...
```

步骤4： 在mysampleadapter中编写业务逻辑

mysampleadapter要做的事情是：读取mixer穿过来的数据，将Instance和相关的Type信息写入到配置的文件中。这就需要在配置阶段，把metric的Type信息存储起来，在请求阶段来使用。修改Adapter的代码如下:

```go
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -f mixer/adapter/mysampleadapter/config/config.proto
package mysampleadapter

import (
	"context"

	// "github.com/gogo/protobuf/types"
	"fmt"
	"os"
	"path/filepath"
	config "istio.io/istio/mixer/adapter/mysampleadapter/config"
	"istio.io/istio/mixer/pkg/adapter"
	"istio.io/istio/mixer/template/metric"
)

type (
	builder struct {
		adpCfg      *config.Params
		metricTypes map[string]*metric.Type
	}
	handler struct {
		f           *os.File
		metricTypes map[string]*metric.Type
		env         adapter.Env
	}
)

// ensure types implement the requisite interfaces
var _ metric.HandlerBuilder = &builder{}
var _ metric.Handler = &handler{}

///////////////// Configuration-time Methods ///////////////

// adapter.HandlerBuilder#Build
func (b *builder) Build(ctx context.Context, env adapter.Env) (adapter.Handler, error) {
	var err error
	var file *os.File
	file, err = os.Create(b.adpCfg.FilePath)
	return &handler{f: file, metricTypes: b.metricTypes, env: env}, err

}

// adapter.HandlerBuilder#SetAdapterConfig
func (b *builder) SetAdapterConfig(cfg adapter.Config) {
	b.adpCfg = cfg.(*config.Params)
}

// adapter.HandlerBuilder#Validate
func (b *builder) Validate() (ce *adapter.ConfigErrors) {
	// Check if the path is valid
	if _, err := filepath.Abs(b.adpCfg.FilePath); err != nil {
		ce = ce.Append("file_path", err)
	}
	return
}

// metric.HandlerBuilder#SetMetricTypes
func (b *builder) SetMetricTypes(types map[string]*metric.Type) {
	b.metricTypes = types
}

////////////////// Request-time Methods //////////////////////////
// metric.Handler#HandleMetric
func (h *handler) HandleMetric(ctx context.Context, insts []*metric.Instance) error {
	for _, inst := range insts {
		if _, ok := h.metricTypes[inst.Name]; !ok {
			h.env.Logger().Errorf("Cannot find Type for instance %s", inst.Name)
			continue
		}
		h.f.WriteString(fmt.Sprintf(`HandleMetric invoke for :
		Instance Name  :'%s'
		Instance Value : %v,
		Type           : %v`, inst.Name, *inst, *h.metricTypes[inst.Name]))
	}
	return nil
}

// adapter.Handler#Close
func (h *handler) Close() error {
	return h.f.Close()
}

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
		DefaultConfig: &config.Params{},
	}
}
```

进入$ISTIO目录，执行以下命令:
```bash
go build ./...
```

步骤5： 将mysampleadapter添加到Mixer的adapter体系中

我们已经创建好了一个adapter，但是Mixer还不知道这个mysampleadapter。我们需要把它添加到Mixer中去。Mixer支持的所有adapter在一个yaml文件中定义了:
`$MIXER_REPO/adapter/inventory.yaml`。把mysampleadapter添加到该文件中:
```yaml
prometheus: "istio.io/istio/mixer/adapter/prometheus"
servicecontrol: "istio.io/istio/mixer/adapter/servicecontrol"
stackdriver: "istio.io/istio/mixer/adapter/stackdriver"
statsd: "istio.io/istio/mixer/adapter/statsd"
stdio: "istio.io/istio/mixer/adapter/stdio"
mysampleadapter: "istio.io/istio/mixer/adapter/mysampleadapter"
...
```

然后进入`$MIXER_REPO/adapter`目录，执行以下命令:
```bash
go generate $MIXER_REPO/adapter/doc.go
```
现在，mysampleadapter被添加到Mixer中了，可以开始接收数据了。

步骤6： 编写一个管理配置的sample

为了测试mysampleadapter是否工作正常，我们需要一份管理配置，即operator cofiguration。
创建一个文件:`$MIXER_REPO/adapter/mysampleadapter/testdata/mysampleadapter.yaml`，内容如下:
```yaml
# instance configuration for template 'metric'
apiVersion: "config.istio.io/v1alpha2"
kind: metric
metadata:
 name: requestcount
 namespace: istio-system
spec:
 value: "1"
 dimensions:
   target: destination.service | "unknown"

---
# handler configuration for adapter 'metric'
apiVersion: "config.istio.io/v1alpha2"
kind: mysampleadapter
metadata:
 name: hndlrTest
 namespace: istio-system
spec:
 file_path: "out.txt"
---
# rule to dispatch to your handler
apiVersion: "config.istio.io/v1alpha2"
kind: rule
metadata:
 name: mysamplerule
 namespace: istio-system
spec:
 match: "true"
 actions:
 - handler: hndlrTest.mysampleadapter
   instances:
   - requestcount.metric
```

另外，把attributes.yaml文件也复制到testdata目录:
```bash
cp $MIXER_REPO/testdata/config/attributes.yaml $MIXER_REPO/adapter/mysampleadapter/testdata
```

步骤7： 启动Mixer Server
执行以下命令来启动Mixer Server:
```bash
cd $ISTIO
make mixs
// locate mixs binary, should be $GOPATH/out/linux_amd64/release/mixs on linux os and 
// $GOPATH/out/darwin_amd64/release/mixs on mac os. 
// Choose command below according to your os:
$GOPATH/out/linux_amd64/release/mixs server --configStoreURL=fs://$(pwd)/mixer/adapter/mysampleadapter/testdata
```

终端会输出类似如下的内容:
```text
..
..
Mixer started with
MaxMessageSize: 1048576
MaxConcurrentStreams: 1024
APIWorkerPoolSize: 1024
AdapterWorkerPoolSize: 1024
ExpressionEvalCacheSize: 1024
APIPort: 9091
MonitoringPort: 9093
SingleThreaded: false
ConfigStore2URL: fs:///usr/local/google/home/guptasu/go/src/istio.io/istio/mixer/adapter/mysampleadapter/sampleoperatorconfig
ConfigDefaultNamespace: istio-system
ConfigIdentityAttribute: destination.service
ConfigIdentityAttributeDomain: svc.cluster.local
LoggingOptions: log.Options{OutputPaths:[]string{"stdout"}, ErrorOutputPaths:[]string{"stderr"}, RotateOutputPath:"", RotationMaxSize:104857600, RotationMaxAge:30, RotationMaxBackups:1000, JSONEncoding:false, IncludeCallerSourceLocation:false, stackTraceLevel:"none", outputLevel:"info"}
TracingOptions: tracing.Options{ZipkinURL:"", JaegerURL:"", LogTraceSpans:false}

2018-01-06T01:43:12.305995Z	info	template Kind: kubernetesenv, &InstanceParam{SourceUid:,SourceIp:,DestinationUid:,DestinationIp:,OriginUid:,OriginIp:,AttributeBindings:map[string]string{},}
...
```

步骤9： 使用Mixer Client发送数据给Mixer Server

接下来我们在Mixer Client中向Mixer Server发起一个REPORT调用。这将会导致Mixer调用mysampleadapter。

执行以下命令来启动Mixer Client
```bash
cd $ISTIO
make mixc
mixc report -s destination.service="svc.cluster.local" -t request.time="2017-02-01T10:12:14Z"
```
命令执行完后，可以在`$ISTIO`目录看到一个out.txt文件。内容大概如下:
```text
HandleMetric invoke for
       Instance Name  : requestcount.metric.istio-system
       Instance Value : {requestcount.metric.istio-system 1 map[response_code:200 service:unknown source:unknown target:unknown version:unknown method:unknown] UNSPECIFIED map[]}
       Type           : {INT64 map[response_code:INT64 service:STRING source:STRING target:STRING version:STRING method:STRING] map[]}
```

可以通过Mixer Client发送其他的数据到Mixer Server，比如:
```bash
cd $ISTIO
make mixc
mixc report -s="destination.service=svc.cluster.local,destination.service=mySrvc" -i="response.code=400" --stringmap_attributes="destination.labels=app:dummyapp"
```

