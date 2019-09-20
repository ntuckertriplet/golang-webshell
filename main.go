package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"os/exec"
	s "strings"
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
	fmt.Fprintf(w, string(commandOutput))
}

func execute(inputCommd string) []byte {
	var commd []string

	commd = s.Split(inputCommd, " ")
	var out []byte
	if len(commd) > 1 {
		out, _ = exec.Command(string(commd[0]), string(commd[1])).Output()
		fmt.Printf("%s", out)
	}
	out, _ = exec.Command(string(commd[0])).Output();
	fmt.Printf("%s", out)
	return out
}


