package utils

func CheckValue(currentValue, defaultValue string) string {
    if currentValue == "" {
        return currentValue
    }

    return defaultValue
}
