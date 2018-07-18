---
layout: post
title: Istio Mixer Source Code
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---


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
