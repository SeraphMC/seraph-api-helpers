package cubelify

import (
	"github.com/SeraphMC/seraph-api-helpers/src/validation"
	"time"
)

type CubelifyResponseBuilder struct {
	response *CubelifyResponse
}

// NewCubelifyResponseBuilder creates and initialises a new instance of CubelifyResponseBuilder to construct a CubelifyResponse.
func NewCubelifyResponseBuilder() *CubelifyResponseBuilder {
	return &CubelifyResponseBuilder{response: &CubelifyResponse{}}
}

// SetScore sets the Score field of the CubelifyResponse with the provided CubelifyScore instance and returns the updated CubelifyResponseBuilder.
func (b *CubelifyResponseBuilder) SetScore(score *CubelifyScore) *CubelifyResponseBuilder {
	b.response.Score = score
	return b
}

// AddSniperScore increments the existing score value of the response by the value of the given CubelifyScore. If no score already exists, it initialises the score with the provided one. Returns the updated builder instance.
func (b *CubelifyResponseBuilder) AddSniperScore(score *CubelifyScore) *CubelifyResponseBuilder {
	if b.response.Score == nil {
		b.response.Score = &CubelifyScore{}
	}
	b.response.Score.Value += score.Value
	return b
}

// SetTags sets the tags of the CubelifyResponse to the provided slice of CubelifyResponseTag and returns the builder instance.
func (b *CubelifyResponseBuilder) SetTags(tags []CubelifyResponseTag) *CubelifyResponseBuilder {
	b.response.Tags = &tags
	return b
}

// AddTags appends a slice of CubelifyResponseTag to the Tags field and initialises the Tags field if nil, then returns the builder itself.
func (b *CubelifyResponseBuilder) AddTags(tags []CubelifyResponseTag) *CubelifyResponseBuilder {
	if b.response.Tags == nil {
		b.response.Tags = &[]CubelifyResponseTag{}
	}
	*b.response.Tags = append(*b.response.Tags, tags...)
	return b
}

// AddTag appends a single CubelifyResponseTag to the Tags slice in the response, initialising the Tags slice if it is nil, and returns the updated CubelifyResponseBuilder.
func (b *CubelifyResponseBuilder) AddTag(tags CubelifyResponseTag) *CubelifyResponseBuilder {
	if b.response.Tags == nil {
		b.response.Tags = &[]CubelifyResponseTag{}
	}
	*b.response.Tags = append(*b.response.Tags, tags)
	return b
}

// SetTimestamp sets the timestamp of the CubelifyResponse to the provided time value and returns the builder instance.
func (b *CubelifyResponseBuilder) SetTimestamp(timestamp *time.Time) *CubelifyResponseBuilder {
	b.response.Timestamp = timestamp
	return b
}

// SetError sets the error message in the CubelifyResponse and returns the updated CubelifyResponseBuilder.
func (b *CubelifyResponseBuilder) SetError(error string) *CubelifyResponseBuilder {
	b.response.Error = error
	return b
}

// SetType sets the type of the response to the specified string and returns the updated CubelifyResponseBuilder instance.
func (b *CubelifyResponseBuilder) SetType(typeStr string) *CubelifyResponseBuilder {
	b.response.Type = typeStr
	return b
}

// Build finalizes and returns the constructed CubelifyResponse object. If the Timestamp field is nil, it sets it to the current time.
func (b *CubelifyResponseBuilder) Build() *CubelifyResponse {
	if b.response.Timestamp == nil {
		b.response.Timestamp = validation.ToTimePointer(time.Now())
	}

	return b.response
}
