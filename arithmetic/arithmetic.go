package arithmetic

import (
	"dkstar88/mathgen/generator"
	"dkstar88/mathgen/types"
	"fmt"
	"github.com/Knetic/govaluate"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

//const valid_operators = {"+", "-", "*", "/"}

type GenerationRules struct {
	Operations     []string
	Nums           []int
	Len            int
	Max            int
	Min            int
	Difficulty     int
	MustBeInt      bool
	IncludeBracket bool
}

var MaxRetry = 10

func Arithmetic(rules GenerationRules) (*generator.QuestionAnswer, error) {
	log.Infof("Arithmetic: %v", rules)
	n1 := rules.Nums[0]
	n2 := rules.Nums[1]
	questionNums := make([]int, rules.Len)
	operators := make([]string, rules.Len-1)
	retry := 0
	answer := 0
	question := ""
	for i := 0; i < rules.Len; i++ {
		questionNums[i] = randomMinMax(n1, n2)
		// Check if generated number is within difficulty setting
		isValid := intIsWithin(questionNums[i], rules.Min, rules.Max, rules.Difficulty)
		if !isValid {
			// Regenerate
			i--
			if retry >= MaxRetry {
				log.Warnf("retried too many times %d", retry)
				return nil, fmt.Errorf("failed to generate addition %v", rules)
			}
			retry++
			continue
		}
		if i > 0 {
			operator := randomOperator(rules.Operations)
			operators[i-1] = operator
			question = formQuestion(questionNums, operators, i)
			log.Info(question)
			expression, err := govaluate.NewEvaluableExpression(question)
			if err != nil {
				log.Errorf("NewEvaluableExpression: %v", err)
			}
			result, err := expression.Evaluate(nil)
			if err != nil {
				log.Errorf("Evaluate: %v", err)
			}
			answer = int(result.(float64))
			isValid = intIsWithin(answer, rules.Min, rules.Max, rules.Difficulty)
			if rules.MustBeInt {
				isValid = isValid && mustBeInt(result.(float64))
			}
			if !isValid {
				// Regenerate
				i--
				if retry >= MaxRetry {
					log.Warnf("retried too many times %d", retry)
					return nil, fmt.Errorf("failed to generate addition %v", rules)
				}
				retry++
				continue
			}
		}
	}

	return &generator.QuestionAnswer{
		Question: question,
		Answer:   fmt.Sprintf("%d", answer),
	}, nil

}

const IntPrecision = 10000

func formQuestion(numbers []int, operators []string, length int) string {
	result := ""
	for i := 0; i <= length; i++ {
		result = result + fmt.Sprintf("%d", numbers[i])
		if i < length {
			result = result + operators[i]
		}
	}
	return result
}

func mustBeInt(result float64) bool {
	flt := result
	return int(flt)*IntPrecision == int(flt*IntPrecision)
}

func randomOperator(operations []string) string {
	r := rand.Intn(len(operations))
	return operations[r]
}

func intIsWithin(i, min, max, diff int) bool {
	overMax := i > max
	if overMax {
		log.Warnf("%d is exceeding max: %d", i, max)
		return false
	}
	lessMin := i < min
	if lessMin {
		log.Warnf("%d is less than min: %d", i, min)
		return false
	}
	thisDiff := detectDifficultyInt(i)
	overDifficulty := diff > 0 && thisDiff > diff
	if overDifficulty {
		log.Warnf("%d difficulty level is %d, and is over difficult level %d", i, thisDiff, diff)
		return false
	}
	return true
}

func init() {
	generator.RegisterGenerator("arithmetic", func(config map[string]interface{}) (*generator.QuestionAnswer, error) {
		log.Infof("Wrap Arithmetic: %v", config)
		rule := GenerationRules{
			Operations: types.StrArrDef(config["operations"], []string{"+", "-"}),
			Nums:       types.IntArrDef(config["nums"], []int{1, 10}),
			Len:        types.IntDef(config["len"], 2),
			Max:        types.IntDef(config["max"], 100),
			Min:        types.IntDef(config["min"], 1),
			Difficulty: types.IntDef(config["difficulty"], 0),
			MustBeInt:  types.BoolDef(config["mustBeInt"], true),
		}
		return Arithmetic(rule)
	})
}
