package config

import (
	"fmt"
	"io"

	fs "github.com/hack-pad/hackpadfs"
	"go.yaml.in/yaml/v3"
)

const (
	configPath = "config.yml"
)

func InitHolder(fs fs.FS) (Holder, error) {
	file, err := fs.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling config file: %w", err)
	}

	return NewHolder(config), nil
}
