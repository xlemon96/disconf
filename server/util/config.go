package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var G_conf *Config

type Config struct {
	MySqlConfig  *MySqlConfig  `yaml:"mysql"`
	LogConfig    *LogConfig    `yaml:"log"`
	ServerConfig *ServerConfig `yaml:"server"`
}

type MySqlConfig struct {
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
	IP           string `yaml:"ip"`
	Port         string `yaml:"port"`
	DataBase     string `yaml:"database"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}

type LogConfig struct {
	LogPath string `yaml:"logpath"`
}

type ServerConfig struct {
	IP       string `yaml:"ip"`
	Port     string `yaml:"port"`
	HttpPort string `yaml:"http_port"`
}

func InitConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	G_conf = new(Config)
	err = yaml.Unmarshal(data, G_conf)
	if err != nil {
		return err
	}
	return nil
}
