package main

import 	(
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseFiles("layout.html"))

type Server struct {
	
}

func (s *Server) Template(req *http.Request, w http.ResponseWriter) {
	tmpl.Execute(w, nil)
}