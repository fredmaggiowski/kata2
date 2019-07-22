// Package bowl is not supposed to be empty.
package bowl

// Score is amazing.
type Score struct {
	Total int          `json:"total"`
	Valid bool         `json:"valid"`
	Error BowlingError `json:"error"`
}

// SetError is sad.
func (s *Score) SetError(err BowlingError) {
	s.Valid = false
	if s.Error.code == 0 || err.code < s.Error.code {
		s.Error = err
	}
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

		if !isValidRoll(roll) {
			score.SetError(InvalidRoll)
		}

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

		if !isValidRoll(rolls[i+1]) {
			score.SetError(InvalidRoll)
		}

		if !isValidFrame(roll, rolls[i+1]) {
			score.SetError(InvalidFrame)
		}

		currentScoreMultiplier, tScoreMultipliers = getAndShift(scoreMultipliers)
		scoreMultipliers = tScoreMultipliers
		score.Total += rolls[i+1] * currentScoreMultiplier

		if currentFrame > 10 {
			score.SetError(TooManyRolls)
		}

		i++
	}

	return score
}
