package arithmetic

import (
	"dkstar88/mathgen/generator"
	"dkstar88/mathgen/types"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func Addition(rules GenerationRules) (*generator.QuestionAnswer, error) {
	log.Infof("Addition: %v", rules)
	n1 := rules.Nums[0]
	n2 := rules.Nums[1]
	questionNums := make([]int, rules.Len)
	answer := 0
	difficulty := 0
	retry := 0
	for i := 0; i < rules.Len; i++ {
		questionNums[i] = randomMinMax(n1, n2)
		thisDiff := detectDifficultyInt(questionNums[i])
		overMax := questionNums[i] > rules.Max
		if overMax {
			log.Warnf("%d is exceeding max: %d", questionNums[i], rules.Max)
		}
		overDifficulty := rules.Difficulty > 0 && thisDiff > rules.Difficulty
		if overDifficulty {
			log.Warnf("%d difficulty level is %d, and is over difficult level %d", questionNums[i], thisDiff, rules.Difficulty)
		}
		if overMax || overDifficulty {
			// Regenerate
			i--
			if retry > MaxRetry {
				log.Warnf("retried too many times %d", retry)
				return nil, fmt.Errorf("failed to generate addition %v", rules)
			}
			retry++
			continue
		}
		answer += questionNums[i]
		difficulty += thisDiff
	}
	return &generator.QuestionAnswer{
		Question:   joinInts(questionNums, "+"),
		Answer:     fmt.Sprintf("%d", answer),
		Difficulty: difficulty,
	}, nil
}

func init() {
	generator.RegisterGenerator("Addition", func(config map[string]interface{}) (*generator.QuestionAnswer, error) {
		log.Infof("Wrap Addition: %v", config)
		log.Infof("IntArrDef %T\n", config["nums"])
		rule := GenerationRules{
			Nums:       types.IntArrDef(config["nums"], []int{1, 10}),
			Len:        types.IntDef(config["len"], 2),
			Max:        types.IntDef(config["max"], 100),
			Min:        types.IntDef(config["min"], 1),
			Difficulty: types.IntDef(config["difficulty"], 0),
		}
		return Addition(rule)
	})
}
