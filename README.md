# go-fragments

golang示例代码片段，基于 go mod 模式编写，go版本1.16。

## 代码导航

[1、通过channel进行协程间通信](https://github.com/dictxwang/go-fragments/blob/main/src/channel/channel_sample.go)

**[2、配置文件解析](https://github.com/dictxwang/go-fragments/tree/main/src/config)**

[2.1、配置文件解析（goconfig）](https://github.com/dictxwang/go-fragments/blob/main/src/config/goconfig_sample.go)

[2.2、配置文件解析（toml）](https://github.com/dictxwang/go-fragments/blob/main/src/config/toml_sample.go)

[2.3、配置文件解析（viper）_推荐](https://github.com/dictxwang/go-fragments/blob/main/src/config/viper_sample.go)

[3、通过context进行协程控制和信息传递](https://github.com/dictxwang/go-fragments/blob/main/src/context/context_sample.go)

[4、业界准标准的异常处理包errors](https://github.com/dictxwang/go-fragments/blob/main/src/errors/errors_sample.go)

[5、异常捕获与恢复（defer、panic与recover）](https://github.com/dictxwang/go-fragments/blob/main/src/exception/exception_sample.go)

[6、文件操作示例](https://github.com/dictxwang/go-fragments/blob/main/src/file/file_sample.go)

[7、go中的finally实现（正确使用defer）](https://github.com/dictxwang/go-fragments/blob/main/src/finally/finally_sample.go)

[8、函数式编程实践](https://github.com/dictxwang/go-fragments/blob/main/src/functional/functional_sample.go)

**[9、泛型的实现与应用（go1.17前无泛型）](https://github.com/dictxwang/go-fragments/tree/main/src/generic)**

[9.1、模拟实现泛型的map/reduce/filer](https://github.com/dictxwang/go-fragments/blob/main/src/generic/generic_sample.go)

[9.1、泛型的map/reduce/filter实现的健壮版本](https://github.com/dictxwang/go-fragments/blob/main/src/generic/generic_power_sample.go)

[10、http的应用示例](https://github.com/dictxwang/go-fragments/blob/main/src/http/http_sample.go)

[11、interface接口相关](https://github.com/dictxwang/go-fragments/blob/main/src/interface/interface_sample.go)

**[12、日志处理](https://github.com/dictxwang/go-fragments/tree/main/src/log)**

[12.1、通过log输出std日志和file日志](https://github.com/dictxwang/go-fragments/blob/main/src/log/log_sample.go)

[12.2、logrus模块的应用](https://github.com/dictxwang/go-fragments/blob/main/src/log/logrus_sample.go)

[12.3、基于logrus实现滚动日志](https://github.com/dictxwang/go-fragments/blob/main/src/log/logrus_rotate_sample.go)

[13、方法中指针参数和非指针参数的区别](https://github.com/dictxwang/go-fragments/blob/main/src/method/method_sample.go)

[14、reflect反射的应用](https://github.com/dictxwang/go-fragments/blob/main/src/reflect/reflect_sample.go)

[15、通过runtime获取系统信息（cpu、操作系统等）](https://github.com/dictxwang/go-fragments/blob/main/src/runtime/runtime_sample.go)

[16、基于http搭建服务端](https://github.com/dictxwang/go-fragments/blob/main/src/server/server_sample.go)

[17、使用sync模块实现协程同步](https://github.com/dictxwang/go-fragments/blob/main/src/sync/sync_sample.go)

[18、通过net模块实现tcp服务端和客户端](https://github.com/dictxwang/go-fragments/blob/main/src/tcp/tcp_sample.go)

### 工程构建

[1、go工程的交叉编译（跨平台编译）](https://github.com/dictxwang/go-fragments/blob/main/build.sh)