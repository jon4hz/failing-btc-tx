package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	config *Config
)

type Config struct {
	RPCUser     string `yaml:"rpcuser"`
	RPCPassword string `yaml:"rpcpassword"`
	RPCHost     string `yaml:"rpchost"`
}

func Load() {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal("error while reading the config file: ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("error while parsing the config file: ", err)
	}
}

func Get() *Config {
	if config == nil {
		panic("config not loaded")
	}
	return config
}
