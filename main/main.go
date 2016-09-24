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
		context := Context{
			[3]string{"Apple", "Mango", "Banana"},
			"Go Web App",
		}
		tmpl.Execute(w, context)
	}
}

const doc = `
	<!DOCTYPE html>
		<html>
			<head><title>{{.Title}}</title></head>
			<body>
				<h1>List of fruits</h1>
				<ul>
				{{range .Fruit}}
					<li>{{.}}</li>
				{{end}}
				</ul>
			</body>
		</html>
`

type Context struct {
	Fruit [3]string
	Title string
}
