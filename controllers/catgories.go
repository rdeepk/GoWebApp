package controllers

import (
	"github.com/gorilla/mux"
	"github.com/rdeepk/GoWebApp/controllers/util"
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"strconv"
	"text/template"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetCategories()
	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProducts(id)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, r)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)
	} else {
		w.WriteHeader(404)
	}
}
