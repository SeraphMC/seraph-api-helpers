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
	if mode != "" && (mode == "add" || mode == "set") {
		*b.score.Mode = mode
	}
	return b
}

func (b *CubelifyScoreBuilder) AddValue(value float64) *CubelifyScoreBuilder {
	b.score.Value = b.score.Value + value
	return b
}

func (b *CubelifyScoreBuilder) RemoveValue(value float64) *CubelifyScoreBuilder {
	b.score.Value = b.score.Value - value
	return b
}

func (b *CubelifyScoreBuilder) Build() *CubelifyScore {
	return b.score
}
