package cubelify

import (
	"github.com/Clemintina/common_utilities-for-apis/src/validation"
	"time"
)

type CubelifyResponseBuilder struct {
	response *CubelifyResponse
}

func NewCubelifyResponseBuilder() *CubelifyResponseBuilder {
	return &CubelifyResponseBuilder{response: &CubelifyResponse{}}
}

func (b *CubelifyResponseBuilder) SetScore(score *CubelifyScore) *CubelifyResponseBuilder {
	b.response.Score = score
	return b
}

func (b *CubelifyResponseBuilder) SetTags(tags []CubelifyResponseTag) *CubelifyResponseBuilder {
	b.response.Tags = &tags
	return b
}

func (b *CubelifyResponseBuilder) SetTimestamp(timestamp *time.Time) *CubelifyResponseBuilder {
	b.response.Timestamp = timestamp
	return b
}

func (b *CubelifyResponseBuilder) SetError(error string) *CubelifyResponseBuilder {
	b.response.Error = error
	return b
}

func (b *CubelifyResponseBuilder) SetType(typeStr string) *CubelifyResponseBuilder {
	b.response.Type = typeStr
	return b
}

func (b *CubelifyResponseBuilder) Build() *CubelifyResponse {
	if b.response.Timestamp == nil {
		b.response.Timestamp = validation.ToTimePointer(time.Now())
	}

	return b.response
}
