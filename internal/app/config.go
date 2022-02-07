package app

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		HttpAddress string `yaml:"http_address"`
		GrpcAddress string `yaml:"grpc_address"`
	}
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		Database int    `yaml:"database"`
	}
}

func readConfig(configPath string) (*config, error) {
	config := &config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
