package arithmetic

import (
	"dkstar88/mathgen/generator"
	"dkstar88/mathgen/types"
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

const (
	ADDITION       = "+"
	SUBTRACTION    = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
)

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
	question := ""
	for i := 0; i < rules.Len; i++ {
		questionNums[i] = randomMinMax(n1, n2)
		if !intIsWithin(questionNums[i], rules.Min, rules.Max, rules.Difficulty) {
			// Regenerate
			i--
			if retry > MaxRetry {
				log.Warnf("retried too many times %d", retry)
				return nil, fmt.Errorf("failed to generate addition %v", rules)
			}
			retry++
			continue
		}
		if i >= rules.Len-1 {
			question = question + fmt.Sprintf("%d", questionNums[i])
		} else {
			operator := randomOperator(rules.Operations)
			operators[i] = operator
			question = question + fmt.Sprintf("%d%s", questionNums[i], operator)
		}
	}
	expression, err := govaluate.NewEvaluableExpression(question)
	if err != nil {
		log.Errorf("NewEvaluableExpression: %v", err)
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Errorf("Evaluate: %v", err)
	}
	answer := int(result.(float64))

	if intIsWithin(answer, rules.Min, rules.Max, rules.Difficulty) {
		return &generator.QuestionAnswer{
			Question: question,
			Answer:   fmt.Sprintf("%d", answer),
		}, nil
	}
	return nil, errors.New("failed to generate")

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
		}
		return Arithmetic(rule)
	})
}
