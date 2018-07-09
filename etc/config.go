package etc

import (
	"os"

	"github.com/go-yaml/yaml"
)

type mysqlConfig struct {
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Net       string `yaml:"net"`
	Addr      string `yaml:"addr"`
	DBName    string `yaml:"db_name"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parse_time"`
}

type DingAlert struct {
	Usage string `json:"usage"`
	Url   string `json:"url"`
}

type Config struct {
	Mysql     mysqlConfig `yaml:"mysql"`
	DingAlert DingAlert   `yaml:"ding_alert"`
}

func LoadConfigFromFile(filePath string) (*Config, error) {
	var cfg Config
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
