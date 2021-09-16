package serialization

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	ExampleYamlFile = "example.yaml"
)

func NewPerson(yamlData string) (*Person, error) {
	person := Person{}
	err := yaml.Unmarshal([]byte(yamlData), &person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (person *Person) ToString() (string, error) {
	data, err := yaml.Marshal(person)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadExampleFile() (string, error) {
	data, err := os.ReadFile(ExampleYamlFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}