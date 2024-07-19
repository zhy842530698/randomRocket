package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name   string
	Passwd string
	Port   uint16
	Server *Server
}
type Server struct {
	Ip    [4]uint8
	Port  uint16
	Alias string
}

func NewConfig(filName string) (*Config, error) {
	data, err := os.ReadFile(filName)
	if err != nil {
		return nil, err
	}
	var c *Config
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
