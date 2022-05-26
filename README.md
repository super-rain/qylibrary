# qylibrary——提供简单易用、功能完整、高质量的 Go 微服务工具包


## 说明

1. Go Version 1.13+
2. 使用该库需要设置如下环境变量
   ```
   GOPROXY=http://goproxy.qingting-hz.com
   
   GOSUMDB=off
   ```
3. 可通过如下指令引入依赖库
   ```
   go get git2.qingtingfm.com/neirong/library
   ```
4. 如有其他问题，欢迎提[Issue](https://git2.qingtingfm.com/neirong/library/issues/new)

## 目录

以下为基础目录，每个包的使用方法见包下ReadMe说明及Example

- base:基础组件
    - ctime:自定义时间类型
    - decimal:小数处理常用工具
    - deepcopy:深拷贝结构体
    - encrypt:加密常用工具
    - filewriter:可分割文件的writer
    - net:网络常用工具
    - logrender:日志输出渲染工具
    - null: 空值处理工具
    - reflect:反射常用工具
    - runtime:运行时常用工具
    - strings:字符串处理常用工具
- database:数据库组件
    - etcd:etcd客户端
    - mongo:mongo客户端
    - redis:redis客户端
    - sql:mysql客户端
- goroutine:协程组件
- kafka:Kafka消息队列组件
- log:日志组件
- net:网络组件
    - circuitbreaker:断路器
    - cm:配置中心客户端
    - errcode:错误码
    - gin:gin常用工具
    - httpclient:http客户端
    - metric:数据监控工具
    - middleware:中间件
    - redlock:redis分布式锁工具
    - response:基础响应
    - tracing:链路跟踪工具
- queue:AMQP消息队列组件

## 版本说明

[版本说明](https://git2.qingtingfm.com/neirong/library/releases)