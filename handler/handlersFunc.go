package handler

import (
	"html/template"
	"log"
	"net/http"
)

var temp *template.Template

func RenderTemplate() {
	temp = template.Must(template.ParseGlob("templates/*"))

}
func errors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type pupil struct {
	Name string
	age  int
}

func data() pupil {
	pup1 := pupil{
		Name: "Jeremiah",
		age:  3,
	}
	return pup1
}

//Home page function handler
func Home(w http.ResponseWriter, req *http.Request) {
	hey := data()
	err := temp.ExecuteTemplate(w, "index.gohtml", hey)
	errors(err)
}

//About page handler
func About(w http.ResponseWriter, req *http.Request) {
	err := temp.ExecuteTemplate(w, "about.gohtml", "This is the about page")
	errors(err)
}
