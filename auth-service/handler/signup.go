package handler

import (
	"auth-service/models"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		err = models.CreateUser(newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		return
	}
}
