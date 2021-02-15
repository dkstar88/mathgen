package generator

import "errors"

type GenerateFunc func(config map[string]interface{}) (*QuestionAnswer, error)

type Generator struct {
	name    string
	execute GenerateFunc
}

var (
	generators map[string]*Generator
)

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
	generators[g.name] = g
}

func RegisterGenerator(name string, execute GenerateFunc) {
	Register(NewGenerator(name, execute))
}

func Generate(name string, config map[string]interface{}) (*QuestionAnswer, error) {
	g, err := generators[name]
	if !err {
		return nil, errors.New(name + " generator does not exist!")
	}
	return g.execute(config)
}
