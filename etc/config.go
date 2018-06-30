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

type config struct {
	Mysql mysqlConfig `yaml:"mysql"`
}

func LoadConfigFromFile(filePath string) (*config, error) {
	var cfg config
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
