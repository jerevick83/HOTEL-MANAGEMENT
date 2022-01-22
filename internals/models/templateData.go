package models

import "github.com/jerevick83/HOTEL-MGT/internals/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// DataMap holds the data to the rendered pages
	DataMap map[string]interface{}

	FlashMessage   string
	ErrorMessage   string
	WarningMessage string
	CSRFToken      string
	Form           *forms.Form
}
