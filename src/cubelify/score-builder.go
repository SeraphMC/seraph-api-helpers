package cubelify

type CubelifyScoreBuilder struct {
	score *CubelifyScore
}

func NewCubelifyScoreBuilder() *CubelifyScoreBuilder {
	return &CubelifyScoreBuilder{score: &CubelifyScore{}}
}

func (b *CubelifyScoreBuilder) SetValue(value float64) *CubelifyScoreBuilder {
	b.score.Value = value
	return b
}

func (b *CubelifyScoreBuilder) SetMode(mode string) *CubelifyScoreBuilder {
	b.score.Mode = mode
	return b
}

func (b *CubelifyScoreBuilder) Build() *CubelifyScore {
	return b.score
}
