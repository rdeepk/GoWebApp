package controllers

import (
	"github.com/rdeepk/GoWebApp/controllers/util"
	"github.com/rdeepk/GoWebApp/models"
	"github.com/rdeepk/GoWebApp/viewmodels"
	"net/http"
	"text/template"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetHome()
	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}

func (this *homeController) login(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		member, err := models.GetMember(email, password)
		if err == nil {
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId()
				responseWriter.Header().Add("Set-Cookie", cookie.String())
			}
		}
	}
	vm := viewmodels.GetLogin()
	w.Header().Add("Content-Type", "text/html")
	this.loginTemplate.Execute(responseWriter, vm)
}
