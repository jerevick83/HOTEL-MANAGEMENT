package models

type TemplateData struct {
	StringMap      map[string]string
	IntMap         map[string]int
	FloatMap       map[string]float32
	DataMap        map[string]interface{}
	FlashMessage   string
	ErrorMessage   string
	WarningMessage string
	CSRFToken      string
}
