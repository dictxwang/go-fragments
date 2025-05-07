package main

import (
	"fmt"
	"go-fragments/src/encrypt"
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

	//ldflagsSample()
	//
	//exception.SampleMain()
	//finally.SampleMain()
	//method.SampleMain()
	//channel.SampleMain()
	//sync.SampleMain()
	//context.SampleMain()
	//
	//config.SampleMainToml()
	//config.SampleMainViperYaml01()
	//config.SampleMainViperYaml02()
	//config.SampleMainViperFlag01()
	//config.SampleMainViperFlag02()
	//config.SampleMainViperYmalUnmarshal()
	//config.SampleMainViperJsonUnmarshal()
	//config.SampleMainGoConfig()
	//
	//file.SampleMain()
	//log.SampleMainLog()
	//log.SampleMainLogrus()
	//log.SampleMainLogrusRotate()
	//runtime.SampleMain()
	//
	//_interface.SampleMain()
	//_reflect.SampleMain()
	//_reflect.SampleMainType()
	//_http.SampleMain()
	//_errors.SampleMain()
	//functional.SampleMain()
	//
	//generic.SampleMainSimple()
	//generic.SampleMainPower()
	//server.SampleMain()
	//tcp.SampleMain()

	//_gin.SampleMainFirst()
	//_gin.SampleMainSecond()

	//net.NewMain()

	//id.UuidMain()

	encrypt.SampleMain()
}
