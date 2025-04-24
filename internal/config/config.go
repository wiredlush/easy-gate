package config

import (
	"encoding/json"
	"fmt"

	"github.com/wiredlush/easy-gate/internal/group"
	"github.com/wiredlush/easy-gate/internal/theme"
	"gopkg.in/yaml.v3"
)

// Service - Easy Gate service configuration struct
type Service struct {
	Icon     string   `json:"icon" yaml:"icon"`
	Name     string   `json:"name" yaml:"name"`
	URL      string   `json:"url" yaml:"url"`
	Category string   `json:"category" yaml:"category"`
	Groups   []string `json:"groups" yaml:"groups"`
}

// Note - Easy Gate note configuration struct
type Note struct {
	Name     string   `json:"name" yaml:"name"`
	Text     string   `json:"text" yaml:"text"`
	Category string   `json:"category" yaml:"category"`
	Groups   []string `json:"groups" yaml:"groups"`
}

// Config - Easy Gate configuration struct
type Config struct {
	Theme       theme.Theme   `json:"theme" yaml:"theme"`
	Addr        string        `json:"addr" yaml:"addr"`
	Title       string        `json:"title" yaml:"title"`
	CertFile    string        `json:"cert_file" yaml:"cert_file"`
	KeyFile     string        `json:"key_file" yaml:"key_file"`
	Groups      []group.Group `json:"groups" yaml:"groups"`
	Services    []Service     `json:"services" yaml:"services"`
	Notes       []Note        `json:"notes" yaml:"notes"`
	BehindProxy bool          `json:"behind_proxy" yaml:"behind_proxy"`
	UseTLS      bool          `json:"use_tls" yaml:"use_tls"`
}

type format int

const (
	jsonFormat format = iota + 1
	yamlFormat
)

// Unmarshal - Unmarshal config bytes into config struct
func Unmarshal(configBytes []byte) (*Config, error) {
	config := &Config{}

	allowedFormats := []format{jsonFormat, yamlFormat}
	for _, format := range allowedFormats {
		switch format {
		case jsonFormat:
			err := json.Unmarshal(configBytes, config)
			if err == nil {
				return config, nil
			}
		case yamlFormat:
			err := yaml.Unmarshal(configBytes, config)
			if err == nil {
				return config, nil
			}
		}
	}

	return nil, fmt.Errorf("invalid configuration file format")
}
