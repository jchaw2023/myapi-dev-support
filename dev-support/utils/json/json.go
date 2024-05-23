package json

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func MarshalToStr(v interface{}) string {
	b, err := json.Marshal(v)
	if err == nil {
		return string(b)
	}
	fmt.Errorf("MarshalToStr,err:=%#v", err)
	return ""
}

func MarshalSortKeyToStr(v interface{}) string {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.MarshalToString(v)
	if err != nil {
		fmt.Errorf("MarshalToStr,err:=%#v", err)
	}
	return jsonStr
}
func UnMarshalStrToObj(v string, o interface{}) error {
	var vb = []byte(v)
	err := json.Unmarshal(vb, o)
	return err
}
