package bowl

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

func isValidRoll(roll int) bool {
	return roll >= 0 && roll <= 10
}

func isValidFrame(rollA, rollB int) bool {
	return isValidRoll(rollA + rollB)
}
