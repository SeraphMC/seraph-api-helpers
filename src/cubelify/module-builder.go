package cubelify

import (
	"strings"

	"github.com/flosch/pongo2/v6"
)

type ResponseTagBuilder struct {
	tag CubelifyResponseTag
}

// NewCubelifyTagBuilder creates and returns a new instance of ResponseTagBuilder, used to configure and construct a CubelifyResponseTag.
func NewCubelifyTagBuilder() *ResponseTagBuilder {
	return &ResponseTagBuilder{}
}

// GetRawData returns the underlying CubelifyResponseTag instance being built by the ResponseTagBuilder.
func (b *ResponseTagBuilder) GetRawData() CubelifyResponseTag {
	return b.tag
}

// SetIconName sets the icon name for the response tag, automatically prefixing it with "mdi-" if the provided name does not already have this prefix. Returns the updated ResponseTagBuilder instance.
func (b *ResponseTagBuilder) SetIconName(iconName string) *ResponseTagBuilder {
	if !strings.HasPrefix(iconName, "mdi-") {
		iconName = "mdi-" + iconName
	}
	b.tag.IconName = iconName
	return b
}

// SetTextLabel sets the text label of the response tag and returns the updated ResponseTagBuilder instance.
func (b *ResponseTagBuilder) SetTextLabel(textLabel string) *ResponseTagBuilder {
	b.tag.TextLabel = textLabel
	return b
}

// SetTextLabelWithTemplating sets the TextLabel property using the provided default text and applies a template if valid.
// If the template execution fails, the default text label is retained.
func (b *ResponseTagBuilder) SetTextLabelWithTemplating(defaultTextLabel, templateTextLabelString string, valueMap map[string]interface{}) *ResponseTagBuilder {
	b.tag.TextLabel = defaultTextLabel
	template, err := pongo2.FromString(templateTextLabelString)
	if err != nil {
		return b
	}

	result, err := template.Execute(valueMap)
	if err != nil {
		return b
	}

	b.tag.TextLabel = result
	return b
}

// SetTextColour sets the text colour of the tag using a uint32 value and returns the ResponseTagBuilder for chaining.
func (b *ResponseTagBuilder) SetTextColour(textColour uint32) *ResponseTagBuilder {
	b.tag.TextColour = textColour
	return b
}

// SetToolTipLabel sets the tooltip label for the associated tag and returns the updated ResponseTagBuilder.
func (b *ResponseTagBuilder) SetToolTipLabel(toolTipLabel string) *ResponseTagBuilder {
	b.tag.ToolTipLabel = toolTipLabel
	return b
}

// SetToolTipLabelWithTemplating updates the ToolTipLabel property using a default value and a templated string.
// It applies the provided value map to the template and falls back to the default if template execution fails.
// Returns the modified ResponseTagBuilder.
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

// SetBackgroundColour sets the background colour of the tag using the provided ColourCode and returns the updated ResponseTagBuilder.
func (b *ResponseTagBuilder) SetBackgroundColour(colour ColourCode) *ResponseTagBuilder {
	b.tag.Colour = colour
	return b
}

// SetTagName sets the tag name for the response tag and returns the updated ResponseTagBuilder instance.
func (b *ResponseTagBuilder) SetTagName(tagName string) *ResponseTagBuilder {
	b.tag.TagName = tagName
	return b
}

// Build finalises and returns the constructed CubelifyResponseTag instance.
func (b *ResponseTagBuilder) Build() CubelifyResponseTag {
	return b.tag
}
