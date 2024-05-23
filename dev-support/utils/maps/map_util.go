package maps

func GetMapString(dataMap map[string]interface{}, key string) string {
	val := dataMap[key]
	strValue, ok := val.(string)
	if ok {
		return strValue
	}
	return ""
}
func GetMapInt(dataMap map[string]interface{}, key string) int {
	val := dataMap[key]
	strValue, ok := val.(int)
	if ok {
		return strValue
	}
	return 0
}
func GetMapFloat64(dataMap map[string]interface{}, key string) float64 {
	val := dataMap[key]
	strValue, ok := val.(float64)
	if ok {
		return strValue
	}
	return 0
}

func GetMapBool(dataMap map[string]interface{}, key string) bool {
	val := dataMap[key]
	boolValue, ok := val.(bool)
	if ok {
		return boolValue
	}
	panic("类型不匹配!")
}
