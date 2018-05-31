package config

import (
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Mongo MongoConfig
	Http HTTP
}

type MongoConfig struct {
	Host string
	Name string
}

type HTTP struct {
	HostAddress string
	Port int
}

func LoadConfig(configPath string) Config {

	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	configString := string(buf)
	if _, err := toml.Decode(configString, &config); err != nil {
		panic(err)
	}

	return config		
}

var Conf Config
