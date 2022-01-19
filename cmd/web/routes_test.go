package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jerevick83/HOTEL-MGT/internals/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	r := routes(&app)

	switch v := r.(type) {
	case *chi.Mux:
	//do nothing
	default:
		t.Error(fmt.Sprintf("type is not and http.handler, %T", v))
	}
}
