package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Forward *ForwardConfig `yaml:"forward"`
}

type ForwardConfig struct {
	TargetHost string `yaml:"target_host"`
	ServerPort string `yaml:"server_port"`
}

func GetConfig() (*Config, error) {
	f, err := os.Open("./config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("打开config.yaml文件失败: %v", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("读取config.yaml文件失败: %v", err)
	}
	conf := &Config{}
	if err = yaml.Unmarshal(b, conf); err != nil {
		return nil, fmt.Errorf("反序列化config.yaml文件失败: %v", err)
	}
	return conf, nil
}
