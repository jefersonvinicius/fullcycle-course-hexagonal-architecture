package handlers

import "encoding/json"

func jsonError(msg string) string {
	object := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	r, _ := json.Marshal(object)
	return string(r)
}
