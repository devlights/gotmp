# gotmp

個人用のGo言語メモプログラム置き場

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/devlights/gotmp)

```sh
$ task init
task: [init] go mod init app
go: creating new go.mod: module app
task: [init] go get -u github.com/devlights/gomy@latest
go: added github.com/devlights/gomy v0.6.0
task: [init] echo -e "package main\nfunc main(){\n}" > main.go
task: [fmt] goimports -w .
task: [build] go mod tidy
task: [build] go build -o app main.go

$ task run
task: Task "build" is up to date
task: [run] ./app

$ task clean
task: [OS:clean] go clean
task: [OS:clean] rm go.* *.go
task: [OS:clean] rm -rf ./.task
```
