package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorf struct {
	Error string `json:"error"`
}

func Json(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Json(w, statusCode, errorf{Error: err.Error()})

	} else {
		Json(w, statusCode, nil)
	}
}
