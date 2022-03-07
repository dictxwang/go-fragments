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
	fmt.Printf("name=%s\n", name)
	fmt.Printf("version=%f\n", version)
}