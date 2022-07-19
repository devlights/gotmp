# gotmp

個人用のGo言語メモプログラム置き場

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/devlights/gotmp)

[gotip](https://pkg.go.dev/golang.org/dl/gotip) is installed, so you can try the latest version of Go at any time.

```sh
$ task version
task: [version] gotip version
go version devel go1.19-de8101d Mon Jul 18 18:04:23 2022 +0000 linux/amd64

$ task init
task: [init] gotip mod init app
go: creating new go.mod: module app
task: [init] touch main.go
task: [init] gp open main.go

$ task run
task: [run] gotip run main.go

$ task clean
task: [clean] rm go.mod *.go
```
