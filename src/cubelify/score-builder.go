package cubelify

type CubelifyScoreBuilder struct {
	score *CubelifyScore
}

// NewCubelifyScoreBuilder creates and returns a new instance of CubelifyScoreBuilder with an initialised CubelifyScore.
func NewCubelifyScoreBuilder() *CubelifyScoreBuilder {
	return &CubelifyScoreBuilder{score: &CubelifyScore{}}
}

// SetValue sets the score's value to the specified floating-point number and returns the builder instance.
func (b *CubelifyScoreBuilder) SetValue(value float64) *CubelifyScoreBuilder {
	b.score.Value = value
	return b
}

// SetMode sets the mode of the score to either "add" or "set", depending on the input string, and returns the builder. Invalid or empty values are ignored.
func (b *CubelifyScoreBuilder) SetMode(mode string) *CubelifyScoreBuilder {
	if mode != "" && (mode == "add" || mode == "set") {
		b.score.Mode = &mode
	}
	return b
}

// AddValue increments the current score value of the CubelifyScore by the specified value and returns the updated builder instance.
func (b *CubelifyScoreBuilder) AddValue(value float64) *CubelifyScoreBuilder {
	b.score.Value = b.score.Value + value
	return b
}

// RemoveValue subtracts the specified value from the current score's value and returns the updated builder.
func (b *CubelifyScoreBuilder) RemoveValue(value float64) *CubelifyScoreBuilder {
	b.score.Value = b.score.Value - value
	return b
}

// Build finalises the CubelifyScoreBuilder and returns a CubelifyScore instance.
// Returns nil if the score value is zero.
func (b *CubelifyScoreBuilder) Build() *CubelifyScore {
	if b.score.Value == 0 {
		return nil
	}
	return b.score
}
