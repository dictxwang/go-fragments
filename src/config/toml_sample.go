package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

type Person struct {
	Name string
	OnLine bool
	DOB time.Time
	Authors []string
	Databases map[string]database `config:"dbs"`
	Clients clients
}

type database struct {
	Host string
	Port int
	DBName string
}

type clients struct {
	Data [][]interface{}
	Hosts []string
}

func SampleMainToml() {

	fmt.Println("\n[config_toml_sample]")
	var config Person
	if _, err := toml.DecodeFile("default.conf.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	// like var_dump in PHP
	fmt.Printf("%v\n", config)
	fmt.Println(config.Clients.Data)
}