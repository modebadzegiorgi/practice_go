package main

import (
	"log"
	"net/http"

	"github.com/modebadzegiorgi/practice_go/beginner_programs/todo_app/backend/router"
)

func main() {
	r := router.Router()

	log.Fatal(http.ListenAndServe(":3000", r))

}
