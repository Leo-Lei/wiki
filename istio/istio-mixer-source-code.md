---
layout: post
title: Istio Mixer Source Code
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---


# 编译mixer二进制文件和Docker镜像
Makefile:
`istio/Makefile`:
```bash
MIXER_GO_BINS:=${ISTIO_OUT}/mixs ${ISTIO_OUT}/mixc
mixc:
	bin/gobuild.sh ${ISTIO_OUT}/mixc ./mixer/cmd/mixc
mixs:
	bin/gobuild.sh ${ISTIO_OUT}/mixs ./mixer/cmd/mixs
...
include tools/istio-docker.mk
```
其中:
* mixc: Mixer客户端，通过mixc可以和运行的Mixer进行交互
* mixs: Mixer服务端。和Envoy，adapter交互。部署Istio时启动。





`istio/mixer/cmd/mixs/main.go`
```go
func supportedTemplates() map[string]template.Info {
	return generatedTmplRepo.SupportedTmplInfo
}

func supportedAdapters() []adptr.InfoFn {
	return adapter.Inventory()
}

func main() {
	rootCmd := cmd.GetRootCmd(os.Args[1:], supportedTemplates(), supportedAdapters(), shared.Printf, shared.Fatalf)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
```
