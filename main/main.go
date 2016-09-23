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
		tmpl.Execute(w, r.URL.Path)
	}
}

const doc = `
	<!DOCTYPE html>
		<html>
			<head><title>Go Web App</title></head>
			<body>
				<h1>Hello Templates</h1>
			</body>
		</html>
`
