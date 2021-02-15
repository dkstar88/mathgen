package generator

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

type GenerateFunc func(config map[string]interface{}) (*QuestionAnswer, error)

type Generator struct {
	name    string
	execute GenerateFunc
}

var (
	generators map[string]*Generator
)

func init() {
	generators = make(map[string]*Generator)
}

func NewGenerator(name string, execute GenerateFunc) *Generator {
	generator := &Generator{
		name:    name,
		execute: execute,
	}
	return generator
}

func (g *Generator) GetName() string {
	return g.name
}

func (g *Generator) Execute(config map[string]interface{}) (*QuestionAnswer, error) {
	return g.execute(config)
}

func Register(g *Generator) {
	log.Infof("Generator Registering: %v", g)
	generators[strings.ToLower(g.name)] = g
	log.Infof("Generator Registered: %s", g.name)
}

func RegisterGenerator(name string, execute GenerateFunc) {
	log.Infof("Generator Registering: %s", name)
	Register(NewGenerator(name, execute))
}

func Generate(name string, config map[string]interface{}) (*QuestionAnswer, error) {
	g, err := generators[strings.ToLower(name)]
	if !err {
		return nil, errors.New(name + " generator does not exist!")
	}
	return g.execute(config)
}
