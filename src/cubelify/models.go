package cubelify

import "time"

type ColourCode = uint32

const (
	Green    ColourCode = 0x4ADE80
	Yellow   ColourCode = 0xFACC15
	Amber    ColourCode = 0xFCD34D
	Red      ColourCode = 0xEF4444
	DarkRed  ColourCode = 0xb91c1c
	LightRed ColourCode = 0xf87171
	Cyan     ColourCode = 0x22D3EE
	Pink     ColourCode = 0xE879F9
	Grey     ColourCode = 0x6b7280
)

// CubelifyResponse represents the response structure for a Cubelify operation, containing scoring information, associated tags, timestamp, error details, and response type.
type CubelifyResponse struct {
	Score     *CubelifyScore         `json:"score,omitempty"`
	Tags      *[]CubelifyResponseTag `json:"tags,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
	Error     string                 `json:"error,omitempty"`
	Type      string                 `json:"-,omitempty"`
}

// CubelifyScore represents a scoring system with a numerical value and an optional mode of evaluation.
type CubelifyScore struct {
	Value float64 `json:"value,omitempty"`
	Mode  *string `json:"mode,omitempty"`
}

// CubelifyResponseTag represents a tag associated with a cubelify response, including visual attributes such as icon, text, color, tooltip, and name.
type CubelifyResponseTag struct {
	IconName     string `json:"icon,omitempty"`
	TextLabel    string `json:"text,omitempty"`
	TextColour   uint32 `json:"textColor,omitempty"`
	ToolTipLabel string `json:"tooltip,omitempty"`
	Colour       uint32 `json:"color,omitempty"`
	TagName      string `json:"tag_name,omitempty"`
}
