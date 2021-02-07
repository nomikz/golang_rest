package controllers

import (
	"../views"
)

type Static struct {
	Home     *views.View
	Contacts *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("bootstrap", "static/index"),
		Contacts: views.NewView("bootstrap", "static/contacts"),
	}
}
