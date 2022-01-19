package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/searchAvail", "GET", []postData{}, http.StatusOK},
	{"generals", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suites", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"makeReservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"search availability", "/searchAvail", "POST", []postData{
		{key: "start", value: "2022-01-19"},
		{key: "end", value: "2022-01-25"},
	}, http.StatusOK},
	{"search Availability", "/searchAvail-json", "POST", []postData{
		{key: "start", value: "2022-01-19"},
		{key: "end", value: "2022-01-25"},
	}, http.StatusOK},
	{"make Reservation", "/make-reservation", "POST", []postData{
		{key: "fName", value: "Jeremiah Victor"},
		{key: "lName", value: "Harding"},
		{key: "arrival-date", value: "2022-01-19"},
		{key: "departure-date", value: "2022-01-25"},
		{key: "email", value: "jere@mail.com"},
		{key: "phone", value: "0874892"},
	}, http.StatusOK},
	//{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := testServer.Client().Get(testServer.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := testServer.Client().PostForm(testServer.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}
	}
}
