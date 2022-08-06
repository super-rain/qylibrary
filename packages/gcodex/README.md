# qylibrary.gcodex包


## 说明

1. Golang 版本要求最低 1.13
2. 使用该库需要设置如下环境变量
3. 使用 wire 框架实现
   ```
   GOPROXY=https://goproxy.io,direct
   
   GOSUMDB=off
   ```
4. 可通过如下指令引入依赖库
   ```
   go get git@github.com:super-rain/qylibrary.git
   ```

## 功能

以下为基础目录，每个包的使用方法见包下ReadMe说明及Example

├── gcodex
    │   ├── README.md 
    │   ├── doc.go
    │   ├── errors.go
    │   ├── example_test.go
    │   ├── gcodex.go
    │   ├── gcodex_test.go
    │   └── types.go

## 设计实现

>领域建模(PB文件)->CLI命令->meta data->go yacc/go generate->Source code

## TODO

- [ ] grpc
- [ ] go generate
- [ ] go doc comment
- [ ] gpromise

## 使用

- 输入 `gcodex generate` 命令生成 ent 模版源码文件
- 进入 ./ent 同级目录，执行 `go generate ./ent` 或者 在根目录执行“`go run -mod=mod entgo.io/ent/cmd/ent generate ./dist/ent/schema`” 生成模型 CRUD 代码
- `gocodex generate --type=svc` 生成 `erpo/service/type/endpoint/transport` 源码文件
- `gcpdex generate --type=pb` 生成 grpc 模版源码文件
## 版本说明

[版本说明](https://github.com/super-rain/qylibrary/releases)