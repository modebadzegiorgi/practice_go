package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoTask struct {
	Body string `json:"todo"`
}

var tasks []TodoTask

func CreateTask(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var requestBody TodoTask
	json.Unmarshal(body, &requestBody)

	tasks = append(tasks, requestBody)
	fmt.Println(tasks)

}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	r.HandleFunc("/todo", CreateTask)
	return r
}
