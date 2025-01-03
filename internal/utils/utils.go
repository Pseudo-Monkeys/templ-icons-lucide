package utils

import (
	"github.com/Pseudo-Monkeys/templ-icons-lucide/lib"
	"github.com/a-h/templ"
)

func CheckValue(currentValue, defaultValue string) string {
    if currentValue == "" {
        return defaultValue
    }
    return currentValue
}

func getBasicProps(props lib.IconProps, iconClass string) *templ.Attributes{
    return &templ.Attributes{
        "xmlns":"http://www.w3.org/2000/svg",
        "width":CheckValue(props.Size, "24"),
        "height":CheckValue(props.Size, "24"),
        "viewBox":CheckValue(props.ViewBox, "0 0 24 24"),
        "fill":CheckValue(props.Fill, "none"),
        "stroke":CheckValue(props.Stroke, "currentColor"),
        "stroke-width":CheckValue(props.StrokeWidth, "2"),
        "stroke-linecap":CheckValue(props.StrokeLinecap, "round"),
        "stroke-linejoin":CheckValue(props.StrokeLinejoin, "round"),
        "class":props.Class + " lucide " + iconClass,
        "style":"display:block;",
    }
}
