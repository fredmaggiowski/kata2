package bowl

// BowlingError is an error. :(
type BowlingError struct {
	code    int
	message string
}

var (
	// InvalidRoll is an error. :(
	InvalidRoll = BowlingError{1, "invalid roll"}

	// InvalidFrame is an error. :(
	InvalidFrame = BowlingError{2, "invalid frame"}

	// TooManyRolls is an error. :(
	TooManyRolls = BowlingError{3, "too many rolls"}

	// NotEnoughRolls is an error. :(
	NotEnoughRolls = BowlingError{4, "not enough rolls"}
)
