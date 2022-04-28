package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdUrl(name string, r *http.Request) uint64 {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars[name], 10, 64)
	if err != nil {
		return 0
	}
	return id
}
