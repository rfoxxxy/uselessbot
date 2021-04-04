package main

func replacer(key string, value interface{}) (bool, string) {
	str, ok := value.(string)
	// cut the string
	if ok && len(str) > 30 {
		return true, `"` + str[0:35] + `..."`
	}
	return false, ""
}
