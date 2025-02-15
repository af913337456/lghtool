package LghTool

import "encoding/json"

func EchoJson(obj interface{}) string {
	return JsonString(obj)
}

func JsonString(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(data)
}

func JsonStringIndent(obj interface{}) string {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return ""
	}
	return string(data)
}
