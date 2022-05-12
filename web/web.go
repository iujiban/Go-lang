package web

import (
	"log"
	"net/http"
	"text/template"
)

func index(res http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("template/index.htm"))
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func signUp(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/signUp.htm"))
		err := tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln("error executing template", err)
		}
	}
	if req.Method == "POST" {
		//Create(db, req)
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
}
