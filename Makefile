#!/usr/bin/env bash
## 下载源码
# git clone https://github.com/golang/tools.git
# 安装
# go install golang.org/x/tools/cmd/goimports
fmt:
	goimports -l -w ./