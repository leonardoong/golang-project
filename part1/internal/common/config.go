package common

import (
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	GrpcPort string      `yaml:"grpc_port"`
	Redis    RedisConfig `yaml:"redis"`
	Kafka    KafkaConfig `yaml:"kafka"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

type KafkaConfig struct {
	Broker    string `yaml:"broker"`
	Zookeeper string `yaml:"zookeeper"`
}

func ReadConfig() (cfg Config, err error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	configPath := filepath.Join(exPath, "../../files/config.yaml")

	file, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&cfg)

	return
}
