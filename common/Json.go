package common

import "encoding/json"

//Takes an arbitrary struct and returns the json string of it
func Json(v interface{}) string {
	d, _ := json.Marshal(v)
	return string(d)
}
