/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"dkstar88/mathgen/arithmetic"
	Gen "dkstar88/mathgen/generator"
	"dkstar88/mathgen/output"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"hash/fnv"
	"math/rand"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mathgen",
	Short: "Math Test Paper Generator",
	Long: `mathgen is a math test paper generator outputs as PDF paper with answers, 
currently support simple arithmetics including addition,
subtraction, multiplication and division. You can also 
set difficulties such as no negatives and no greater than 999, etc. 

To generate example paper try included "ks2" config file:

mathgen ks2.yaml -t "KS2 Simple Arithmetics" -o ks2.pdf -a ks2-answers.pdf

For other more advanced configurations please find it in readme.md file. 

`,
	Args: cobra.MinimumNArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

var (
	configFile  string
	title       string
	questionPdf string
	answerPdf   string
	seed        string
	logLevel    uint32
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&title, "title", "t", "MathGen Test Paper", "Title of the output paper")
	rootCmd.Flags().StringVarP(&questionPdf, "output", "o", "test.pdf", "Output test paper file path")
	rootCmd.Flags().StringVarP(&answerPdf, "answer", "a", "", "Output answer paper file path")
	rootCmd.Flags().StringVarP(&seed, "seed", "s", "", "MathGen generator seed")
	rootCmd.Flags().Uint32VarP(&logLevel, "verbose-level", "v", 1, "Verbose level: [1,2,3,4]")
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))

		// Load Generators
		arithmetic.Register()

		configFile = args[0]
		if seed != "" {
			h := fnv.New64a()
			_, _ = h.Write([]byte(seed))
			rand.Seed(int64(h.Sum64()))
		} else {
			rand.Seed(time.Now().UnixNano())
		}

		Generate(Gen.LoadGeneratorFile(configFile))
	}
}

func initConfig() {

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
	output.GenerateTestPaper(qas, title, questionPdf, answerPdf)
}
