package cubelify

type ResponseTagBuilder struct {
	tag CubelifyResponseTag
}

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

func NewCubelifyTagBuilder() *ResponseTagBuilder {
	return &ResponseTagBuilder{}
}

func (b *ResponseTagBuilder) SetIconName(iconName string) *ResponseTagBuilder {
	b.tag.IconName = iconName
	return b
}

func (b *ResponseTagBuilder) SetTextLabel(textLabel string) *ResponseTagBuilder {
	b.tag.TextLabel = textLabel
	return b
}

func (b *ResponseTagBuilder) SetTextColour(textColour uint32) *ResponseTagBuilder {
	b.tag.TextColour = textColour
	return b
}

func (b *ResponseTagBuilder) SetToolTipLabel(toolTipLabel string) *ResponseTagBuilder {
	b.tag.ToolTipLabel = toolTipLabel
	return b
}

func (b *ResponseTagBuilder) SetBackgroundColour(colour ColourCode) *ResponseTagBuilder {
	b.tag.Colour = colour
	return b
}

func (b *ResponseTagBuilder) SetTagName(tagName string) *ResponseTagBuilder {
	b.tag.TagName = tagName
	return b
}

func (b *ResponseTagBuilder) Build() CubelifyResponseTag {
	return b.tag
}
