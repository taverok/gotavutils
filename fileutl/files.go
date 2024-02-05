package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func MapYml(file string, target any) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, target)
}
