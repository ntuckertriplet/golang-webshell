package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
    "os/exec"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()

	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rawCommand := r.PostForm.Get("command")
	commandOutput := execute(rawCommand)
	fmt.Printf("%s", commandOutput)
	http.Redirect(w, r, "/", 302)
}

func execute(commd string) string {
	out, err := exec.Command(commd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out[:])
	fmt.Printf("%s", output)
	return output
}


