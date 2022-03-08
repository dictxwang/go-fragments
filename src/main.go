package main

import (
	"fmt"
	"go-fragments/src/channel"
	"go-fragments/src/config"
	"go-fragments/src/context"
	"go-fragments/src/exception"
	"go-fragments/src/file"
	"go-fragments/src/finally"
	"go-fragments/src/log"
	"go-fragments/src/method"
	"go-fragments/src/server"
	"go-fragments/src/sync"
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
	file.SampleMain()
	log.SampleMainLog()
	log.SampleMainLogrus()
	log.SampleMainLogrusRotate()

	config.SampleMainGoConfig()

	server.SampleMain()
}
