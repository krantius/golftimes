package app

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("internal/templates/index.html"))

type Server struct {
	API *API
}

func (s *Server) Courses(w http.ResponseWriter, req *http.Request) {
	date := "01-30-2021"
	times, err := s.API.GetTimes(MilesSquare, date)
	if err != nil {
		panic(err)
	}

	d := Data{
		Date: date,
	}

	ct := CourseTimes{Name: MilesSquare.Name}
	for _, t := range times {
		ct.Times = append(ct.Times, TeeTime{
			Time:  t.Time,
			Price: t.CartFee + t.GreenFee,
		})
	}

	d.Courses = append(d.Courses, ct)

	if err := tmpl.Execute(w, d); err != nil {
		w.WriteHeader(500)
		return
	}
}

func (s *Server) Template(req *http.Request, w http.ResponseWriter) {
	tmpl.Execute(w, struct{ Name string }{Name: "asdf"})
}
