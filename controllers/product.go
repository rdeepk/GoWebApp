package controllers

import (
	"github.com/gorilla/mux"
	"github.com/rdeepk/GoWebApp/controllers/util"
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"strconv"
	"text/template"
)

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProduct(id)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, r)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)
	} else {
		w.WriteHeader(404)
	}
}
