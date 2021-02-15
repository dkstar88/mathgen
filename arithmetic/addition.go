package arithmetic

import (
	"dkstar88/mathgen/generator"
	"dkstar88/mathgen/types"
	"fmt"
)

func Addition(rules GenerationRules) (*generator.QuestionAnswer, error) {
	n1 := rules.Nums[0]
	n2 := rules.Nums[1]
	questionNums := make([]int, rules.Len)
	answer := 0
	difficulty := 0
	retry := 0
	for i := 0; i < rules.Len; i++ {
		questionNums[i] = randomMinMax(n1, n2)
		thisDiff := detectDifficultyInt(questionNums[i])
		if answer+questionNums[i] > rules.Max || difficulty+thisDiff > rules.Difficulty {
			// Regenerate
			i--
			if retry > MaxRetry {
				return nil, fmt.Errorf("failed to generate addition %v", rules)
			}
			retry++
			continue
		}
		answer += questionNums[i]
		difficulty += thisDiff
	}
	return &generator.QuestionAnswer{
		Question:   joinInts(questionNums, "-"),
		Answer:     fmt.Sprintf("%d", answer),
		Difficulty: difficulty,
	}, nil
}

func init() {
	generator.RegisterGenerator("Addition", func(config map[string]interface{}) (*generator.QuestionAnswer, error) {
		rule := GenerationRules{
			Nums:       types.IntArrDef(config["min"], []int{1, 10}),
			Len:        types.IntDef(config["len"], 2),
			Max:        types.IntDef(config["max"], 100),
			Min:        types.IntDef(config["min"], 1),
			Difficulty: types.IntDef(config["difficulty"], 2),
		}
		return Addition(rule)
	})
}
