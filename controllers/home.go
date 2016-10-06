package controllers

import (
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetHome()
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w, vm)
}
