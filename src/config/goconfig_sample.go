package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func SampleMainGoConfig() {

	fmt.Println("\n[goconfig_sample]")

	conf, err := goconfig.LoadConfigFile("default.goconfig.conf")
	if err != nil {
		panic("load config file error")
	}
	name, _ := conf.GetValue(goconfig.DEFAULT_SECTION, "name")
	version := conf.MustFloat64(goconfig.DEFAULT_SECTION, "version")
	ip := conf.MustValue("server", "ip")
	port := conf.MustInt("server", "port", 1024)

	// 获取配置字段注释信息
	versionComments := conf.GetKeyComments(goconfig.DEFAULT_SECTION, "version")

	// 获取array配置
	clientsHosts := conf.MustValueArray("clients", "hosts", ",")

	fmt.Printf("name=%s\n", name)
	fmt.Printf("version=%f\n", version)
	fmt.Printf("server.ip=%s\n", ip)
	fmt.Printf("server.port=%d\n", port)
	fmt.Printf("versionComments=%s\n", versionComments)
	fmt.Printf("versionComments=%s\n", clientsHosts)
}