package utils

import "github.com/a-h/templ"

func CheckValue(currentValue, defaultValue string) string {
    if currentValue == "" {
        return defaultValue
    }
    return currentValue
}

type IconProps struct {
    Class,
    ViewBox,
    Fill,
    Stroke,
    StrokeWidth,
    StrokeLinecap,
    StrokeLinejoin,
    Size string
}

func getBasicProps(props IconProps) *templ.Attributes{
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
        "class":CheckValue(props.Class, "lucide lucide-chevron-left"),
        "style":"display:block;",
    }
}
