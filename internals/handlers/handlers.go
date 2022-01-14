package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jerevick83/HOTEL-MGT/internals/config"
	"github.com/jerevick83/HOTEL-MGT/internals/models"
	"github.com/jerevick83/HOTEL-MGT/internals/render"
	"log"
	"net/http"
)

// TemplateData holds data sent to the templates

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["gender"] = "Jeremiah Victor Harding"
	remoteIp := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_Ip", remoteIp)
	render.RenderTemplate(w, req, "home.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, req, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Availability(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, req, "search-availability.page.gohtml", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start")
	end := req.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is: %s\n, and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, req *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available",
	}

	output, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (m *Repository) Generals(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, req, "generals.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Majors(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, req, "majors-suites.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, req, "contact.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) MakeReservation(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, req, "makeReservation.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
