package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/hashicorp/raft"
	"k8s.io/klog"
)

var SysConfig *Config

type Server struct {
	ID      raft.ServerID
	Address raft.ServerAddress
	Http    string
}

// Config 不同节点配置文件
type Config struct {
	ServerName  string `yaml:"server-name"`
	ServerID    string `yaml:"server-id"`
	LogStore    string
	StableStore string
	Snapshot    string //快照保存的位置
	Transport   string
	Servers     []Server
	Port        string
	LocalCache  string `yaml:"local-cache"`
}

func NewConfig() *Config {
	return &Config{}
}

func loadConfigFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		klog.Errorf("load file err: %s", err)
		return nil
	}
	return b
}

// LoadConfig 读取配置文件
func LoadConfig(path string) (*Config, error) {
	config := NewConfig()
	if b := loadConfigFile(path); b != nil {
		err := yaml.Unmarshal(b, config)
		if err != nil {
			klog.Errorf("unmarshal err: %s", err)
			return nil, err
		}
		return config, err
	} else {
		return nil, fmt.Errorf("load config file error")
	}
}
