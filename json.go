package resputil

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(obj)
	w.WriteHeader(200)
}

func JSONIndent(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonBytes, _ := json.MarshalIndent(obj, "", "    ")
	w.Write(jsonBytes)
	w.WriteHeader(200)
}
