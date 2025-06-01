package handler

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("App is running fine"))
	}
}
