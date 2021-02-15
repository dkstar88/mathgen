package main

import (
	_ "dkstar88/mathgen/arithmetic"
	Gen "dkstar88/mathgen/generator"
	"dkstar88/mathgen/output"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Generate(Gen.LoadGeneratorFile("ks2.yaml"))
}

func Generate(generatorFile Gen.GeneratorFile) {
	log.Infof("Generator: %v", generatorFile)
	var qas []Gen.QuestionAnswer
	for _, generator := range generatorFile.Generators {
		log.Infof("Generating %s ...", generator.Type)
		i := 0
		for {
			qa, err := Gen.Generate(generator.Type, generator.Params)
			if err != nil {
				log.Errorf("failed to generate %v, error: %v", generator, err.Error())
			}
			if qa != nil {
				qas = append(qas, *qa)
				i++
			}
			if i >= generator.Quantity {
				break
			}
		}
	}
	output.GenerateTestPaper(qas, "Test Paper 2")
}
