package cubelify

type ResponseTagBuilder struct {
	tag CubelifyResponseTag
}

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
