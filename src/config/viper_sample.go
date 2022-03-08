package config

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"sync"
	"time"
)

// 支持诸多配置类型： "yaml", "yml", "json", "toml", "hcl", "tfvars", "prop", "props", "properties","dotenv", "env", "ini"

type ConfigPerson struct {
	Name string `mapstructure:"name"`
	Clothing map[string]string
	Hobbies []string
}

type ConfigPersonNested struct {
	Name string
	Clothing ConfigClothing
}

type ConfigClothing struct {
	Jacket string
	Trousers string
}

func SampleMainViperYaml01()  {

	fmt.Println("\n[config_viper_sample_yaml01]")

	myViper := viper.New()
	myViper.AddConfigPath("./")
	// 加载本地配置文件
	myViper.SetConfigFile("default.viper.yaml")
	// 设置配置类型
	myViper.SetConfigType("yaml")
	err := myViper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 监听配置文件的变化
	myViper.WatchConfig()

	myViper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed", in.Name)
	})

	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	cond.L.Lock()

	go func() {
		for _, num := range []int{1,2,3,4,5,6,7,8,9,10} {
			fmt.Printf("viper watch %d: name=%s\n", num, myViper.GetString("name"))
			time.Sleep(1 * time.Second)
		}
		cond.Signal()
	}()

	cond.Wait()
	fmt.Println("viper watch: after cond wait...")
}

func SampleMainViperYaml02() {

	fmt.Println("\n[config_viper_sample_yaml02]")

	myViper := viper.New()
	myViper.SetConfigType("yaml")
	// 加载内存中的配置
	var configCb = []byte(`
name: wangqiang
hobbies:
- walking
- reading
app:
  cache1:
    items: 100
  cache2:
    items: 200
`)
	// 加载远程配置
	//myViper.AddRemoteProvider("remoteEtc", "http://127.0.0.1", "/config/conf.json")

	myViper.ReadConfig(bytes.NewBuffer(configCb))
	fmt.Println(myViper.GetString("name"))

	// 获取slice
	fmt.Println(myViper.GetStringSlice("hobbies"))

	// 提取配置子树，这里返回的是一个新的Viper示例
	subv := myViper.Sub("app.cache1")
	fmt.Println(subv.GetString("items"))
}

func SampleMainViperFlag01() {

	fmt.Println("\n[config_viper_sample_flag01]")
	// 与flag并存
	viperflag01 := flag.String("viperflag01", "viperflag01Value", "viper with flag")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	myViper := viper.New()
	myViper.BindPFlags(pflag.CommandLine)
	fmt.Println(myViper.GetString("viperflag01"))
	fmt.Println(*viperflag01)
}

func SampleMainViperFlag02() {

	fmt.Println("\n[config_viper_sample_flag02]")
	// 单独使用pflag
	viperflag02 := pflag.String("viperflag02", "viperflag02Value", "viper with flag")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	myViper := viper.New()
	myViper.BindPFlags(pflag.CommandLine)
	fmt.Println(myViper.GetString("viperflag02"))
	fmt.Println(*viperflag02)
}

func SampleMainViperYmalUnmarshal() {

	fmt.Println("\n[config_viper_sample_yaml_unmarshal]")
	myViper := viper.New()
	myViper.SetConfigFile("default.viper.yaml")
	myViper.SetConfigType("yaml")
	err := myViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	person := new(ConfigPerson)
	if err := myViper.Unmarshal(person); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	fmt.Printf("%v\n", person)
	fmt.Printf("clothing.jacket: %s\n", person.Clothing["jacket"])

	// 结构体嵌套解析
	personNested := new(ConfigPersonNested)
	if err := myViper.Unmarshal(personNested); err != nil {
		panic(fmt.Errorf("unmarshal confNested failed, err:%s \n", err))
	}
	fmt.Printf("%v\n", personNested)
	fmt.Printf("clothing.jacket: %s\n", personNested.Clothing.Jacket)
}

func SampleMainViperJsonUnmarshal() {

	fmt.Println("\n[config_viper_sample_json_unmarshal]")
	myViper := viper.New()
	myViper.SetConfigFile("default.viper.json")
	// 可以不设置类型
	//myViper.SetConfigType("json")
	err := myViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	person := new(ConfigPerson)
	if err := myViper.Unmarshal(person); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	fmt.Printf("%v\n", person)
	fmt.Printf("clothing.jacket: %s\n", person.Clothing["jacket"])
}