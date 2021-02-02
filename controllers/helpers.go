package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func ParseForm(r *http.Request, form interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(form, r.PostForm)
	if err != nil {
		return err
	}
	return nil
}
