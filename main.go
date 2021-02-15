package main

import (
	Gen "dkstar88/mathgen/generator"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	Generate(Gen.LoadGeneratorFile("ks2.yaml"))
}

func Generate(generatorFile Gen.GeneratorFile) {
	var qas []Gen.QuestionAnswer
	for _, generator := range generatorFile.Generators {
		qa, err := Gen.Generate(generator.Type, generator.Params)
		if err != nil {
			log.Errorf("failed to generate %v, error: %v", generator, err.Error())
		}
		if qa != nil {
			qas = append(qas, *qa)
		}
	}
	GenerateTestPaper(qas, "Test Paper 2")
}
