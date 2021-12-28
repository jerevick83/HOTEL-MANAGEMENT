package handlers

import (
	"github.com/jerevick83/pkg/config"
	"github.com/jerevick83/pkg/models"
	"github.com/jerevick83/pkg/render"
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
	stringMap["name"] = "Jeremiah"
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
