package cubelify

type ResponseTagBuilder struct {
	tag CubelifyResponseTag
}

type ColourCode = uint32

const (
	GREEN    ColourCode = 0x4ADE80
	YELLOW   ColourCode = 0xFACC15
	AMBER    ColourCode = 0xFCD34D
	RED      ColourCode = 0xEF4444
	DarkRed  ColourCode = 0xb91c1c
	LightRed ColourCode = 0xf87171
	CYAN     ColourCode = 0x22D3EE
	PINK     ColourCode = 0xE879F9
	GREY     ColourCode = 0x6b7280
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

func (b *ResponseTagBuilder) Build() CubelifyResponseTag {
	return b.tag
}
