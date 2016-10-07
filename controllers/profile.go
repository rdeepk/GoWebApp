package controllers

import (
	"github.com/rdeepk/GoWebApp/controllers/util"
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"text/template"
)

type profileController struct {
	template *template.Template
}

func (this *profileController) handle(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()

	vm := viewmodels.GetProfile()
	if r.Method == "POST" {

		vm.User.Email = r.FormValue("email")
		vm.User.FirstName = r.FormValue("firstName")
		vm.User.LastName = r.FormValue("lastName")
		vm.User.Stand.Address = r.FormValue("standAddress")
		vm.User.Stand.City = r.FormValue("standCity")
		vm.User.Stand.State = r.FormValue("standState")
		vm.User.Stand.Zip = r.FormValue("standZip")
	}

	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}
