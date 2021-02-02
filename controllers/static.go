package controllers

import (
	"../views"
)

type Static struct {
	HomeView     *views.View
	ContactsView *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:     views.NewView("bootstrap", "./views/static/index.gohtml"),
		ContactsView: views.NewView("bootstrap", "./views/static/contacts.gohtml"),
	}
}
