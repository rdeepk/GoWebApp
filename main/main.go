package main

import (
	//"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("test").Parse(doc)
	if err == nil {
		//context := Context{"the message"}
		tmpl.Execute(w, r.URL.Path)
	}
}

const doc = `
	<!DOCTYPE html>
		<html>
			<head><title>Go Web App</title></head>
			<body>
			{{if eq . "/Google"}}
				<h1>Hey, Google made Go!</h1>
			{{else}}
				<h1>Hello {{.}}</h1>
			{{end}}
			</body>
		</html>
`

/*type Context struct {
	Message string
}*/
