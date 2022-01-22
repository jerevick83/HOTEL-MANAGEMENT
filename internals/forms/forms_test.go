package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}
	postedData := url.Values{}
	postedData.Add("a", "b")
	postedData.Add("b", "a")
	postedData.Add("c", "a")
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when is does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}
	postedData := url.Values{}
	postedData.Add("a", "b")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}
func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for a non-existent field")
	}
	postedData := url.Values{}
	postedData.Add("some_field", "some value")

	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}
	postedData = url.Values{}
	postedData.Add("another_field", "abc123")

	form = New(postedData)

	form.MinLength("some_field", 1)
	if form.Valid() {
		t.Error("shows minlength of 1 is not met when data it is")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "sget@mail.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows valid email when email is actually invalid")
	}

	postedData = url.Values{}
	postedData.Add("email", "sget@.com")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows valid email when eamil is actually invalid")
	}
}
