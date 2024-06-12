package cubelify

import "time"

type CubelifyResponse struct {
	Score     *CubelifyScore         `json:"score,omitempty"`
	Tags      *[]CubelifyResponseTag `json:"tags,omitempty"`
	Timestamp time.Time              `json:"timestamp,omitempty"`
	Error     string                 `json:"error,omitempty"`
	Type      string                 `json:"-,omitempty"`
}

type CubelifyScore struct {
	Value float64 `json:"value,omitempty"`
	Mode  string  `json:"mode,omitempty"`
}

type CubelifyResponseTag struct {
	IconName     string `json:"icon,omitempty"`
	TextLabel    string `json:"text,omitempty"`
	TextColour   uint32 `json:"textColor,omitempty"`
	ToolTipLabel string `json:"tooltip,omitempty"`
	Colour       uint32 `json:"color,omitempty"`
}
