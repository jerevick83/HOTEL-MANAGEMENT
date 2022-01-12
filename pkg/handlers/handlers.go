package handlers

import (
	"github.com/jerevick83/HOTEL-MGT/pkg/config"
	"github.com/jerevick83/HOTEL-MGT/pkg/models"
	"github.com/jerevick83/HOTEL-MGT/pkg/render"
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
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Reservation(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, "reservation.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Generals(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, "generals.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Majors(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, "majors-suites.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(req.Context(), "remote_Ip")
	stringMap["remoteIP"] = remoteIp
	render.RenderTemplate(w, "contact.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
