package main

import (
	"fmt"
	"go-fragments/src/channel"
	"go-fragments/src/config"
	"go-fragments/src/context"
	"go-fragments/src/exception"
	"go-fragments/src/file"
	"go-fragments/src/finally"
	_http "go-fragments/src/http"
	_interface "go-fragments/src/interface"
	"go-fragments/src/log"
	"go-fragments/src/method"
	_reflect "go-fragments/src/reflect"
	"go-fragments/src/runtime"
	"go-fragments/src/server"
	"go-fragments/src/sync"
	"go-fragments/src/tcp"
)

// 获取编译参数 ldflags
var ENV string
var VERSION string

func ldflagsSample() {
	fmt.Println("[ldflags_sample]")
	fmt.Printf("ENV=%s\n", ENV)
	fmt.Printf("VERSION=%s\n", VERSION)
}

func main() {

	ldflagsSample()

	exception.SampleMain()
	finally.SampleMain()
	method.SampleMain()
	channel.SampleMain()
	sync.SampleMain()
	context.SampleMain()

	config.SampleMainToml()
	config.SampleMainViperYaml01()
	config.SampleMainViperYaml02()
	config.SampleMainViperFlag01()
	config.SampleMainViperFlag02()
	config.SampleMainViperYmalUnmarshal()
	config.SampleMainViperJsonUnmarshal()
	config.SampleMainGoConfig()

	file.SampleMain()
	log.SampleMainLog()
	log.SampleMainLogrus()
	log.SampleMainLogrusRotate()
	runtime.SampleMain()

	_interface.SampleMain()
	_reflect.SampleMain()
	_http.SampleMain()

	server.SampleMain()
	tcp.SampleMain()
}
