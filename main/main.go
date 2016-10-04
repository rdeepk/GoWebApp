package main

import (
	"bufio"
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content Type", "text/html")
	templates := populateTemplates()
	requestedFile := r.URL.Path[1:]
	var context interface{} = nil
	switch requestedFile {
	case "home":
		context = viewmodels.GetHome()
	case "categories":
		context = viewmodels.GetCategories()

		//case "products":
		//context = viewmodels.GetProducts()
	}
	template := templates.Lookup(requestedFile + ".html")
	if template != nil {
		template.Execute(w, context)
	} else {
		w.WriteHeader(404)
	}
}

func serveResource(w http.ResponseWriter, r *http.Request) {
	path := "public" + r.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func populateTemplates() *template.Template {
	result := template.New("template")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}
	result.ParseFiles(*templatePaths...)
	return result
}
