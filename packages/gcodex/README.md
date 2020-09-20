# qylibrary.gcodex包


## 说明

1. Golang 版本要求最低 1.13
2. 使用该库需要设置如下环境变量
   ```
   GOPROXY=https://goproxy.io,direct
   
   GOSUMDB=off
   ```
3. 可通过如下指令引入依赖库
   ```
   go get git@github.com:super-rain/qylibrary.git
   ```
4. 如有其他问题，欢迎提[Issue](https://github.com/super-rain/qylibrary/issues/new)

## 包结构

以下为基础目录，每个包的使用方法见包下ReadMe说明及Example

├── gcodex
    │   ├── README.md 
    │   ├── doc.go
    │   ├── errors.go
    │   ├── example_test.go
    │   ├── gcodex.go
    │   ├── gcodex_test.go
    │   └── types.go

## 版本说明

[版本说明](https://github.com/super-rain/qylibrary/releases)