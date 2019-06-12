package handler

import (
	"encoding/json"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

// JSONServer -
func jsonServer(code int, obj interface{}, w http.ResponseWriter) {
	MyRender(code, JSON{Data: obj}, w)
}

// Rende -
type Rende interface {
	// Render writes data with custom ContentType.
	render(http.ResponseWriter) error
	// WriteContentType writes custom ContentType.
	WriteContentType(w http.ResponseWriter)
}

// JSON -
type JSON struct {
	Data interface{}
}

// Render (JSON) writes data with custom ContentType.
func (r JSON) render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// WriteContentType (JSON) writes JSON ContentType.
func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// MyRender writes the response headers and calls render.Render to render data.
func MyRender(code int, r Rende, w http.ResponseWriter) {
	w.WriteHeader(code)

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(w)
		//	c.Writer.WriteHeaderNow()
		return
	}

	if err := r.render(w); err != nil {
		panic(err)
	}
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}
