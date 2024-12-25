package utils

func CheckValue(currentValue, defaultValue string) string {
    if currentValue == "" {
        return currentValue
    }

    return defaultValue
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
