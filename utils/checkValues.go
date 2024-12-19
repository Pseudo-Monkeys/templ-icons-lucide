package utils

func CheckValue(currentValue, defaultValue string) string {
    if currentValue == ""{
        return defaultValue
    }else{
        return currentValue
    }
}
