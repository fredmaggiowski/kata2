// Package bowl is not supposed to be empty.
package bowl

// Score is amazing.
type Score struct {
	Total int `json:"total"`
}

func isStrike(roll int) bool {
	return roll == 10
}

func isSpare(rollA int, rollB int) bool {
	return isStrike(rollA + rollB)
}

func getAndShift(multipliers []int) (int, []int) {
	current := multipliers[0]
	return current, []int{multipliers[1], 1}
}

// GetScore is funny.
func GetScore(rolls []int) Score {
	var score Score

	scoreMultipliers := []int{1, 1}
	currentFrame := 0

	for i := 0; i < len(rolls); i++ {
		currentFrame++

		roll := rolls[i]
		currentScoreMultiplier, tScoreMultipliers := getAndShift(scoreMultipliers)
		scoreMultipliers = tScoreMultipliers
		//  := scoreMultipliers[0]
		// scoreMultipliers[0] = scoreMultipliers[1]
		// scoreMultipliers[1] = 1

		score.Total += roll * currentScoreMultiplier
		if isStrike(roll) && currentFrame != 10 {
			scoreMultipliers[0]++
			scoreMultipliers[1]++
			continue
		}

		if i == len(rolls)-1 {
			continue
		}

		if isSpare(roll, rolls[i+1]) && currentFrame != 10 {
			scoreMultipliers[1]++
		}

		currentScoreMultiplier, tScoreMultipliers = getAndShift(scoreMultipliers)
		scoreMultipliers = tScoreMultipliers
		score.Total += rolls[i+1] * currentScoreMultiplier

		i++
	}
	return score
}
