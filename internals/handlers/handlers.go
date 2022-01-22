package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jerevick83/HOTEL-MGT/internals/config"
	"github.com/jerevick83/HOTEL-MGT/internals/forms"
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
func (m *Repository) PostAvailability(w http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start")
	end := req.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is: %s\n, and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
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
	// emptyReservation gets all the fields of the Reservation and pass them on to the reservation field with the data they hold
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, req, "makeReservation.page.gohtml", &models.TemplateData{
		Form:    forms.New(nil),
		DataMap: data,
	})
}

// PostMakeReservation handles the data posted from the make-reservation form
func (m *Repository) PostMakeReservation(w http.ResponseWriter, req *http.Request) {
	// ParseForm initially checks the form for any errors and stores any error in the error variables
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	//reservation getting the values from the posted form and storing them in the reservation struct fields
	reservation := models.Reservation{
		Firstname:     req.Form.Get("fName"),
		Lastname:      req.Form.Get("lName"),
		Email:         req.Form.Get("email"),
		Phone:         req.Form.Get("phone"),
		ArrivalDate:   req.Form.Get("arrival-date"),
		DepartureDate: req.Form.Get("departure-date"),
	}

	// PostForm contains the data from ParseForm PATCH POST PUT form body parameter
	form := forms.New(req.PostForm)

	// Required checks whether the values from the form are not empty, else will return
	form.Required("fName", "lName", "email", "arrival-date", "departure-date")

	// MinLength checks whether the values from the form are equal to the specified length, else will return
	form.MinLength("fName", 5)

	// IsEmail checks whether the value from the form email field has the proper structure of and email address, else will return
	form.IsEmail("email")

	// Valid returns true if there are no form errors, otherwise false
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, req, "makeReservation.page.gohtml", &models.TemplateData{
			Form:    form,
			DataMap: data,
		})
		return
	}

	//Put adds a key and corresponding value to the session data. Any existing value for the key will be replaced. The session data status will be set to Modified.
	m.App.Session.Put(req.Context(), "reservation", reservation)

	//redirecting client to another page (reservation-summary) after completing the booking
	http.Redirect(w, req, "/reservation-summary", http.StatusSeeOther)
}
func (m *Repository) ReservationSummary(w http.ResponseWriter, req *http.Request) {
	// Getting item out of the session
	reservationInfo, ok := m.App.Session.Get(req.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("Cannot get item from session")
		m.App.Session.Put(req.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, req, "/make-reservation", http.StatusTemporaryRedirect)
		return
	}

	// Removing items out of the session
	m.App.Session.Remove(req.Context(), "reservation")
	data := make(map[string]interface{})

	// assigning a field to the DataMap struct
	data["reservation"] = reservationInfo
	render.RenderTemplate(w, req, "reservation-summary.page.gohtml", &models.TemplateData{
		DataMap: data,
	})
}
