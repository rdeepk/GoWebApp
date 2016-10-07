package controllers

import (
	"encoding/json"
	"github.com/rdeepk/GoWebApp/controllers/util"
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"text/template"
)

type standLocatorController struct {
	template *template.Template
}

func (this *standLocatorController) get(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

/*func (this *standLocatorController) apiSearch(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocations()

	responseWriter.Header().Add("Content-Type", "application/json")
	this.template.Execute(responseWriter, vm)

	data, err := json.Marshal(vm)
	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(404)
	}
}*/

func (this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetStandLocations()
	responseWriter.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(vm)
	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(404)
	}
}
