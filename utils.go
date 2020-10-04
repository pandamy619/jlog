package jlog

import (
	"encoding/json"
	"fmt"
)

func convertToJson(obj interface{}) []byte{
	jsonObj, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return jsonObj
}