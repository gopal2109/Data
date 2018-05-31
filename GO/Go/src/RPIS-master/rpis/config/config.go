package config

import (
	"fmt"
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Mongo Mongo
	Http HTTP
	Logging Logging
	Application Application
}

type Mongo struct {
	Host string
	Name string
}

type HTTP struct {
	HostAddress string
	Port int
}

type Logging struct {
	File string
	Level string
}

type Application struct {
	TimeZone string
	DateFormat string
}

func LoadConfig(configPath string) (error, Config) {
	var config Config
	
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
		return err, config
	}

	configString := string(buf)
	if _, err := toml.Decode(configString, &config); err != nil {
		fmt.Println(err.Error())
		return err, config
	}

	return nil, config		
}

var Conf Config
