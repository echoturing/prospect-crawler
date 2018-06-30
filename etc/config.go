package etc

import (
	"os"
	"gopkg.in/yaml.v2"
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

type config struct {
	Mysql mysqlConfig `yaml:"mysql"`
}

func LoadConfigFromFile(filePath string) (*config, error) {
	cfg := &config{}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
