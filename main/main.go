package main

import (
	"github.com/rdeepk/GoWebApp/controllers"
	"net/http"
	"os"
	"text/template"
)

func main() {
	templates := populateTemplates()
	controllers.Register(templates)
	http.ListenAndServe(":8000", nil)
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
