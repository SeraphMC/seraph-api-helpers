package cubelify

import (
	"github.com/flosch/pongo2/v6"
	"strings"
)

type ResponseTagBuilder struct {
	tag CubelifyResponseTag
}

func NewCubelifyTagBuilder() *ResponseTagBuilder {
	return &ResponseTagBuilder{}
}

func (b *ResponseTagBuilder) GetRawData() CubelifyResponseTag {
	return b.tag
}

func (b *ResponseTagBuilder) SetIconName(iconName string) *ResponseTagBuilder {
	if !strings.HasPrefix(iconName, "mdi-") {
		iconName = "mdi-" + iconName
	}
	b.tag.IconName = iconName
	return b
}

func (b *ResponseTagBuilder) SetTextLabel(textLabel string) *ResponseTagBuilder {
	b.tag.TextLabel = textLabel
	return b
}

func (b *ResponseTagBuilder) SetTextLabelWithTemplating(defaultTextLabel, templateTextLabelString string, valueMap map[string]interface{}) *ResponseTagBuilder {
	b.tag.ToolTipLabel = defaultTextLabel
	template, err := pongo2.FromString(templateTextLabelString)
	if err != nil {
		return b
	}

	result, err := template.Execute(valueMap)
	if err != nil {
		return b
	}

	b.tag.ToolTipLabel = result
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

func (b *ResponseTagBuilder) SetToolTipLabelWithTemplating(defaultTooltip, templateTooltipString string, valueMap map[string]interface{}) *ResponseTagBuilder {
	b.tag.ToolTipLabel = defaultTooltip
	template, err := pongo2.FromString(templateTooltipString)
	if err != nil {
		return b
	}

	result, err := template.Execute(valueMap)
	if err != nil {
		return b
	}

	b.tag.ToolTipLabel = result
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
