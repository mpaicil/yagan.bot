package config

import(
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Redis struct {
		Host string `yaml: "host"`
		Port string `yaml:"port"`
	} `yaml: "redis"`
	Token string `yaml: "token"`
}

var LoadedConfig *Config

func LoadConfig(){
	LoadedConfig = &Config{}

	file, err := os.Open("./config.yml")
    if err != nil {
        fmt.Println("Quedo la cagá")
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)
	
	if err := d.Decode(LoadedConfig); err != nil {
        fmt.Println("Requedo la cagá")
	}
}