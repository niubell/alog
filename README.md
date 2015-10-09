alog
===========
# Overview 
alog是个人基于[Go](http://golang.org/)开发的异步日志组件，实现部分参照了[aiwuTech团队](https://github.com/aiwuTech)的repo，目前功能比较简单，后期会不断完善分支，欢迎大家贡献代码。

# Features 
----------
* 日志文件可以按照日期进行切割
* 基于goroutine、channel实现异步文件检查和日志写入
* 目前只支持Linux/MacOs

# Installation

    go get github.com/bigpyer/alog
# Test
    cd src/example
    go run example.go
    tail -f example.log 观察日志打印情况
