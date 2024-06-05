package common

import "net/http"

func SimpleSuccess(w http.ResponseWriter, status int, data any) {
	WriteJSON(w, status, map[string]interface{}{"success": true, "data": data})
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]interface{}{"success": false, "error": message})
}
