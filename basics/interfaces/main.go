package main

import (
	"encoding/json"
	"net/http"
)

type UserNotifier interface {
	NotifyUserCreated(user) error
}

type user struct {
	Username string `json:"user_name"`
	Email    string `json:"user_email"`
}

type UserHandler struct {
	UserNotifier UserNotifier
}

func (u UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {

	var user user
	json.NewDecoder(r.Body).Decode(&user)

	u.UserNotifier.NotifyUserCreated(user)

	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func main() {
	mux := http.NewServeMux()

	UserHandler := UserHandler{
		UserNotifier: BetterNotifier{},
	}

	mux.HandleFunc("POST /create_user", UserHandler.handleCreateUser)

	http.ListenAndServe(":3000", mux)

}
