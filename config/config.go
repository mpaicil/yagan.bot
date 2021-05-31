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
	Database struct {
		User string `yaml: "user"`
		Pass string `yaml: "pass"`
		Host string `yaml: "host"`
		Port string `yaml: "port"`
		DBName string `yaml: "dbname"`
		Queries map[string]string `yaml: "queries,flow`
	} `yaml: "database"`
	
}

func DBProperty() string{
	return LoadedConfig.Database.User + ":"+ LoadedConfig.Database.Pass+"@tcp("+LoadedConfig.Database.Host+":"+LoadedConfig.Database.Port+")/"+LoadedConfig.Database.DBName
}

func RetrieveQuery(name string) string {
	return LoadedConfig.Database.Queries[name]
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

	fmt.Println(LoadedConfig)
}