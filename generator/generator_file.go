package generator

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type GeneratorFile struct {
	Meta struct {
		Author string `yaml:"author,omitempty"`
		Date   struct {
		} `yaml:"date,omitempty"`
		Version string   `yaml:"version,omitempty"`
		Label   []string `yaml:"label,omitempty"`
		Year    []int    `yaml:"year,omitempty"`
	} `yaml:"meta"`
	Generators []struct {
		Type     string                 `yaml:"type"`
		Quantity int                    `yaml:"quantity,omitempty"`
		Score    int                    `yaml:"score,omitempty"`
		Params   map[string]interface{} `yaml:"params,omitempty"`
	} `yaml:"generators"`
}

func LoadGeneratorFile(filename string) GeneratorFile {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	t := GeneratorFile{}

	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	return t
}
